import {
  getPlayerUsername,
  getOwnedGames,
  getSteamLevel,
  resolveVanityUrl,
} from './steamApiRequests.js';
import { getYoutubeChannelInfo } from './youtubeApiRequests.js';
import { getTwitchChannelInfo } from './twitchApiRequests.js';
import { getPageInfo } from './vkApiRequests.js';
import fastify from 'fastify';
import NodeCache from 'node-cache';
import { parse } from 'node-html-parser';
import { authMiddleware } from './token.js';

const server = fastify({ logger: true });

const cache = new NodeCache({ stdTTL: 86400, checkperiod: 120 });

//steam

server.get(
  '/steam/username',
  { preHandler: authMiddleware },
  async (request) => {
    const key = request.query.key;
    const id = request.query.id;
    if (!cache.has(`username-${id}`)) {
      cache.set(`username-${id}`, await getPlayerUsername(key, id));
    }
    return { username: cache.get(`username-${id}`) };
  },
);

server.get('/steam/games', { preHandler: authMiddleware }, async (request) => {
  const key = request.query.key;
  const id = request.query.id;
  if (!cache.has(`games-${id}`)) {
    const games = await getOwnedGames(key, id);
    cache.set(`games-${id}`, games);
    for (let game of games) {
      cache.set(`game-${id}-${game.appid}`, {
        game_name: game.name,
        hours: Math.floor(game.playtime_forever / 60),
        icon: game.img_icon_url,
      });
    }
  }
  return { games: cache.get(`games-${id}`) };
});

server.get('/steam/game', { preHandler: authMiddleware }, async (request) => {
  const key = request.query.key;
  const id = request.query.id;
  const appid = request.query.appid;
  if (!cache.has(`games-${id}`)) {
    const games = await getOwnedGames(key, id);
    cache.set(`games-${id}`, games);
    for (let game of games) {
      cache.set(`game-${id}-${game.appid}`, {
        game_name: game.name,
        hours: Math.floor(game.playtime_forever / 60),
        icon: game.img_icon_url,
      });
    }
  }
  return { game: cache.get(`game-${id}-${appid}`) };
});

server.get('/steam/level', { preHandler: authMiddleware }, async (request) => {
  const key = request.query.key;
  const id = request.query.id;
  if (!cache.has(`level-${id}`)) {
    cache.set(`level-${id}`, await getSteamLevel(key, id));
  }
  return { level: cache.get(`level-${id}`) };
});

server.get('/steam/vanity', { preHandler: authMiddleware }, async (request) => {
  const key = request.query.key;
  const vanityurl = request.query.vanityurl;
  if (!cache.has(`id-${vanityurl}`)) {
    cache.set(`id-${vanityurl}`, await resolveVanityUrl(key, vanityurl));
  }
  return { id: cache.get(`id-${vanityurl}`) };
});

//youtube

server.get(
  '/youtube/channel',
  { preHandler: authMiddleware },
  async (request) => {
    const key = request.query.key;
    const handle = request.query.handle ?? request.query.id;
    const isId = !request.query.handle;
    if (!cache.has(`yt-${handle}`)) {
      cache.set(`yt-${handle}`, await getYoutubeChannelInfo(key, handle, isId));
    }
    return { info: cache.get(`yt-${handle}`) };
  },
);

//twitch

server.get(
  '/twitch/channel',
  { preHandler: authMiddleware },
  async (request) => {
    const key = request.query.key;
    const clientId = request.query.clientId;
    const handle = request.query.handle;
    if (!cache.has(`twitch-${handle}`)) {
      cache.set(
        `twitch-${handle}`,
        await getTwitchChannelInfo(key, clientId, handle),
      );
    }
    return { followers: cache.get(`twitch-${handle}`) };
  },
);

//vk

server.get('/vk/page', { preHandler: authMiddleware }, async (request) => {
  const key = request.query.key;
  const handle = request.query.handle;
  if (!cache.has(`vk-${handle}`)) {
    cache.set(`vk-${handle}`, await getPageInfo(key, handle));
  }
  return { info: cache.get(`vk-${handle}`) };
});

server.get('/tg/:channel/:post', async (request, response) => {
  const { channel, post } = request.params;
  response.header('Content-Type', 'text/html; charset=utf-8');
  response.header('Access-Control-Allow-Origin', '*');

  const { color, dark } = request.query;
  console.log(color);
  const html = await fetch(
    `https://t.me/${channel}/${post}?embed=1&userpic=false&color=${encodeURIComponent(color)}&dark-color=${encodeURIComponent(color)}&dark=${dark}`,
  ).then((res) => res.text());
  const root = parse(html);
  root.querySelectorAll('script').forEach((a) => a.remove());

  root.querySelectorAll('tg-emoji').forEach((a) => (a.tagName = 'b'));
  root
    .querySelector('body')
    .append(
      parse(
        '<script>\n' +
          '    function sendHeight() {\n' +
          '      const height = document.body.scrollHeight;\n' +
          "      parent.postMessage({ type: 'resize', height }, '*');\n" +
          '    }\n' +
          "    window.addEventListener('load', sendHeight);\n" +
          "    window.addEventListener('resize', sendHeight);\n" +
          '    const observer = new MutationObserver(sendHeight);\n' +
          '    observer.observe(document.body, { childList: true, subtree: true, attributes: true });\n' +
          '  </script>',
      ),
    );

  return root.outerHTML;
});

server.listen({ port: 1488, host: '0.0.0.0' });

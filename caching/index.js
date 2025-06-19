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

const server = fastify({ logger: true });

const cache = new NodeCache({ stdTTL: 86400, checkperiod: 120 });

//steam

server.get('/getPlayerUsername', async (request) => {
  const key = request.query.key;
  const id = request.query.id;
  if (!cache.has(`username-${id}`)) {
    cache.set(`username-${id}`, await getPlayerUsername(key, id));
  }
  return { username: cache.get(`username-${id}`) };
});

server.get('/getOwnedGames', async (request) => {
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

server.get('/getGameHours', async (request) => {
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

server.get('/getSteamLevel', async (request) => {
  const key = request.query.key;
  const id = request.query.id;
  if (!cache.has(`level-${id}`)) {
    cache.set(`level-${id}`, await getSteamLevel(key, id));
  }
  return { level: cache.get(`level-${id}`) };
});

server.get('/resolveVanityUrl', async (request) => {
  const key = request.query.key;
  const vanityurl = request.query.vanityurl;
  if (!cache.has(`id-${vanityurl}`)) {
    cache.set(`id-${vanityurl}`, await resolveVanityUrl(key, vanityurl));
  }
  return { id: cache.get(`id-${vanityurl}`) };
});

//youtube

server.get('/youtube/getChannelInfo', async (request) => {
  const key = request.query.key;
  const handle = request.query.handle ?? request.query.id;
  const isId = !request.query.handle;
  if (!cache.has(`yt-${handle}`)) {
    cache.set(`yt-${handle}`, await getYoutubeChannelInfo(key, handle, isId));
  }
  return { info: cache.get(`yt-${handle}`) };
});

//twitch

server.get('/twitch/getChannelInfo', async (request) => {
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
});

//vk

server.get('/vk/getPageInfo', async (request) => {
  const key = request.query.key;
  const handle = request.query.handle;
  if (!cache.has(`vk-${handle}`)) {
    cache.set(`vk-${handle}`, await getPageInfo(key, handle));
  }
  return { info: cache.get(`vk-${handle}`) };
});

server.listen({ port: 1488, host: '0.0.0.0' });

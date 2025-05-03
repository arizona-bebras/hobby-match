import { STEAM_API_KEY as key } from '$env/static/private';
import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ({ request }) => {
  const requestData = await request.json();
  const steamId = getSteamID(key, requestData.link);
  console.log(steamId)
  const data = await fetch(
    `https://api.steampowered.com/IPlayerService/GetSteamLevel/v1/?key=${key}&steamid=${steamId}`,
  ).then((result) => result.json());
  const lvl = data.player_level;
  console.log(lvl);
  return new Response(
    JSON.stringify({
      success: true,
      player_level: parseInt(lvl),
    }),
    {
      headers: { 'Content-Type': 'application/json' },
    },
  );
};

async function someFunc(url: string) {
  const userProfileURL = url;
  let steamID;
  const name = userProfileURL.split('/')[4];

  if (isFinite(name)) {
    steamID = name;
  } else {
    steamID = await getSteamID(key, name);
    steamID = steamID.steamid;
  }

  const games = await getGames(key, steamID);
  const gameList = [];
  for (const game of games) {
    game.img_icon_url = `https://media.steampowered.com/steamcommunity/public/images/apps/${game.appid}/${game.img_icon_url}.jpg`;
    // console.log(game.appid, game.name, Math.trunc(game.playtime_forever/60));
    gameList.push(game);
  }
  return gameList;
}

async function getSteamID(key: string, name: string) {
  const response = await fetch(
    `http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=${key}&vanityurl=${name}`,
  );
  const json = await response.json();
  return json.response;
}

async function getGames(key: string, steamID: string) {
  const response = await fetch(
    `https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=${key}&steamid=${steamID}&include_appinfo=true&include_played_free_games=true&format=json`,
  );
  const json = await response.json();
  return json.response.games;
}

// console.log(/steamcommunity\.com\/id\/(.+)\//.exec('https://steamcommunity.com/id/xrystikonelove/dasdasda')[1])

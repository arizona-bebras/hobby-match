import { STEAM_API_KEY as key } from '$env/static/private';
import { json, type RequestHandler } from '@sveltejs/kit';
type Game = {
  appid: string;
  name: string;
  playtime_forever: string;
  img_icon_url: string;
};

export const GET: RequestHandler = async ({ request, url }) => {
  const userGames = await getUserGames(url.searchParams.get('link') as string);
  // const userGames = await getUserGames(url.searchParams.get('link'));
  return json(userGames);
  // const steamId = await getSteamID(key, requestData.link);
  // console.log(steamId);
  // //const lvldata = await fetch(
  // //  `https://api.steampowered.com/IPlayerService/GetSteamLevel/v1/?key=${key}&steamid=${steamId}`,
  // //).then((result) => result.json());
  // const steamUserData = await fetch(
  //   `https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=${key}&steamids=${steamId}`,
  // ).then((result) => result.json());
  // const username = steamUserData.personaname;
  // //const lvl = lvldata.player_level;
  // const lvl = '11';
  // console.log(steamUserData.response.players);
  // return new Response(
  //   JSON.stringify({
  //     success: true,
  //     player_name: username,
  //     player_level: parseInt(lvl),
  //   }),
  //   {
  //     headers: { 'Content-Type': 'application/json' },
  //   },
  // );
};

async function getUserGames(url: string) {
  let steamID;
  const steamRegex =
    /^(?:https:\/\/)?steamcommunity\.com\/((?:id)|(?:profiles))\/(\w+)/gm;
  const match = steamRegex.exec(url);

  if (match![1] === 'profiles') {
    steamID = match![2];
  } else if (match![1] === 'id') {
    steamID = await getSteamID(key, match![2]);
    steamID = steamID.steamid;
  }

  const games = await getGames(key, steamID);
  const gameList: Game[] = [];
  for (const game of games) {
    // console.log(game.appid, game.name, Math.trunc(game.playtime_forever/60));
    game.img_icon_url = `https://media.steampowered.com/steamcommunity/public/images/apps/${game.appid}/${game.img_icon_url}.jpg`;
    gameList.push({
      appid: game.appid,
      name: game.name,
      playtime_forever: game.playtime_forever,
      img_icon_url: game.img_icon_url,
    });
  }
  return gameList;
}

async function getSteamID(key: string, name: string) {
  const response = await fetch(
    `https://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=${key}&vanityurl=${name}`,
  );
  const json = await response.json();
  return json.response;
}

async function getGames(key: string, steamID: string) {
  const response = await fetch(
    `https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=${key}&steamid=${steamID}&include_appinfo=true&include_played_free_games=true&format=json`,
  );
  const json = await response.text();
  console.log(json);
  return JSON.parse(json).response.games;
}

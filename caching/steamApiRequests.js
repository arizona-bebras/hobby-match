export async function getPlayerUsername(key, id) {
  let response = await fetch(
    `http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=${key}&steamids=${id}`,
  ).then((result) => result.json());
  return response.response.players[0].personaname;
}

export async function getSteamLevel(key, id) {
  let response = await fetch(
    `http://api.steampowered.com/IPlayerService/GetSteamLevel/v1/?key=${key}&steamid=${id}`,
  ).then((response) => response.json());
  return response.response.player_level;
}

export async function resolveVanityUrl(key, vanityurl) {
  let response = await fetch(
    `http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=${key}&vanityurl=${vanityurl}`,
  ).then((response) => response.json());
  return response.response.steamid;
}

export async function getOwnedGames(key, id) {
  let response = await fetch(
    `http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=${key}&steamid=${id}&include_appinfo=true&include_played_free_games=true&format=json`,
  ).then((response) => response.json());
  return response.response.games;
}

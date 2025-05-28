/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

function resolveSteamLink(link) {
  const handle = /^(?:https?:\/\/)?steamcommunity\.com\/(profiles|id)\/([a-zA-Z0-9]+)/.exec(link);
  if (!link || !handle || handle.length !== 3) return null;
  if (handle[1] === 'profiles') {
    return handle[2];
  } else if (handle[1] === 'id') {
    const vanity = $http.send({
      method: "GET",
      url: `https://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=${$os.getenv("STEAMAPI_KEY")}&vanityurl=${handle[2]}`,
    });
    if(!vanity.json.response['steamid']) return null;
    return vanity.json.response['steamid']
  } else {
    return null;
  }
}

module.exports = {
  getUserInfo: (link) => {
    const id = resolveSteamLink(link);
    const level = $http.send({
      method: "GET",
      url: `https://api.steampowered.com/IPlayerService/GetSteamLevel/v1/?key=${$os.getenv("STEAMAPI_KEY")}&steamid=${id}`,
    });
    if (level.json.response.player_level === undefined) return null;
    const player = $http.send({
      method: "GET",
      url: `https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=${$os.getenv("STEAMAPI_KEY")}&steamids=${id}`,
    });
    if (!player.json.response || !player.json.response.players || player.json.response.players.length !== 1) return null;

    return {
      level: level.json.response.player_level,
      username: player.json.response.players[0]['personaname'],
    };
  },
};

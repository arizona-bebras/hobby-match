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
      url: `${$os.getenv("CACHED_STEAM_API")}/resolveVanityUrl?key=${$os.getenv("STEAM_API_KEY")}&vanityurl=${handle[2]}`,
    });
    if(!vanity.json.id) return null;
    return vanity.json.id
  } else {
    return null;
  }
}

module.exports = {
  getUserInfo: (link) => {
    const id = resolveSteamLink(link);
    const level = $http.send({
      method: "GET",
      url: `${$os.getenv("CACHED_STEAM_API")}/getSteamLevel?key=${$os.getenv("STEAM_API_KEY")}&id=${id}`,
    });
    if (level.json.level === undefined) return null;
    const player = $http.send({
      method: "GET",
      url: `${$os.getenv("CACHED_STEAM_API")}/getPlayerUsername?key=${$os.getenv("STEAM_API_KEY")}&id=${id}`,
    });
    if (!player.json.username) return null;

    return {
      level: level.json.level,
      username: player.json.username,
    };
  },
  getGameInfo: (link, appid) => {
    const id = resolveSteamLink(link);
    const gameData = $http.send({
      method: "GET",
      url: `${$os.getenv("CACHED_STEAM_API")}/getGameHours?key=${$os.getenv("STEAM_API_KEY")}&id=${id}&appid=${appid}`,
    })
    if (!gameData.json?.game?.game_name) return null;
    return {
      icon: `https://media.steampowered.com/steamcommunity/public/images/apps/${appid}/${gameData.json.game.icon}.jpg`,
      title: gameData.json.game.game_name,
      hours_played: gameData.json.game.hours,
    }
  }
};

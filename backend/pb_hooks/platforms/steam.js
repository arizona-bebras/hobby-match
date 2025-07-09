/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

function resolveSteamLink(link) {
  const handle = /^(?:https?:\/\/)?steamcommunity\.com\/(profiles|id)\/([a-zA-Z0-9_.-]+)/.exec(link);
  if (!link || !handle || handle.length !== 3) return null;
  if (handle[1] === 'profiles') {
    return handle[2];
  } else if (handle[1] === 'id') {
    const vanity = require(`${__hooks}/steamapi.js`).request(
      'GET',
      'vanity',
      `vanityurl=${handle[2]}`,
    );
    if(!vanity.json.id) return null;
    return vanity.json.id
  } else {
    return null;
  }
}

module.exports = {
  getUserInfo: (link) => {
    const id = resolveSteamLink(link);
    const level = require(`${__hooks}/steamapi.js`).request(
      'GET',
      'level',
      `id=${id}`,
    );
    if (level.json.level === undefined) return null;
    const player = require(`${__hooks}/steamapi.js`).request(
      'GET',
      'username',
      `id=${id}`,
    );
    if (!player.json.username) return null;
    return {
      level: level.json.level,
      username: player.json.username,
    };
  },
  getGameInfo: (link, appid) => {
    const id = resolveSteamLink(link);
    const gameData = require(`${__hooks}/steamapi.js`).request(
      'GET',
      'game',
      `id=${id}&appid=${appid}`,
    );
    if (!gameData.json?.game?.game_name) return null;
    return {
      icon: `https://media.steampowered.com/steamcommunity/public/images/apps/${appid}/${gameData.json.game.icon}.jpg`,
      title: gameData.json.game.game_name,
      hours_played: gameData.json.game.hours,
    }
  }
};

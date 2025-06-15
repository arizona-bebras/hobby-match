/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

module.exports = {
  getChannelInfo: (link) => {
    const handle = /^(?:https?:\/\/)?(?:www\.)?twitch\.tv\/([\w-]+)/.exec(link);
    const followers = require(`${__hooks}/twitchapi.js`).request(
      'GET',
      `handle=${handle[1]}`,
    );
    return {
      title: handle[1],
      followers: followers.json.followers,
    };
  },
};

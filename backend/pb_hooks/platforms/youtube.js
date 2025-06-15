/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

module.exports = {
  getChannelInfo: (link) => {
    const handle = /^(?:https?:\/\/)?(?:www\.)?youtube\.com\/([\w@-]+)/.exec(link);
    if (!link || !handle || handle.length !== 2) return null;
    const response = require(`${__hooks}/youtubeapi.js`).request(
      'GET',
      `handle=${handle[1]}`,
    );
    return {
      title: handle[1],
      subscribers: response.json.subs,
    };
  },
};

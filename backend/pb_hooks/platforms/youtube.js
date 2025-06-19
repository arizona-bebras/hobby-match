/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

module.exports = {
  getChannelInfo: (link) => {
    const handle = /^https?:\/\/(?:www\.)?youtube\.com\/(channel|user|c|@)\/?([a-zA-Z0-9_-]+)/.exec(link);
    if (!link || !handle || handle.length !== 3) return null;
    let data;
    if (handle[1] === '@') {
      data = `handle=${handle[2]}`;
    } else {
      data = `id=${handle[2]}`;
    }
    const response = require(`${__hooks}/youtubeapi.js`).request('GET', data);
    console.log(response.json)
    return response.json.info;
  },
};

/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

module.exports = {
  getChannelInfo: (link) => {
    const handle = /^(?:https?:\/\/)?(?:www\.)?youtube\.com\/([\w@-]+)/.exec(link);
    if (!link || !handle || handle.length !== 2) return null;
    const response = $http.send({
      method: "GET",
      url: `https://www.googleapis.com/youtube/v3/channels?part=snippet,statistics&forHandle=${handle[1]}&key=${$os.getenv("GOOGLEAPI_TOKEN")}`,
    });
    if (!response.json.items || response.json.items.length !== 1) return null;

    return {
      title: response.json.items[0].snippet.title,
      subscribers: response.json.items[0].statistics.subscriberCount,
    };
  },
};

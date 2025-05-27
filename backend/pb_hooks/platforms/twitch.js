/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

module.exports = {
  getChannelInfo: (link) => {
    const handle = /^(?:https?:\/\/)?(?:www\.)?twitch\.tv\/([\w-]+)/.exec(link);
    if (!link || !handle || handle.length !== 2) return null;
    const headers = {
      'Client-Id': $os.getenv("TWITCH_CLIENT_ID"),
      'Authorization': `Bearer ${$os.getenv("TWITCH_TOKEN")}`,
    };
    const info = $http.send({
      method: "GET",
      url: `https://api.twitch.tv/helix/users?login=${handle[1]}`,
      headers,
    });
    if (!info.json.data || info.json.data.length !== 1) return null;
    const id = info.json.data[0].id;
    const title = info.json.data[0]['display_name'];

    console.log(JSON.stringify(info.json));
    const followers = $http.send({
      method: "GET",
      url: `https://api.twitch.tv/helix/channels/followers?broadcaster_id=${id}`,
      headers,
    });

    return {
      title,
      followers: followers.json.total,
    };
  },
};

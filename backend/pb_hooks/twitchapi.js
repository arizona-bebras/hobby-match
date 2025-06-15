/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

module.exports = {
  request: function (method, data) {
    const baseUrl = `${$os.getenv('CACHED_API')}/twitch/getChannelInfo?key=${$os.getenv('TWITCH_TOKEN')}&clientId=${$os.getenv('TWITCH_CLIENT_ID')}&${data}`;
    return $http.send({
      method,
      url: baseUrl,
    });
  },
};
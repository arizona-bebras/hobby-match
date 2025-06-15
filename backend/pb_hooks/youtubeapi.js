/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

module.exports = {
  request: function (method, data) {
    const baseUrl = `${$os.getenv('CACHED_API')}/youtube/getChannelInfo?key=${$os.getenv('YOUTUBE_API_KEY')}&${data}`;
    return $http.send({
      method,
      url: baseUrl,
    });
  },
};
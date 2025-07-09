/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

module.exports = {
  request: function (method, data) {
    const baseUrl = `${$os.getenv('CACHED_API')}/youtube/channel?key=${$os.getenv('YOUTUBE_API_KEY')}&${data}`;
    return $http.send({
      method,
      url: baseUrl,
      headers: { 'Authorization': `Bearer ${$os.getenv('CACHED_TOKEN')}` }
    });
  },
};

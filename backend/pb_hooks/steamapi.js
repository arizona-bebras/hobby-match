/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

module.exports = {
  request: function (method, endpoint, data) {
    const baseUrl = `${$os.getenv('CACHED_STEAM_API')}/${endpoint}?key=${$os.getenv('STEAM_API_KEY')}&${data}`;
    console.log('СРАБОТАЛО', baseUrl);
    return $http.send({
      method,
      url: baseUrl,
    });
  },
};
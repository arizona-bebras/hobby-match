/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

module.exports = {
  request: function (method, endpoint, data) {
    const baseUrl = `${$os.getenv('CACHED_API')}/steam/${endpoint}?key=${$os.getenv('STEAM_API_KEY')}&${data}`;
    return $http.send({
      method,
      url: baseUrl,
      headers: { 'Authorization': `Bearer ${$os.getenv('CACHED_TOKEN')}` }
    });
  },
};

/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

module.exports = {
  request: function (method, data) {
    const baseUrl = `${$os.getenv('CACHED_API')}/vk/getPageInfo?key=${$os.getenv('VK_ACCESS_TOKEN')}&${data}`;
    return $http.send({
      method,
      url: baseUrl,
    });
  },
};
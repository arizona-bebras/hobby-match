/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

module.exports = {
  request: function (method, endpoint, data) {
    return $http.send({
      method,
      url: `${$os.getenv('WORKER_ENDPOINT')}/${endpoint}`,
      body: JSON.stringify(data),
      headers: {
        Authorization: `Bearer ${$os.getenv('WORKER_SECRET')}`,
      },
    });
  },
};

/// <reference types="../pb_data/types.d.ts" />

module.exports = {
  getUserInfo: (link) => {
    const handle = /vk\.com\/((?=id)\d+|\w+)/i.exec(link);
    console.log(handle)
    const response = require(`${__hooks}/vkapi.js`).request(
      'GET',
      `handle=${handle[1]}`,
    );
    return {
      name: response.json.info.name,
      followers: response.json.info.followers,
    };
  },
};
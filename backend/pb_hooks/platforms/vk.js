/// <reference types="../pb_data/types.d.ts" />

module.exports = {
  getUserInfo: (link) => {
    const handle = /vk\.com\/((?=id)\d+|\w+)/i.exec(link);
    console.log('-----VK-----');
    console.log(handle);
    //if (!link || !handle || handle.length !== 2) return null;
    const response = $http.send({
      method: "GET",
      url: `https://api.vk.ru/method/users.get?user_ids=${handle[1]}&fields=followers_count&access_token=${$os.getenv("VK_ACCESS_TOKEN")}&v=5.199`,
    });
    console.log(JSON.stringify(response));
    console.log('-----VK-----');
    if (!response.json.response || response.json.response.length !== 1) return null;
    return {
      name: response.json.response[0].first_name + ' ' + response.json.response[0].last_name,
      followers: response.json.response[0].followers_count,
    };
  },
};
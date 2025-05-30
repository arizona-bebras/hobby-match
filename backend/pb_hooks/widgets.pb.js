/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

console.log("Loading widget hooks!");

onRecordEnrich((e) => {
  e.record.withCustomData(true);
  const data = JSON.parse(e.record.get('data'));

  switch (data.type) {
    case 'photo':
      e.record.set("additionalData", {
        type: 'photo',
        urls: e.record.get('files').map(file => e.record.baseFilesPath() + "/" + file)
      })
      break;
    case 'social_media':
      switch (data.platform) {
        case 'YouTube':
          e.record.set("additionalData", {
            type: data.platform,
            ...require(`${__hooks}/platforms/youtube.js`).getChannelInfo(data.link)
          });
          break;
        case 'Twitch':
          e.record.set("additionalData", {
            type: data.platform,
            ...require(`${__hooks}/platforms/twitch.js`).getChannelInfo(data.link)
          });
          break;
        case 'Steam':
          e.record.set("additionalData", {
            type: data.platform,
            ...require(`${__hooks}/platforms/steam.js`).getUserInfo(data.link)
          });
          break;
        case 'VK':
          e.record.set("additionalData", {
            type: data.platform,
            followers: 1337,
            // ...require(`${__hooks}/platforms/vk.js`).getUserInfo(data.link)
          });
          break;
        case 'X':
          e.record.set("additionalData", {
            type: data.platform,
            followers: 1337,
            tweets: 1234,
            // ...require(`${__hooks}/platforms/x.js`).getUserInfo(data.link)
          });
          break;
        default:
          e.record.set("additionalData", {});
          break;
      }
      break;
      case 'survey':
        let surveyData = {};
        for (let option of data.options) {
          let exp = $dbx.hashExp({
            "survey": e.record.getString("id"),
            "selected_option": option.description
          });
          surveyData[option.description] = $app
            .findAllRecords("votes", exp)
            .map(record => record.getString("user"));
        }
        e.record.set("additionalData", {
          type: 'survey',
          stats: surveyData
        })
        break;
  }

  console.log(JSON.stringify(e.record));

  e.next();
}, 'widgets');

onRecordUpdate((e) => {
  console.log(JSON.stringify(e, undefined, 2));
  const url = e.record.getString('user_photo')
  if (!url.startsWith('http')) return e.next();

  const res = $http.send({
    method:  "GET",
    url,
  });

  const formats = {
    'image/jpeg': 'jpg',
    'image/png': 'png',
    'image/svg+xml': 'svg',
    'image/gif': 'gif',
    'image/webp': 'webp',
  };

  const mime = res.headers['Content-Type'];
  console.log(mime);
  if(!(mime in formats)) return null;

  const file = $filesystem.fileFromBytes(res.body, `avatar.${formats[mime]}`);

  e.record.set('user_photo', file);
  e.next();
}, 'users');

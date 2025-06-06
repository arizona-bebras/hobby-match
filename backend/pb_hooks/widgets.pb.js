/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

console.log("Loading widget hooks!");

onRecordEnrich((e) => {
  if(e.record.getDateTime("birth_date")) {
    e.record.withCustomData(true);
    const currentDate = new Date().getTime() / 1000;
    const birthDate = e.record.getDateTime("birth_date").unix();
    const age = currentDate - birthDate;
    e.record.set('age', Math.floor(age / (60 * 60 * 24 * 365)));
  }
  e.next();
}, 'users')

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
        let surveyData = [];
        let myVote = undefined;
        for (let i = 0; i < data.options.length; i++) {
          let exp = $dbx.hashExp({
            "survey": e.record.getString("id"),
            "selected_option": i
          });
          surveyData.push($app.countRecords("votes", exp));

          if ($app.countRecords("votes", exp, $dbx.hashExp({
            "user": e.requestInfo.auth.id,
          })) > 0)
            myVote = i;
        }
        e.record.set("additionalData", {
          type: 'survey',
          stats: surveyData,
          myVote
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

onRecordUpdate((e) => {
  const oldOrder = $app.unsafeWithoutHooks().findRecordById('widgets', e.record.id).getInt('order');

  if (oldOrder !== e.record.getInt('order')) {
    const [conflict] = $app.findRecordsByFilter('widgets',
      'user = {:user} && order = {:order} && id != {:id}', 'order', 1, 0, {
        "user": e.record.getString("user"),
        "order": e.record.getInt("order"),
        "id": e.record.getInt("id"),
      });

    console.log("conflict", oldOrder, e.record.getInt('order'), JSON.stringify(conflict, null, 2));
    if (conflict) {
      conflict.set("order", oldOrder);
      $app.unsafeWithoutHooks().save(conflict);
    }
  }
  e.next()

  let userRecords = $app.unsafeWithoutHooks().findRecordsByFilter('widgets',
    'user = {:user}', 'order', 0, 0, {
      "user": e.record.getString("user")
    });

  for (let i = 0; i < userRecords.length; i++) {
    let record = userRecords[i];
    if (record.getInt("order") !== i) {
      record.set("order", i);
      $app.unsafeWithoutHooks().save(record);
    }
  }
}, 'widgets')

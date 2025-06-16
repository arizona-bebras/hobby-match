/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

console.log("Loading worker!");

routerAdd("GET", "/worker/autocomplete", (e) => {
  const query = e.requestInfo().query['query'];
  const response = require(`${__hooks}/worker.js`).request("POST", "interests/query", { query });

  return e.json(response.statusCode, { query, response: response.json })
}, $apis.requireAuth());

/** @param {core.RecordEvent} e */
function upsertInterests(e) {
  const response = require(`${__hooks}/worker.js`).request("POST", "interests/upsert", {
    id: e.record.id,
    tag: e.record.getString("tag"),
  });

  if(response.statusCode !== 200)
    return console.error("Failed to create worker vector!", JSON.stringify(response.json));

  e.next()
}

onRecordDelete((e) => {
  const response = require(`${__hooks}/worker.js`).request("DELETE", "interests/delete", {
    ids: [e.record.id]
  });

  if(response.statusCode !== 200)
    return console.error("Failed to delete worker vector!", JSON.stringify(response.json));

  e.next()
}, "interests");

onRecordCreate(upsertInterests, "interests");
onRecordUpdate(upsertInterests, "interests");

/** @param {core.RecordEvent} e */
function upsertUser(e) {
  e.next()
  const request = {
    id: e.record.id,
    interest_ids: e.record.getRaw("interests"),
    text: e.record.getString("user_info"),
  };
  const response = require(`${__hooks}/worker.js`).request("POST", "feed/upsert", request);

  if(response.statusCode !== 200)
    return console.error("Failed to create worker vector!", JSON.stringify(response.json));
}

routerAdd("GET", "/worker/feed", (e) => {
  const views = arrayOf(new DynamicModel({
    "pageOwner": "",
    "view_count": 0,
  }))
  $app.db().newQuery(`
WITH view_counts AS (
  SELECT
    u.id AS pageOwner,
    COALESCE(COUNT(v.viewer), 0) AS view_count
  FROM users u
  LEFT JOIN views v ON v.pageOwner = u.id AND v.viewer = {:current_user_id}
  WHERE u.id != {:current_user_id}
  GROUP BY u.id
),
ranked_users AS (
  SELECT
      pageOwner,
    view_count,
    RANK() OVER (ORDER BY view_count ASC) AS view_rank
  FROM view_counts
)
SELECT
  ru.pageOwner,
  ru.view_count
FROM ranked_users ru
WHERE ru.view_rank <= (
  SELECT MIN(view_rank)
  FROM (
    SELECT view_rank
    FROM ranked_users
    GROUP BY view_rank
    HAVING COUNT(*) >= 3
    ORDER BY view_rank
    LIMIT 1
  ) AS sub
)
LIMIT 100;
`).bind({"current_user_id": e.auth.id}).all(views);

  const response = require(`${__hooks}/worker.js`).request("POST", "feed/query", {
    id: e.auth.id,
    from: views.map(view => view['pageOwner']),
  });
  if(response.statusCode !== 200) return e.json(response.statusCode, { response: response.json });
  const matches = response.json?.matches;
  if(!matches) return e.json(500, {error: 'failed to get feed'});
  console.log(JSON.stringify(matches));

  let collection = $app.findCollectionByNameOrId("views");
  for (const match of matches) {
    let record = new Record(collection);
    record.set('viewer', e.auth.id);
    record.set('pageOwner', match.id);
    $app.save(record)
  }

  const pages = $app.findRecordsByIds('users', matches.map(m => m.id));
  $app.expandRecords(pages, ["interests"], null);
  $apis.enrichRecords(e, pages);
  return e.json(200, pages.map(r => {
    const widgets = $app.findRecordsByFilter('widgets', 'user = {:id}', 'order', 0, 0, {id: r.id});
    $apis.enrichRecords(e, widgets);
    return {
      id: r.id,
      miniapp_name: r.getString('miniapp_name'),
      age: r.getInt('age'),
      gender: r.getString('gender'),
      location: r.getString('location'),
      user_photo: r.getString('user_photo'),
      user_info: r.getString('user_info'),
      interests: r.expandedAll("interests"),
      widgets
    };
  }));
}, $apis.requireAuth());

onRecordDelete((e) => {
  const response = require(`${__hooks}/worker.js`).request("DELETE", "feed/delete", {
    id: e.record.id
  });

  if(response.statusCode !== 200)
    return console.error("Failed to delete worker vector!", JSON.stringify(response.json));

  e.next()
}, "users");
onRecordCreate(upsertUser, "users");
onRecordUpdate(upsertUser, "users");

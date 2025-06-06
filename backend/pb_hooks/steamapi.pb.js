routerAdd("GET", "/steam/username", (e) => {
  const id = e.requestInfo().query['id']; 
  const response = require(`${__hooks}/steamapi.js`).request("GET", "getPlayerUsername", `id=${id}` )

  return e.json(response.statusCode, { response: response.json })
}, $apis.requireAuth())

routerAdd("GET", "/steam/games", (e) => {
  const id = e.requestInfo().query['id']; 
  const response = require(`${__hooks}/steamapi.js`).request("GET", "getOwnedGames", `id=${id}` )

  return e.json(response.statusCode, { response: response.json })
}, $apis.requireAuth())

routerAdd("GET", "/steam/level", (e) => {
  const id = e.requestInfo().query['id']; 
  const response = require(`${__hooks}/steamapi.js`).request("GET", "getSteamLevel", `id=${id}` )

  return e.json(response.statusCode, { response: response.json })
}, $apis.requireAuth())

routerAdd("GET", "/steam/vanityurl", (e) => {
  const vanityurl = e.requestInfo().query['vanityurl']; 
  const response = require(`${__hooks}/steamapi.js`).request("GET", "resolveVanityUrl", `vanityurl=${vanityurl}` )

  return e.json(response.statusCode, { response: response.json })
}, $apis.requireAuth())

routerAdd("GET", "/steam/hours", (e) => {
  const id = e.requestInfo().query['id']; 
  const appid = e.requestInfo().query['appid']; 
  const response = require(`${__hooks}/steamapi.js`).request("GET", "getGameHours", `id=${id}&appid=${appid}` )

  return e.json(response.statusCode, { response: response.json })
}, $apis.requireAuth())
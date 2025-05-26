/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

console.log('Loading hooks!');

routerAdd("GET", "/hello/{name}", (e) => {
  let name = e.request.pathValue("name")

  return e.json(200, { "message": "Hello " + name })
})
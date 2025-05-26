/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

console.log('Loading hooks!');

routerAdd('GET', '/hello/:name', (c) => {
  const name = c.pathParam('name');

  return c.json(200, { message: 'Hello ' + name });
});

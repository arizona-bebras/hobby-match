/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

console.log("Loading widget hooks!");

onRecordEnrich((e) => {
  e.record.withCustomData(true);

  e.record.set("additionalData", { test: 123 });
  console.log(JSON.stringify(e.record));

  e.next();
}, 'widgets');

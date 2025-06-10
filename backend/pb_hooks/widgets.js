/// <reference types="../pb_data/types.d.ts" />
/* eslint-disable */

module.exports = {
  reorder: function (user) {
    let userRecords = $app
      .unsafeWithoutHooks()
      .findRecordsByFilter('widgets', 'user = {:user}', 'order', 0, 0, {
        user,
      });

    for (let i = 0; i < userRecords.length; i++) {
      let record = userRecords[i];
      if (record.getInt('order') !== i) {
        record.set('order', i);
        $app.unsafeWithoutHooks().save(record);
      }
    }
  },
};

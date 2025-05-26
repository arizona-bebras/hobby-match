import { pb } from '$lib/index';
import type { PageLoad } from './$types';
import type { WidgetWithService } from '$lib/components/widgetConstructors/widgetsConstructor';
import { convertRecordToWidget } from '$lib/index';

export const load: PageLoad = async ({ data }) => {
  //depends('user:widgets');
  if (pb.authStore.isValid) {
    const { textForm } = data;
    //const telegram_id = JSON.parse(localStorage.pocketbase_auth).model
    //  .telegram_id;
    //const widgetsRecords = await pb.collection('widgets').getFullList({
      //filter: `telegram_id = "${telegram_id}"`,
    //  sort: `+order`,
    //});

    const testRecord = await pb.collection('widgets').getOne('bh4d58u66nwg72x');
    const widgetsRecords = [testRecord];
    const widgets: WidgetWithService[] = [];
    for (const record of widgetsRecords) {
      const widget: WidgetWithService = {
        widget: convertRecordToWidget(record),
        deleteStatus: false,
        changeStatus: false,
        additionalData: {},
      };
      widgets.push(widget);
    }
    console.log(widgets);
    return { widgets: structuredClone(widgets), textForm };
  }
  throw new Error('Data validation error');
};

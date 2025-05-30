import { pb } from '$lib/index';
import type { PageLoad } from './$types';
import type { WidgetWithService } from '$lib/components/widgetConstructors/widgetsConstructor';
import { convertRecordToWidget } from '$lib/index';

export const load: PageLoad = async ({ data }) => {
  // depends('user:widgets');
  if (pb.authStore.isValid) {
    const { textForm } = data;
    //
    // const collection = await pb.collections.getOne('widgets');
    // console.log(collection.schema);
    console.log(pb.collection('widgets').getOne('20h66bx83j0s3ok'));
    const widgetsRecords = await pb.collection('widgets').getFullList({
      sort: `+order`,
    });

    const widgets: WidgetWithService[] = [];
    for (const record of widgetsRecords) {
      const widget: WidgetWithService = {
        // @ts-expect-error because
        widget: record,
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

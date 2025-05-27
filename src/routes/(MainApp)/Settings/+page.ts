import { pb } from '$lib/index';
import type { PageLoad } from './$types';
import { convertRecordToWidget, getSocialMediaData } from '$lib/index';
import type { WidgetWithService } from '$lib/components/widgetConstructors/widgetsConstructor';

export const load: PageLoad = async () => {
  console.log(localStorage);
  const telegram_id = JSON.parse(localStorage.pocketbase_auth).model
    .telegram_id;
  const widgetsRecords = await pb.collection('widgets').getFullList({
    filter: `telegram_id = "${telegram_id}"`,
    sort: `+order`,
  });
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
  return { widgets: structuredClone(widgets) };
};

import { pb } from '$lib/index';
import type { PageLoad } from './$types';
import type { Widget } from '$lib/widgetTypes/widgetTypes';
import { converRecordToWidget, getSocialMediaData } from '$lib/index';

interface AdditionalData {
  socialMeidaData?: number;
}

interface WidgetWithService {
  widget: Widget;
  deleteStatus: boolean;
  changeStatus: boolean;
  additionalData: AdditionalData;
}

export const load: PageLoad = async () => {
  const telegram_id = JSON.parse(localStorage.pocketbase_auth).model
    .telegram_id;
  const widgetsRecords = await pb.collection('widgets').getFullList({
    filter: `telegram_id = "${telegram_id}"`,
    sort: `+order`,
  });
  const widgets: WidgetWithService[] = [];
  for (const record of widgetsRecords) {
    const widget: WidgetWithService = {
      widget: converRecordToWidget(record),
      deleteStatus: false,
      changeStatus: false,
      additionalData: {},
    };
    if (widget.widget.data.type == 'social_media')
      widget.additionalData.socialMeidaData = await getSocialMediaData(
        widget.widget,
      );
    widgets.push(widget);
  }
  return { widgets: structuredClone(widgets) };
};

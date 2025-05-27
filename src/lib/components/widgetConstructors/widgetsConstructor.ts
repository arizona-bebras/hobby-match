import type { Widget } from '$lib/widgetTypes/widgetTypes';
import { pb } from '$lib/index';

export interface AdditionalData {
  socialMediaData?: number;
}

export interface WidgetWithService {
  widget: Widget;
  deleteStatus: boolean;
  changeStatus: boolean;
  additionalData: AdditionalData;
}

export async function createWidget(
  formData: Widget['data'],
  numberOfWidgets: number,
  files: File[] = [],
): Promise<void> {
  await pb.collection('widgets').create({
    user: pb.authStore.model?.id,
    order: numberOfWidgets,
    files: files,
    data: formData,
  });
}

export async function updateWidgetsOrder(widgets: WidgetWithService[]) {
  for (let i = 1; i <= widgets.length; i++) {
    await pb.collection('widgets').update(widgets[i - 1].widget.id, {
      order: i,
    });
    console.log(widgets[i - 1].widget.id);
  }
}

export async function deleteWidget(widgetId: string): Promise<void> {
  await pb.collection('widgets').delete(widgetId);
  /*
  widget.deleteStatus = true;
  const filtered = [];
  for (let i = 0; i < widgets.length; i++) {
    if (!widgets[i].deleteStatus) {
      filtered.push(widgets[i]);
    }
  }
  widgets = filtered;
  updateWidgetsOrder(widgets);
  */
}

export async function updateWidget(
  widget: WidgetWithService,
  formData: Widget,
) {
  await pb.collection('widgets').update(widget.widget.id, {
    data: formData,
  });
  widget.changeStatus = false;
}

export async function changeWidgetPosition(
  widgets: WidgetWithService[],
  widget: WidgetWithService,
  posChange: 1 | -1,
): Promise<void> {
  const record = await pb.collection('widgets').getFullList({
    filter: `telegram_id = "${pb.authStore.model?.telegram_id}" && order = "${widget.widget.order + posChange}"`,
  });
  //.then((record) => record);

  const widgetOrder = widget.widget.order;
  const recordOrder = record[0].order;
  await pb.collection('widgets').update(record[0].id, {
    order: recordOrder + posChange * -1,
  });
  await pb.collection('widgets').update(widget.widget.id, {
    order: widgetOrder + posChange,
  });
  const temp = widgets[recordOrder - 1];
  widgets[recordOrder - 1].widget.order += posChange * -1;
  widgets[recordOrder - 1] = widgets[widgetOrder - 1];
  widgets[widgetOrder - 1].widget.order += posChange;
  widgets[widgetOrder - 1] = temp;
}

export async function chooseOption(survey: WidgetWithService, option: string) : Promise<void> {
  await pb.collection('votes').create({
    user: pb.authStore.model?.id,
    survey: survey.widget.id,
    option: option
  });
}

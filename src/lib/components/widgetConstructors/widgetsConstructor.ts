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
    //telegram_id: pb.authStore.model?.telegram_id,
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

export async function updateWidget(widgetId: string, formData: Widget['data']) {
  await pb.collection('widgets').update(widgetId, {
    data: formData,
  });
  //widget.changeStatus = false;
}

export async function changeWidgetPosition(
  widget: Widget,
  posChange: 1 | -1,
): Promise<void> {
  await pb.collection('widgets').update(widget.id, {
    order: widget.order + posChange,
  });
}

export async function chooseOption(
  surveyId: string,
  option: number,
): Promise<void> {
  await pb.collection('votes').create({
    user: pb.authStore.record!.id,
    survey: surveyId,
    selected_option: option,
  });
}

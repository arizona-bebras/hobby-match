import type { Widget } from '$lib/widgetTypes/widgetTypes';
import { pb } from '$lib';

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
  files: File[] = [],
): Promise<void> {
  await pb.collection('widgets').create({
    user: pb.authStore.record?.id,
    order: -1,
    files: files,
    data: formData,
  });
}

export async function deleteWidget(widgetId: string): Promise<void> {
  await pb.collection('widgets').delete(widgetId);
}

export async function updateWidget(
  widgetId: string,
  formData: Widget['data'],
  files: File[] = [],
) {
  console.log(files);
  await pb.collection('widgets').update(widgetId, {
    data: formData,
    'files+': files,
  });
  //widget.changeStatus = false;
}

export async function changeWidgetPosition(
  widget: Widget,
  posChange: 1 | -1,
): Promise<void> {
  await pb.collection('widgets').update(widget.id, {
    'order+': posChange,
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

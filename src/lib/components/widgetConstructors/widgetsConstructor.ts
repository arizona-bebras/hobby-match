import type { Widget } from '$lib/widgetTypes/widgetTypes';
import { pb } from '$lib';
import { db } from '$lib';
import { invalidate } from '$app/navigation';

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
  var form = new FormData()
  form.append("data", JSON.stringify(formData))
  for (let i = 0; i < files.length; i++) {
    form.append(`file_${i}`, files[i])
  }
  const authHeader: HeadersInit = new Headers()
  authHeader.set('Authorization', `Bearer ${window.localStorage.getItem("access_token")}`)
  await fetch(`${db}/api/me/widgets`, {
    method: "POST",
    headers: authHeader,
    body: form
  })
  await invalidate('user:widgets');
}

export async function deleteWidget(widgetId: string): Promise<void> {
  const authHeader: HeadersInit = new Headers()
  authHeader.set('Authorization', `Bearer ${window.localStorage.getItem("access_token")}`)
  await fetch(`${db}/api/me/widgets?widget_id=${widgetId}`, {
    method: "DELETE",
    headers: authHeader
  })
}

export async function updateWidget(
  widgetId: string,
  formData: Widget['data'],
  files: File[] = [],
) {
  var form = new FormData()
  form.append("widget_id", widgetId)
  form.append("data", JSON.stringify(formData))
  for (let i = 0; i < files.length; i++) {
    form.append(`file_${i}`, files[i])
  }
  const authHeader: HeadersInit = new Headers()
  authHeader.set('Authorization', `Bearer ${window.localStorage.getItem("access_token")}`)
  await fetch(`${db}/api/me/widgets`, {
    method: "PUT",
    headers: authHeader,
    body: form
  })
  await invalidate('user:widgets');
  //widget.changeStatus = false;
}

export async function changeWidgetPosition(
  widget: Widget,
  posChange: 1 | -1,
): Promise<void> {
  var form = new FormData()
  form.append("widget_id", widget.id)
  form.append("order", `${widget.order}`)
  form.append("order", `${posChange}`)
  const authHeader: HeadersInit = new Headers()
  authHeader.set('Authorization', `Bearer ${window.localStorage.getItem("access_token")}`)
  await fetch(`${db}/api/me/widgets/order`, {
    method: "PUT",
    headers: authHeader,
    body: form
  })
  await invalidate('user:widgets');
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

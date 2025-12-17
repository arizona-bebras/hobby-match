import { pb } from '$lib';
import { db } from '$lib';
import type { PageLoad } from './$types';
import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';

export const load: PageLoad = async ({ fetch, data, depends }) => {
  depends('user:widgets');
  const authHeader: HeadersInit = new Headers();
  authHeader.set(
    'Authorization',
    `Bearer ${window.localStorage.getItem('access_token')}`,
  );
  const res = await fetch(`${db}/api/me`, {
    method: 'GET',
    headers: authHeader,
  }).then((res) => res.json());
  const currentDate = new Date().getTime() / 1000;
  const birthDate = new Date(res.birth_date).getTime() / 1000;
  const age = currentDate - birthDate;
  const pageData: PageData = res as PageData;
  pageData.age = Math.floor(age / (60 * 60 * 24 * 365.25));
  for (const widget of pageData.widgets) {
    widget.data = JSON.parse(widget.data);
    if (widget.additionalData != '') {
      widget.additionalData = JSON.parse(widget.additionalData);
    }
  }
  pageData.widgets = pageData.widgets.sort((a, b) => {
    return a.order - b.order;
  });
  return pageData;
};

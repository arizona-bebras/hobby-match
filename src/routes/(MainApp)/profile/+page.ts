import { pb } from '$lib';
import { db } from '$lib';
import type { PageLoad } from './$types';
import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';

export const load: PageLoad = async ({ fetch, data, depends }) => {
  depends('user:widgets');
  console.log('RELOADED!!!');
  const authHeader: HeadersInit = new Headers()
  authHeader.set('Authorization', `Bearer ${window.localStorage.getItem("access_token")}`)
  const res = await fetch(`${db}/api/me`, {
    method: "GET",
    headers: authHeader
  }).then(res => res.json())
  var pageData: PageData = JSON.parse(res) as PageData
  return pageData
}

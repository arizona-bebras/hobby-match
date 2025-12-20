import { db } from '$lib/index';

export { default as Information } from './information/Information.svelte';
export { default as Photo } from './photo/Photo.svelte';
export { default as Interests } from './interests/Interests.svelte';
export { default as RegisterStages } from './RegisterStages.svelte';
export { default as Test } from './test/Test.svelte';

export type Stages = 'Информация' | 'Фото' | 'Тест' | 'Интересы';

export async function updateData(data: object) {
  const authHeader: HeadersInit = new Headers();
  authHeader.set(
    'Authorization',
    `Bearer ${window.localStorage.getItem('access_token')}`,
  );
  // if ("user_photo" in formData) {
  //     const res = await fetch(`${db}/api/me`, {
  //         method: "POST",
  //         headers: authHeader,
  //         body: formData
  //     }).then(res => res)
  //     return res.status
  // } else {
  //     const res = await fetch(`${db}/api/me`, {
  //         method: "POST",
  //         headers: authHeader,
  //         body: JSON.stringify(formData)
  //     }).then(res => res)
  //     return res.status
  // }
  const res = await fetch(`${db}/api/me`, {
    method: 'POST',
    headers: authHeader,
    body: JSON.stringify(data),
  }).then((res) => res);
  return res.status;
}

export async function updatePhoto(photo: Record<string, File | string>) {
  const authHeader: HeadersInit = new Headers();
  authHeader.set(
    'Authorization',
    `Bearer ${window.localStorage.getItem('access_token')}`,
  );
  const formData = new FormData();
  formData.append('user_photo', photo.user_photo);
  const res = await fetch(`${db}/api/me/photo`, {
    method: 'POST',
    headers: authHeader,
    body: formData,
  }).then((res) => res);
  return res.status;
}

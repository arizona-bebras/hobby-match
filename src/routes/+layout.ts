import { db } from '$lib';
import { accessToken } from '$lib/storage/accessToken.svelte';
import { goto } from '$app/navigation';
import client from '$lib/api/client';
import { userData } from '$lib/storage/userData.svelte';
import type { NamespaceMembers } from '$lib/api/utils';

export const ssr = false;

export const load = async () => {
  if (accessToken.current) {
    console.log('RETURN!');
    return;
  }
  const res = await fetch(`${db}/api/auth`, {
    method: 'POST',
    headers: {
      Authorization: `tma ${window.Telegram.WebApp.initData}`,
    },
  });

  // const response = await client.POST('/api/auth', {
  //   headers: {
  //     Authorization: `tma ${window.Telegram.WebApp.initData}`,
  //   },
  // });

  const data = await res.json();
  const user = data.response.user;
  userData.current = user;
  window.localStorage.setItem('access_token', data.response.access_token);
  window.localStorage.setItem('refresh_token', data.response.refresh_token);
  accessToken.current = data.response.access_token;
  console.log('TOKEN CHANGE!');
  // try {
  //   const ban = await pb.collection('bans').getFirstListItem('');
  //   if (ban.reason) {
  //     await goto(`/ban?reason=${encodeURIComponent(ban.reason)}`);
  //   } else {
  //     await goto(`/ban`);
  //   }
  //   return;
  // } catch (_) {
  //   // ok
  // }

  if (
    !user.miniapp_name ||
    !user.gender ||
    !user.birth_date ||
    !user.location ||
    !user.user_info ||
    !user.user_photo
    // user.interests.length < 3
  ) {
    await goto('/registration');
  } else if (window.location.pathname !== '/search') {
    await goto('/profile');
  }

  // .then(
  // })
  // .catch((error) => {
  //   console.error('Fetch-запрос завершился ошибкой (сработал catch):', error);
  // });
};

// export const handleError = Sentry.handleErrorWithSentry();

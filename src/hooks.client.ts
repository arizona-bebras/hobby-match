import * as Sentry from '@sentry/sveltekit';
import { pb } from '$lib/index';
import { browser } from '$app/environment';
import { goto } from '$app/navigation';
import { db } from '$lib/index'
import Plausible from 'plausible-tracker';

export const plausible = Plausible({
  domain: 'shumi.space',
  apiHost: 'https://look.gesti.tech',
});
// If you don't want to use Session Replay, remove the `Replay` integration,
// `replaysSessionSampleRate` and `replaysOnErrorSampleRate` options.
Sentry.init({
  dsn: 'https://36e7202aadd5eb8af5b8392eb47b6bc2@o4509595179679744.ingest.de.sentry.io/4509595182956624',
  tracesSampleRate: 1,
  replaysSessionSampleRate: 0.1,
  replaysOnErrorSampleRate: 1,
  integrations: [Sentry.replayIntegration()],
});

if (browser) {
  // if (!window.Telegram.WebApp.isVersionAtLeast('7.0')) {
  //   window.location.replace('https://t.me/detoshumibot');
  // } else {
  //   plausible.enableAutoPageviews();
  // }
  window.Telegram.WebApp.disableVerticalSwipes();
  fetch(`${db}/api/auth`, {
    headers: {
      "Authorization": `tma ${window.Telegram.WebApp.initData}`
    },
  }).then(async (res) => {
    const data = await res.json()
    const user = data.response.user
    window.localStorage.setItem("access_token", data.response.access_token)
    window.localStorage.setItem("refresh_token", data.response.refresh_token)
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
  })
  .catch((error) => {
    console.error("Fetch-запрос завершился ошибкой (сработал catch):", error);
  });;
}

// export const handleError = Sentry.handleErrorWithSentry();

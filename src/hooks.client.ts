import * as Sentry from '@sentry/sveltekit';
import { pb } from '$lib/index';
import { browser } from '$app/environment';
import { goto } from '$app/navigation';

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
  if (!window.Telegram.WebApp.isVersionAtLeast('7.0')) {
    window.location.replace('https://t.me/detoshumibot');
  }
  window.Telegram.WebApp.disableVerticalSwipes();
  pb.send('/api/collections/users/auth-with-telegram', {
    method: 'POST',
    body: {
      data: window.Telegram.WebApp.initData,
    },
  }).then((res) => {
    pb.authStore.save(res.token, res.record);
    if (
      !res.record.miniapp_name ||
      !res.record.gender ||
      !res.record.birth_date ||
      !res.record.location ||
      !res.record.user_info ||
      !res.record.user_photo ||
      res.record.interests.length < 3
    ) {
      goto('/registration');
    } else if (window.location.pathname !== '/Search') {
      goto('/Profile');
    }
  });
}
/* 
if (browser) {
  if (!pb.authStore.isValid && window.Telegram?.WebApp?.initData) {
    const authData = await pb.send(
      '/api/collections/users/auth-with-telegram',
      {
        method: 'POST',
        body: {
          data: window.Telegram.WebApp.initData,
        },
      },
    );
    pb.authStore.save(authData.token, authData.record);
  }

localStorage.setItem('tmp_userdata', JSON.stringify(pb.authStore.model));*/
export const handleError = Sentry.handleErrorWithSentry();

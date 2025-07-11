import * as Sentry from '@sentry/sveltekit';
import { pb } from '$lib/index';
import { browser } from '$app/environment';
import { goto } from '$app/navigation';
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
  if (!window.Telegram.WebApp.isVersionAtLeast('7.0')) {
    window.location.replace('https://t.me/detoshumibot');
  } else {
    plausible.enableAutoPageviews();
  }
  window.Telegram.WebApp.disableVerticalSwipes();
  pb.send('/api/collections/users/auth-with-telegram', {
    method: 'POST',
    body: {
      data: window.Telegram.WebApp.initData,
    },
  }).then(async (res) => {
    pb.authStore.save(res.token, res.record);
    try {
      const ban = await pb.collection('bans').getFirstListItem('');
      await goto(`/ban?reason=${encodeURIComponent(ban.reason)}`);
      return;
    } catch (_) {
      // ok
    }
    if (
      !res.record.miniapp_name ||
      !res.record.gender ||
      !res.record.birth_date ||
      !res.record.location ||
      !res.record.user_info ||
      !res.record.user_photo ||
      res.record.interests.length < 3
    ) {
      await goto('/registration');
    } else if (window.location.pathname !== '/Search') {
      await goto('/Profile');
    }
  });
}

export const handleError = Sentry.handleErrorWithSentry();

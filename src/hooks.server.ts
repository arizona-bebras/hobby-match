import { sequence } from '@sveltejs/kit/hooks';
import * as Sentry from '@sentry/sveltekit';
import { pb } from '$lib/index';
import { PB_ADMIN_PASSWORD, PB_ADMIN_EMAIL } from '$env/static/private';

Sentry.init({
  dsn: 'https://36e7202aadd5eb8af5b8392eb47b6bc2@o4509595179679744.ingest.de.sentry.io/4509595182956624',
  tracesSampleRate: 1,
});

async function authAsAdmin() {
  await pb
    .collection('_superusers')
    .authWithPassword(PB_ADMIN_EMAIL, PB_ADMIN_PASSWORD);
}
authAsAdmin();
export const handleError = Sentry.handleErrorWithSentry();
export const handle = sequence(Sentry.sentryHandle());

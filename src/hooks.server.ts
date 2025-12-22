import { sequence } from '@sveltejs/kit/hooks';
import * as Sentry from '@sentry/sveltekit';

Sentry.init({
  dsn: 'https://36e7202aadd5eb8af5b8392eb47b6bc2@o4509595179679744.ingest.de.sentry.io/4509595182956624',
  tracesSampleRate: 1,
  environment: import.meta.env.PROD ? 'production' : 'development',
});

export const handleError = Sentry.handleErrorWithSentry();
export const handle = sequence(Sentry.sentryHandle());

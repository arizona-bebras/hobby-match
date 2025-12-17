import { PersistedState } from 'runed';

export const accessToken = new PersistedState<string | null>(
  'accessToken',
  null,
  { storage: 'session' },
);

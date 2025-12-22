import { PersistedState } from 'runed';
import type { UserData } from '$lib/api/utils';

export const userData = new PersistedState<UserData | undefined>(
  'userData',
  undefined,
  { storage: 'session' },
);

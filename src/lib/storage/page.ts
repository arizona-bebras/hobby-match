import { PersistedState } from 'runed';

export const hideFooter = new PersistedState<boolean>('hideFooter', false, {
  storage: 'session',
});

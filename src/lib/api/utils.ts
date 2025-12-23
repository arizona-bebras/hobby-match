import type { components } from '$lib/api/types';

export type NamespaceMembers =
  components['schemas']['handlers.NamespaceMember'][];

export type UserData = components['schemas']['database.User'];

export type UserNamespaces = components['schemas']['handlers.UserNamespaces'][];

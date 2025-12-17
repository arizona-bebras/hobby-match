import type { Middleware } from 'openapi-fetch';
import { get } from 'svelte/store';
import { accessToken } from '../storage/accessToken.svelte';

export const authMiddleware: Middleware = {
  onRequest({ request }) {
    const token = accessToken.current;

    if (!token) {
      return request;
    }
    // (optional) add logic here to refresh token when it expires

    // add Authorization header to every request
    request.headers.set('Authorization', `Bearer ${token}`);
    return request;
  },
};

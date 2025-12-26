import client from '$lib/api/client';
import { userData } from '$lib/storage/userData.svelte';
import { db } from '$lib';

export async function updateUserData() {
  const response = await client.GET('/api/me');
  userData.current = response.data;
  userData.current.image = `${db}/api/files/users/${userData.current?.tg_user}/undefined`;
}

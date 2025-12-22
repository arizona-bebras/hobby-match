import client from '$lib/api/client';
import { userData } from '$lib/storage/userData.svelte';

export async function updateUserData() {
  const response = await client.GET('/api/me');
  userData.current = response.data;
}

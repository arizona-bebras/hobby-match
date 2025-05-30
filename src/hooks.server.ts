import { pb } from '$lib/index';
import { PB_ADMIN_PASSWORD, PB_ADMIN_EMAIL } from '$env/static/private';
async function authAsAdmin() {
  await pb
    .collection('_superusers')
    .authWithPassword(PB_ADMIN_EMAIL, PB_ADMIN_PASSWORD);
}
authAsAdmin();

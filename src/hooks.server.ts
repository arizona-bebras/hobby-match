import { pb } from '$lib/index';
import { PB_ADMIN_PASSWORD } from '$env/static/private';
async function authAsAdmin() {
  await pb.admins.authWithPassword('maxi.solts@gmail.com', PB_ADMIN_PASSWORD);
}
authAsAdmin();

import { pb } from '$lib/index';
import { browser } from '$app/environment';
import { goto } from '$app/navigation';
if (browser) {
  pb.send('/api/collections/users/auth-with-telegram', {
    method: 'POST',
    body: {
      data: window.Telegram.WebApp.initData,
    },
  }).then((res) => {
    pb.authStore.save(res.token, res.record);
    if (
      !pb.authStore.record?.mini_app_name ||
      !pb.authStore.record?.gender ||
      !pb.authStore.record?.birth_date ||
      !pb.authStore.record?.location ||
      !pb.authStore.record?.user_info ||
      !pb.authStore.record?.photo ||
      !pb.authStore.record?.interests
    ) {
      goto('/registration');
    } else {
      goto('/Profile');
    }
  });
}
/* 
if (browser) {
  if (!pb.authStore.isValid && window.Telegram?.WebApp?.initData) {
    const authData = await pb.send(
      '/api/collections/users/auth-with-telegram',
      {
        method: 'POST',
        body: {
          data: window.Telegram.WebApp.initData,
        },
      },
    );
    pb.authStore.save(authData.token, authData.record);
  }

localStorage.setItem('tmp_userdata', JSON.stringify(pb.authStore.model));*/

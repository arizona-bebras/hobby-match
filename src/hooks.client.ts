import { pb } from '$lib/index';
import { browser } from '$app/environment';
import { goto } from '$app/navigation';
if (browser) {
  if (!window.Telegram.WebApp.isVersionAtLeast('7.0')) {
    window.location.replace('https://t.me/detoshumibot');
  }
  window.Telegram.WebApp.disableVerticalSwipes();
  pb.send('/api/collections/users/auth-with-telegram', {
    method: 'POST',
    body: {
      data: window.Telegram.WebApp.initData,
    },
  }).then((res) => {
    pb.authStore.save(res.token, res.record);
    if (
      !res.record.miniapp_name ||
      !res.record.gender ||
      !res.record.birth_date ||
      !res.record.location ||
      !res.record.user_info ||
      !res.record.user_photo ||
      res.record.interests.length < 3
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

import { pb } from '$lib/index';
import { browser } from '$app/environment';
if (browser) {
  pb.send('/api/collections/users/auth-with-telegram', {
    method: 'POST',
    body: {
      data: window.Telegram.WebApp.initData,
    },
  }).then((res) => {
    console.log(res.record);
    pb.authStore.save(res.token, res.record);
  });
  /*document.cookie = `telegram_id = ${pb.authStore.model?.telegram_id}`;*/
  localStorage.setItem('tmp_userdata', JSON.stringify(pb.authStore.model));
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

import { pb } from '$lib/index';
import { browser } from "$app/environment";
if (browser) {
    pb.send('/api/collections/users/auth-with-telegram', {
            method: 'POST',
            body: {
                data: window.Telegram.WebApp.initData
            }
            }).then(res => {
                console.log(res.record)
                pb.authStore.save(res.token, res.record);
        });
    }
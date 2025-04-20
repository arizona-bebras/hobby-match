<script lang="ts">
  import { pb } from '$lib/index'
  import { goto } from '$app/navigation'
  import { UserRoundPen, Search, Settings  } from '@lucide/svelte';
  import { init, themeParams , secondaryButton  } from '@telegram-apps/sdk-svelte';
  import { browser } from '$app/environment';
  import "./app.css";
  if (browser){
        init();
  }
  let { children } = $props();
  async function getUser() {
    //@ts-ignore 
    let user = await pb.collection('users').getFirstListItem(`telegram_id = "${pb.authStore.record.telegram_id}"`);
    console.log(user)
    if (user.location == "" || user.birth_date == "") {
      goto('./registration')
    }
  }
  getUser();

</script>

<div class="w-screen h-screen flex flex-col">
{@render children()}

<footer class="flex justify-between">
  <button type="button" class="btn preset-tonal-primary p-3 m-2"><a href="/Profile">Profile</a></button>
  <button type="button" class="btn preset-tonal-primary p-3 m-2"><a href="/Search">Search</a></button>
  <button type="button" class="btn preset-tonal-primary p-3 m-2"><a href="/Settings">Settings</a></button>
</footer>
</div>

<script lang="ts">
  import { pb } from '$lib/index';
  import { goto } from '$app/navigation';
  import { UserRoundPen, Search, Settings } from '@lucide/svelte';
  import {
    init,
    themeParams,
    secondaryButton,
  } from '@telegram-apps/sdk-svelte';
  import { browser } from '$app/environment';
  import './app.css';
  if (browser) {
    init();
  }
  let { children } = $props();
  async function getUser() {
    //@ts-ignore
    let user = await pb
      .collection('users')
      .getFirstListItem(`telegram_id = "${pb.authStore.record.telegram_id}"`);
    console.log(user);
    if (user.location == '' || user.birth_date == '') {
      goto('./registration');
    }
  }
  getUser();
</script>

<div class="w-screen max-w-full min-h-screen flex flex-col">
  <div class="flex flex-col bg-white flex-1 items-center text-black">
    {@render children()}
  </div>

  <footer class="flex justify-between w-full">
    <button type="button" class="btn preset-outlined-surface-500 p-3 m-2"
      ><a href="/Profile"><UserRoundPen /></a></button
    >
    <button type="button" class="btn preset-outlined-surface-500 p-3 m-2"
      ><a href="/Search"><Search /></a></button
    >
    <button type="button" class="btn preset-outlined-surface-500 p-3 m-2"
      ><a href="/Settings"><Settings /></a></button
    >
  </footer>
</div>

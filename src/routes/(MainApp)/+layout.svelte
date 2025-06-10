<script lang="ts">
  import { pb } from '$lib/index';
  import { goto } from '$app/navigation';
  import { UserRoundPen, Search, Settings } from '@lucide/svelte';
  import { browser } from '$app/environment';
  // import { initData } from '@telegram-apps/sdk-svelte';
  // if (browser) {
  //   init();
  //   console.log(initData.canSendAfter());
  // }
  let { children } = $props();
  // async function getUser() {
  //   //@ts-ignore
  //   let user = await pb
  //     .collection('users')
  //     .getFirstListItem(`telegram_id = "${pb.authStore.model?.telegram_id}"`);
  //   if (user.location == '' || user.birth_date == '') {
  //     goto('./registration');
  //   }
  // }
  //getUser();
</script>

<div class="w-screen max-w-full min-h-screen flex flex-col">
  <div class="flex flex-col bg-background flex-1 items-center text-text-color">
    {@render children()}
    <footer
      class="flex justify-between w-full fixed bottom-0 bg-background z-2"
    >
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
</div>

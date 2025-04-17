<script>
    let { children } = $props();
    import "./app.css";
    import { UserRoundPen, Search, Settings  } from '@lucide/svelte';
    import { init, themeParams , secondaryButton  } from '@telegram-apps/sdk-svelte';
    import { browser } from '$app/environment';
    if (browser){
        init();
    }
</script>

<div class="w-screen max-w-full min-h-screen flex flex-col">
  <div class="flex flex-col bg-white flex-1 items-center text-black">
    {@render children()}
  </div>

  <footer class="flex justify-between w-full">
    <button type="button" class="btn preset-outlined-surface-500 p-3 m-2"><a href="/Profile"><UserRoundPen/></a></button>
    <button type="button" class="btn preset-outlined-surface-500 p-3 m-2"><a href="/Search"><Search/></a></button>
    <button type="button" class="btn preset-outlined-surface-500 p-3 m-2"><a href="/Settings"><Settings/></a></button>
  </footer>
</div>

<script>
  import { Switch } from '@skeletonlabs/skeleton-svelte';
  let {data} = $props();
  let state = $state(false);
</script>

<main class="flex flex-col bg-red-400 flex-1 items-center">
  <p>{data.name}</p>
  <p>Ещё по одной</p>
  <Switch name="example" checked={state} onCheckedChange={(e) => (state = e.checked)} />
  {#if (state === true)}
    <img src="sticker.gif" alt="sticker" />
  {/if}
</main>



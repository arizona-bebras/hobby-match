<script>
  import { Star } from '@lucide/svelte';
  import GameWidgets from "./gameWidgets.svelte";
  import VideoWidgets from "./videoWidgets.svelte";
  import TextWidgets from "./textWidgets.svelte";
  // let open = $state(false)
</script>

<main class="flex flex-col bg-white flex-1 items-start text-black">
  <img src="Yepp.jpg" alt="person" class="w-[100%] h-95 rounded-b-[24px]">
  <div class="container font-[Inter] p-[16px]">
    <p class="font-extrabold text-[32px] flex items-center">Илона Абудаби, 18<Star /></p>
    <p class="font-semibold text-[20px]"><span>Махачкала, Россия</span></p>
<!--    <button onclick={() => open = true} class="w-[100%] h-15 bg-gray-200">+</button>-->
<!--    {#if (open)}-->
<!--      <Widgets bind:open={open} />-->
<!--    {/if}-->

    <TextWidgets/>
    <VideoWidgets/>
    <GameWidgets/>
    <GameWidgets/>
    <VideoWidgets/>
  </div>
</main>

<style>
  .container *{
      margin-bottom: 8px;

  }
</style>
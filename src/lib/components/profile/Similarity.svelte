<script lang="ts">
  import Emoji from '$lib/components/ui/emoji/emogi.svelte';
  import { Info } from '@lucide/svelte';
  let { userSimilarity }: { userSimilarity: number } = $props();
</script>

{#if userSimilarity < 75}
  <button
    class="flex items-center justify-between border-l-2 w-full p-2 font-medium"
  >
    <div class="flex items-center">
      {#if userSimilarity <= 25}
        <Emoji symbol="🔥" size={5} class="opacity-35" />
        <p>Вы похожи на 15%</p>
      {:else if userSimilarity >= 15 && userSimilarity < 50}
        <Emoji symbol="🔥" size={5} class="opacity-65" />
        <p>Вы похожи на 33%</p>
      {:else if userSimilarity >= 50 && userSimilarity < 75}
        <Emoji symbol="🔥" size={5} class="opacity-85" />
        <p>Вы похожи на 50%</p>
      {/if}
    </div>
    <Info />
  </button>
{:else}
  {@const condition = userSimilarity <= 85}
  <button
    class="flex items-center justify-between border-l-3 w-full {condition
      ? 'border-l-[#FF9900] text-[#FF9900] bg-[#FF9900]/15 font-semibold '
      : 'border-l-[#FF4646] text-[#FF4646] bg-[#FF4646]/25 font-bold '} p-2"
  >
    <div class="flex items-center">
      <Emoji
        symbol="🔥"
        size={condition ? 7 : 8}
        class="drop-shadow-[0_0_7px] {condition
          ? 'drop-shadow-[#FF9900]'
          : 'drop-shadow-[#FF4646]'}"
      />
      <p>{condition ? 'Вы похожи на 75%' : 'ПОЛНОЕ СОВПАДЕНИЕ!'}</p>
    </div>
    <Info />
  </button>
{/if}

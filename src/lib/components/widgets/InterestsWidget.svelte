<script lang="ts">
  import { ChevronUp } from '@lucide/svelte';
  import { cn } from '$lib/utils';
  let { interests }: { interests: { id: string; tag: string }[] } = $props();
  let isOpen = $state(false);
</script>

<div class="w-full p-[10px] transition-all relative" class:h-fit={isOpen}>
  <div
    class={cn(
      'flex flex-row w-fit gap-[8px] flex-wrap transition-all overflow-hidden',
      !isOpen && 'max-h-25',
    )}
    class:overflow-visible={isOpen}
    class:gradient={!isOpen}
  >
    {#each interests as interest}
      <div
        class="flex gap-[8px] h-fit bg-accent/25 rounded-[28px] px-[12px] py-[8px]"
      >
        <p class="text-accent">{interest.tag}</p>
      </div>
    {/each}
    {#if isOpen}
      <button
        onclick={() => (isOpen = false)}
        class="flex gap-[8px] bg-accent/25 rounded-[28px] h-fit px-[12px] py-[8px] text-accent"
      >
        <ChevronUp />
        Свернуть
      </button>
    {/if}
  </div>
  <button
    class="justify-center justify-self-center text-accent absolute bottom-4 left-0 right-0"
    onclick={() => (isOpen = true)}
    class:hidden={isOpen}>Развернуть</button
  >
</div>

<style>
  .gradient {
    mask-image: linear-gradient(to top, transparent 1%, white 100%);
  }
</style>

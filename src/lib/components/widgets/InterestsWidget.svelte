<script lang="ts">
  import { ChevronUp } from '@lucide/svelte';
  import { cn } from '$lib/utils';
  import type { Word } from '$lib/interests';

  let { interests }: { interests: Word[] } = $props();
  let isOpen = $state(false);

  function easeInOutExpo(x: number): number {
    return x === 0
      ? 0
      : x === 1
        ? 1
        : x < 0.5
          ? Math.pow(2, 20 * x - 10) / 2
          : (2 - Math.pow(2, -20 * x + 10)) / 2;
  }
</script>

<div class="w-full py-2 transition-all relative" class:h-fit={isOpen}>
  <div
    class={cn(
      'flex flex-row w-fit gap-[8px] flex-wrap transition-all overflow-hidden',
      !isOpen && 'max-h-25',
    )}
    class:overflow-visible={isOpen}
    class:gradient={!isOpen}
  >
    {#each interests as interest}
      {@const similarity =
        100 - Math.round(easeInOutExpo(interest.similarity) * 100)}
      {@const bgOpacity = similarity > 15 ? similarity : 15}
      <div
        class="flex gap-[8px] h-fit rounded-[28px] px-[12px] py-[8px]"
        style="background-color: color-mix(in srgb, var(--color-accent), transparent {bgOpacity}%)"
        style:box-shadow={similarity <= 50
          ? '0px 3px 10px color-mix(in srgb, var(--color-accent), transparent 40%)'
          : ''}
      >
        <p class="text-text-color">{interest.tag}</p>
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
    class="justify-center justify-self-center text-accent absolute bottom-0 left-0 right-0"
    onclick={() => (isOpen = true)}
    class:hidden={isOpen}>Развернуть</button
  >
</div>

<style>
  .gradient {
    mask-image: linear-gradient(to top, transparent 1%, white 100%);
  }
</style>

<script lang="ts">
  import { ChevronUp } from '@lucide/svelte';
  import { fade } from 'svelte/transition';
  let { interests }: { interests: { id: string; tag: string }[] } = $props();
  let isOpen = $state(false);
</script>

<div class="w-full h-30 p-[10px] transition-all" class:h-fit={isOpen}>
  <div
    class="flex flex-row w-fit gap-[8px] flex-wrap h-30 transition-all overflow-hidden"
    class:overflow-visible={isOpen}
    class:h-fit={isOpen}
    class:test={!isOpen}
  >
    {#each interests as interest}
      <div
        class="flex gap-[8px] h-fit bg-accent/25 rounded-[28px] px-[12px] py-[8px]"
      >
        <!--<img
          class="size-4 my-auto"
          src="https://cdnjs.cloudflare.com/ajax/libs/emoji-datasource-apple/15.1.2/img/apple/64/{[
            ...interest[0],
          ]
            .map((cp) => cp.codePointAt(0).toString(16))
            .join('-')}.png" 
        />-->
        <p class="text-accent">{interest.tag}</p>
      </div>
    {/each}
    {#if isOpen}
      <button
        class="flex gap-[8px] bg-accent/25 rounded-[28px] h-fit px-[12px] py-[8px]"
      >
        <p class="text-accent"><ChevronUp /></p>
        <p class="text-accent" onclick={() => (isOpen = false)}>Свернуть</p>
      </button>
    {/if}
  </div>
  <button
    class="flex justify-self-center text-accent relative -top-11"
    onclick={() => (isOpen = true)}
    class:hidden={isOpen}>Развернуть</button
  >
</div>

<style>
  .test {
    mask-image: linear-gradient(to top, transparent 1%, white 100%);
  }
</style>

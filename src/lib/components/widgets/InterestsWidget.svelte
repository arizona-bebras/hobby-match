<script lang="ts">
  import { ChevronUp } from '@lucide/svelte';
  import { fade } from 'svelte/transition';
  let { interests } = $props();
  let IsOpen = $state(false);
</script>

<div class="w-full h-30 p-[10px] transition-all" class:h-fit={IsOpen}>
  <div
    class="flex flex-row w-fit gap-[8px] flex-wrap h-30 transition-all overflow-hidden"
    class:overflow-visible={IsOpen}
    class:h-fit={IsOpen}
    class:test={!IsOpen}
  >
    {#each interests as interest}
      <div
        class="flex gap-[8px] bg-accent/25 rounded-[28px] px-[12px] py-[8px]"
      >
        <!--<img
          class="size-4 my-auto"
          src="https://cdnjs.cloudflare.com/ajax/libs/emoji-datasource-apple/15.1.2/img/apple/64/{[
            ...interest[0],
          ]
            .map((cp) => cp.codePointAt(0).toString(16))
            .join('-')}.png"
        />-->
        <p class="text-accent">{interest}</p>
      </div>
    {/each}
    <button
      class="flex gap-[8px] bg-accent/25 rounded-[28px] px-[12px] py-[8px]"
    >
      <p class="text-accent"><ChevronUp /></p>
      <p class="text-accent" onclick={() => (IsOpen = false)}>Свернуть</p>
    </button>
  </div>
  <button
    class="flex justify-self-center text-accent relative -top-11"
    onclick={() => (IsOpen = true)}
    class:hidden={IsOpen}>Развернуть</button
  >
</div>

<style>
  .test {
    mask-image: linear-gradient(to top, transparent 1%, white 100%);
  }
</style>

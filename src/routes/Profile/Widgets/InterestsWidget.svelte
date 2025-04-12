<script lang="ts">
  import { fade } from "svelte/transition"
  import { ChevronUp } from '@lucide/svelte'
  let { interests } = $props();
  let IsOpen = $state(false);

</script>

<div class="w-full h-20 p-[10px]" class:h-fit={IsOpen}>
  <div class="flex flex-row w-fit gap-[8px] flex-wrap h-16.5 overflow-hidden" class:overflow-visible={IsOpen} class:h-fit={IsOpen} class:test={!IsOpen}>
    {#each interests as interest}
    <div class="flex gap-[8px] bg-[#40A7E3]/25 rounded-[28px] px-[12px] py-[8px]">
      <img class="size-4 my-auto"src="https://cdnjs.cloudflare.com/ajax/libs/emoji-datasource-apple/15.1.2/img/apple/64/{([...interest[0]].map((cp) => cp.codePointAt(0).toString(16))).join('-')}.png" />
      <p class="text-[#40A7E3]">{interest[1]}</p>
    </div>
    {/each}
    <div class="flex gap-[8px] bg-[#40A7E3]/25 rounded-[28px] px-[12px] py-[8px]">
      <p><ChevronUp/></p>
      <button class="text-[#40A7E3]" onclick={() => IsOpen = false}>Свернуть</button>
    </div>
  </div>
  <button class="flex justify-self-center text-[#40A7E3] relative -top-7" onclick={() => IsOpen = true} class:hidden={IsOpen}>Развернуть</button>
</div>

<style>
  .test{
      mask-image: linear-gradient(to top, transparent 1%, white 40%);
  }
</style>

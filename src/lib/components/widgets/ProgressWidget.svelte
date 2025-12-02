<script lang="ts">
  import type { ProgressBar } from '$lib/widgetTypes/widgetTypes';

  let { data }: {data: ProgressBar} = $props();
</script>

<div class="GameBox flex-col">
  <p class="font-[Inter] text-[20px] font-medium">{data.description}</p>
  <div class="w-full bg-background rounded-lg overflow-hidden relative">
    <div
      class="h-8.25 bg-accent"
      style="
         width: calc(100% * {data.currentProgress} / {data.maxProgress});
         clip-path: polygon(0 0, 100% 0, 90% 100%, 0 100%);"
    ></div>
    <div
      class="absolute inset-0 flex items-center justify-center font-bold text-text-color"
    >
      {data.currentProgress} / {data.maxProgress}
    </div>
  </div>
</div>

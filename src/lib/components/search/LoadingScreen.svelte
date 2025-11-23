<script>
  import { ProgressRing } from '@skeletonlabs/skeleton-svelte';
</script>

<div
  class="flex flex-col items-center justify-center w-screen h-screen font-[Inter] text-lg font-medium"
>
  <ProgressRing
    value={null}
    meterStroke="stroke-accent"
    trackStroke="stroke-accent/25"
  ></ProgressRing>
  <p class="pt-4 text-text-color">Ищем подходящие профили...</p>
</div>

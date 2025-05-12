<script lang="ts">
  import { cn } from '$lib/utils';
  import { ArrowDown, ArrowUp, Pencil } from '@lucide/svelte';
  let { class: className = '' } = $props();
</script>

<div
  class={cn(
    'w-22.5 h-8.5 bg-accent rounded-full border-2 border-solid border-white flex justify-center items-center gap-2.5 absolute right-0 -top-2.5 z-1',
    className,
  )}
>
  <button><Pencil class="size-4 text-white" /></button>
  <button><ArrowUp class="size-4.5 text-white" /></button>
  <button><ArrowDown class="size-4.5 text-white" /></button>
</div>

<script lang="ts">
  import { cn } from '$lib/utils.js';
  import { LoaderCircle } from '@lucide/svelte';
  let {
    onClick,
    isButtonActive,
    class: className = '',
    isLoading = false,
    text = 'Сохранить',
  }: {
    onClick: () => void;
    isButtonActive: boolean;
    class?: string;
    isLoading?: boolean;
    text?: string;
  } = $props();
</script>

<button
  type="button"
  onclick={onClick}
  disabled={!isButtonActive || isLoading}
  class:bg-accent={isButtonActive}
  class:bg-inactive={!isButtonActive}
  class={cn('w-full h-12 rounded-xl mt-9 relative', className)}
>
  {text}
  {#if isLoading}
    <LoaderCircle class="ml-1 animate-spin absolute top-0 right-2 h-full" />
  {/if}
</button>

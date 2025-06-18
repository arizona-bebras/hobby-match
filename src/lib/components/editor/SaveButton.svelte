<script lang="ts">
  import { cn } from '$lib/utils.js';
  let {
    onClick,
    isButtonActive,
    class: className = '',
  }: {
    onClick: () => void;
    isButtonActive: boolean;
    class?: string;
  } = $props();
</script>

<button
  type="button"
  onclick={onClick}
  disabled={!isButtonActive}
  class:bg-accent={isButtonActive}
  class:bg-inactive={!isButtonActive}
  class={cn('w-full h-12 rounded-xl mt-9', className)}>Сохранить</button
>

<script lang="ts">
  import { changeWidgetPosition } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { cn } from '$lib/utils';
  import { ArrowDown, ArrowUp, Pencil } from '@lucide/svelte';
  import { invalidate } from '$app/navigation';
  import type { Widget } from '$lib/widgetTypes/widgetTypes';
  let {
    class: className = '',
    widget,
    onMove,
    onEdit,
  }: {
    class?: string;
    widget: Widget;
    onMove: (delta: number) => void;
    onEdit: CallableFunction;
  } = $props();
</script>

<div
  class={cn(
    'w-22.5 h-8.5 bg-accent rounded-full border-2 border-solid border-white flex justify-center items-center gap-2.5 absolute right-0 -top-2.5 z-1',
    className,
  )}
>
  <button onclick={() => onEdit()}><Pencil class="size-4 text-white" /></button>
  <button
    onclick={async () => {
      onMove(-1);
      await changeWidgetPosition(widget, -1);
      await invalidate('user:widgets');
    }}><ArrowUp class="size-4.5 text-white" /></button
  >
  <button
    onclick={async () => {
      onMove(1);
      await changeWidgetPosition(widget, 1);
      await invalidate('user:widgets');
    }}><ArrowDown class="size-4.5 text-white" /></button
  >
</div>

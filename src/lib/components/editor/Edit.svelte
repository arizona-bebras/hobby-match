<script lang="ts">
  import {
    createWidget,
    deleteWidget,
    updateWidget,
    changeWidgetPosition,
    updateWidgetsOrder,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import EditTgPost from '$lib/components/editor/EditTgPost.svelte';
  import EditSurvey from '$lib/components/editor/EditSurvey.svelte';
  import EditText from '$lib/components/editor/EditText.svelte';
  import EditVideo from '$lib/components/editor/EditVideo.svelte';
  import { cn } from '$lib/utils';
  import { ArrowDown, ArrowUp, Pencil } from '@lucide/svelte';
  import type { WidgetWithService } from '$lib/components/widgetConstructors/widgetsConstructor';
  import EditImage from '$lib/components/editor/EditImage.svelte';
  import { invalidate } from '$app/navigation';
  import EditSocial from '$lib/components/editor/EditSocial.svelte';
  import EditToDo from '$lib/components/editor/EditToDo.svelte';
  import EditProgress from '$lib/components/editor/EditProgress.svelte';
  let {
    class: className = '',
    widgetType,
    widgetId,
    widgets,
    onMove,
  }: {
    class: string;
    widgetType: string;
    widgetId: string;
    widgets: WidgetWithService[];
    onMove: (delta: number) => void;
  } = $props();
  let nextStage = $state(false);
  // Если тип виджета одинаковый, он берёт первый. Нужно сравнивать ещё id виджета
  let selectedWidget: WidgetWithService = $derived(
    widgets.find(
      (widget) =>
        widget.widget.data.type === widgetType && widget.widget.id === widgetId,
    ),
  );
  $effect(() => {
    if (nextStage) {
      window.Telegram.WebApp.MainButton.hide();
    } else {
      window.Telegram.WebApp.MainButton.show();
    }
  });
</script>

<div
  class={cn(
    'w-22.5 h-8.5 bg-accent rounded-full border-2 border-solid border-white flex justify-center items-center gap-2.5 absolute right-0 -top-2.5 z-1',
    className,
  )}
>
  <button
    onclick={() => {
      nextStage = true;
    }}><Pencil class="size-4 text-white" /></button
  >
  <button
    onclick={async () => {
      onMove(-1);
      await changeWidgetPosition(selectedWidget.widget, -1);
      await invalidate('user:widgets');
      console.log($state.snapshot(widgets));
    }}><ArrowUp class="size-4.5 text-white" /></button
  >
  <button
    onclick={async () => {
      onMove(1);
      await changeWidgetPosition(selectedWidget.widget, 1);
      await invalidate('user:widgets');
      console.log($state.snapshot(widgets));
    }}><ArrowDown class="size-4.5 text-white" /></button
  >

  {#if nextStage}
    {#if widgetType === 'video'}
      <EditVideo bind:nextStage numberOfWidgets={widgets.length} {widgetId} />
    {:else if widgetType === 'text'}
      <EditText bind:nextStage numberOfWidgets={widgets.length} {widgetId} />
    {:else if widgetType === 'survey'}
      <EditSurvey bind:nextStage numberOfWidgets={widgets.length} {widgetId} />
    {:else if widgetType === 'social_media'}
      <EditSocial bind:nextStage numberOfWidgets={widgets.length} {widgetId} />
    {:else if widgetType === 'post'}
      <EditTgPost bind:nextStage numberOfWidgets={widgets.length} {widgetId} />
    {:else if widgetType === 'photo'}
      <EditImage bind:nextStage numberOfWidgets={widgets.length} {widgetId} />
    {:else if widgetType === 'todo'}
      <EditToDo bind:nextStage numberOfWidgets={widgets.length} {widgetId} />
    {:else if widgetType === 'progress_bar'}
      <EditProgress
        bind:nextStage
        numberOfWidgets={widgets.length}
        {widgetId}
      />
    {/if}
  {/if}
</div>

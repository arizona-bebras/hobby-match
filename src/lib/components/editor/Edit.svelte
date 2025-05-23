<script lang="ts">
  import {
    createWidget,
    deleteWidget,
    updateWidget,
    changeWidgetPosition,
    updateWidgetsOrder,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import EditTgPost from '$lib/components/editor/EditTgPost.svelte';
  import EditTgAccount from '$lib/components/editor/EditTgAccount.svelte';
  import EditSurvey from '$lib/components/editor/EditSurvey.svelte';
  import EditText from '$lib/components/editor/EditText.svelte';
  import EditVideo from '$lib/components/editor/EditVideo.svelte';
  import { cn } from '$lib/utils';
  import { ArrowDown, ArrowUp, Pencil } from '@lucide/svelte';
  import type { WidgetWithService } from '$lib/components/widgetConstructors/widgetsConstructor';
  let {
    class: className = '',
    widgetType,
    widgetId,
    widgets,
    form,
  }: {
    class: string;
    widgetType: string;
    widgetId: string;
    widgets: WidgetWithService[];
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
    onclick={() => {
      changeWidgetPosition(widgets, selectedWidget, -1);
      console.log($state.snapshot(widgets));
    }}><ArrowUp class="size-4.5 text-white" /></button
  >
  <button
    onclick={() => {
      changeWidgetPosition(widgets, selectedWidget, 1);
      console.log($state.snapshot(widgets));
    }}><ArrowDown class="size-4.5 text-white" /></button
  >

  {#if nextStage}
    {#if widgetType === 'video'}
      <EditVideo bind:nextStage />
    {:else if widgetType === 'text'}
      <EditText bind:nextStage {form} />
    {:else if widgetType === 'survey'}
      <EditSurvey bind:nextStage />
    {:else if widgetType === 'Телеграм Аккаунт'}
      <EditTgAccount bind:nextStage />
    {:else if widgetType === 'Телеграм Пост'}
      <EditTgPost bind:nextStage />
    {/if}
  {/if}
</div>

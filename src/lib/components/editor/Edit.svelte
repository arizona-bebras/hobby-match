<script lang="ts">
  import {
    createWidget,
    deleteWidget,
    updateWidget,
    changeWidgetPostion,
    updateWidgetsOrder,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import EditTgPost from '$lib/components/editor/EditTgPost.svelte';
  import EditTgAccount from '$lib/components/editor/EditTgAccount.svelte';
  import EditSurvey from '$lib/components/editor/EditSurvey.svelte';
  import EditText from '$lib/components/editor/EditText.svelte';
  import EditVideo from '$lib/components/editor/EditVideo.svelte';
  import { cn } from '$lib/utils';
  import { ArrowDown, ArrowUp, Pencil } from '@lucide/svelte';
  let { class: className = '', widgetType } = $props();
  let nextStage = $state(false);

  let someTest = {
    widget: {
      telegram_id: '123',
      order: 1,
      data: {
        type: 'text',
        text: '„Люблю играть в Valorant. Часто говорят что выгляжу как будто сгенерирована нейросетью“',
      },
    },
    deleteStatus: false,
  };

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
  <button><ArrowUp class="size-4.5 text-white" /></button>
  <button><ArrowDown class="size-4.5 text-white" /></button>

  {#if nextStage}
    {#if widgetType === 'video'}
      <EditVideo bind:nextStage />
    {:else if widgetType === 'text'}
      <EditText bind:nextStage />
    {:else if widgetType === 'survey'}
      <EditSurvey bind:nextStage />
    {:else if widgetType === 'Телеграм Аккаунт'}
      <EditTgAccount bind:nextStage />
    {:else if widgetType === 'Телеграм Пост'}
      <EditTgPost bind:nextStage />
    {/if}
  {/if}
</div>

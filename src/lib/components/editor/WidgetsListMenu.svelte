<script lang="ts">
  import * as Dialog from '$lib/components/ui/dialog/index.js';
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { tick } from 'svelte';
  let {
    showWidgetMenu: open = $bindable(),
    addedWidget = $bindable(),
    nextStage = $bindable(),
  } = $props();
  let widgetsNames = [
    'Видео',
    'Социальная сеть',
    'Опрос',
    'Текст',
    'Телеграм Пост',
    'Изображение',
    'Список задач',
  ];
  window.Telegram.WebApp.MainButton.hide();

  function onOpenChange() {
    window.Telegram.WebApp.MainButton.show();
  }
</script>

<Dialog.Root bind:open {onOpenChange}>
  <Dialog.Content>
    <Dialog.Header>
      {#each widgetsNames as name}
        <Dialog.Title
          ><button
            onclick={() => {
              open = false;
              nextStage = true;
              addedWidget = name;
            }}>{name}</button
          ></Dialog.Title
        >
      {/each}
    </Dialog.Header>
  </Dialog.Content>
  <Dialog.Close />
</Dialog.Root>

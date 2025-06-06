<script lang="ts">
  import * as Dialog from '$lib/components/ui/dialog/index.js';
  let { onClick }: { onClick: (name: string) => void } = $props();
  let widgetsNames = [
    'Аудио',
    'Видео',
    'Социальная сеть',
    'Опрос',
    'Текст',
    'Телеграм Пост',
    'Изображение',
    'Список задач',
    'Прогресс',
    'Время игры',
  ];
</script>

<Dialog.Root open={true}>
  <Dialog.Content>
    <Dialog.Header>
      {#each widgetsNames as name}
        <Dialog.Title
          ><button onclick={() => onClick(name)}>{name}</button></Dialog.Title
        >
      {/each}
    </Dialog.Header>
  </Dialog.Content>
  <Dialog.Close />
</Dialog.Root>

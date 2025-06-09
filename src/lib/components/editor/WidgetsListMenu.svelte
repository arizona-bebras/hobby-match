<script lang="ts">
  import * as Dialog from '$lib/components/ui/dialog/index.js';
  import type { WidgetType } from '$lib/widgetTypes/widgetTypes';
  let {
    onClick,
    open = $bindable(),
  }: { onClick: (name?: WidgetType) => void; open: boolean } = $props();
  let widgetNames: Partial<Record<WidgetType, string>> = {
    audio: 'Аудио',
    post: 'Телеграм пост',
    // geo: 'Маршрут',
    photo: 'Изображение',
    progress_bar: 'Прогресс',
    social_media: 'Социальная сеть',
    steam_game: 'Время в игре',
    // sticker: 'Стикер',
    survey: 'Опрос',
    text: 'Текст',
    todo: 'Список задач',
    video: 'Видео',
  } as const;
</script>

<!--if (!state) onClick()-->
<Dialog.Root bind:open onOpenChange={(state) => !state && onClick()}>
  <Dialog.Content>
    <Dialog.Header>
      {#each Object.entries(widgetNames) as [type, name]}
        <Dialog.Title>
          <button
            onclick={() => {
              // @ts-expect-error object entries sucks
              onClick(type);
            }}
          >
            {name}
          </button>
        </Dialog.Title>
      {/each}
    </Dialog.Header>
  </Dialog.Content>
  <Dialog.Close />
</Dialog.Root>

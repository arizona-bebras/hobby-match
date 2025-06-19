<script lang="ts">
  import * as Dialog from '$lib/components/ui/dialog/index.js';
  import type { WidgetType } from '$lib/widgetTypes/widgetTypes';
  import {
    Music,
    Image,
    Loader,
    MessageCircle,
    Send,
    Gamepad2,
    ClipboardPenLine,
    ChartNoAxesCombined,
    CaseSensitive,
    ListChecks,
    Video,
    Icon,
  } from '@lucide/svelte';
  let {
    onClick,
    open = $bindable(),
  }: { onClick: (name?: WidgetType) => void; open: boolean } = $props();
  const WIDGETS: Partial<
    Record<WidgetType, { label: string; icon: typeof Icon }>
  > = {
    audio: {
      label: 'Аудио',
      icon: Music,
    },
    post: { label: 'Телеграм пост', icon: Send },
    // geo: 'Маршрут',
    photo: { label: 'Изображение', icon: Image },
    progress_bar: { label: 'Прогресс', icon: ChartNoAxesCombined },
    social_media: { label: 'Социальная сеть', icon: MessageCircle },
    steam_game: { label: 'Время в игре', icon: Gamepad2 },
    // sticker: 'Стикер',
    survey: { label: 'Опрос', icon: ClipboardPenLine },
    text: { label: 'Текст', icon: CaseSensitive },
    todo: { label: 'Список задач', icon: ListChecks },
    video: { label: 'Видео', icon: Video },
  } as const;
</script>

<!--if (!state) onClick()-->
<Dialog.Root bind:open onOpenChange={(state) => !state && onClick()}>
  <Dialog.Content>
    <Dialog.Header
      ><Dialog.Title
        class="text-accent-foreground bg-accent/25 w-fit p-3 mx-auto rounded-xl"
        >Добавить виджет</Dialog.Title
      ></Dialog.Header
    >
    {#each Object.entries(WIDGETS) as [type, widget]}
      {@const Icon = widget.icon}

      <button
        onclick={() => {
          // @ts-expect-error object entries sucks
          onClick(type);
        }}
      >
        <div class="flex text-accent-foreground text-base font-medium">
          <Icon class="mr-2 mt-0.75 size-6" />
          <p class="">{widget.label}</p>
        </div>
      </button>
    {/each}
  </Dialog.Content>
</Dialog.Root>

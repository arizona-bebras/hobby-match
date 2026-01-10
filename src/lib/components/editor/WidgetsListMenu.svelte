<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { buttonVariants } from '$lib/components/ui/button/index.js';
  import { Input } from '$lib/components/ui/input/index.js';
  import { Label } from '$lib/components/ui/label/index.js';
  import * as Dialog from '$lib/components/ui/dialog/index.js';
  import type { WidgetType } from '$lib/widgetTypes/widgetTypes';
  import { type Widget } from '$lib/widgetTypes/widgetTypes';
  import DuckSoulsIcon from '$lib/assets/DuckSoulsIcon.jpg';
  import DuckHello from '$lib/assets/DuckHello.png';
  import DuckUmbrella from '$lib/assets/DuckUmbrella.png';
  import {
    Audio,
    ToDo,
    Social,
    Photo,
    Survey,
    TgPost,
    Progress,
    Text,
    Game,
    Video,
    Activity,
  } from '$lib/components/widgets/index';
  console.log(DuckHello, DuckUmbrella);
  let {
    onClick,
    open = $bindable(),
  }: { onClick: (name?: WidgetType) => void; open: boolean } = $props();
  let widgetFilter = $state('');
  const WIDGETS: {
    type: WidgetType;
    label: string;
    widget:
      | Audio
      | Game
      | Photo
      | Progress
      | Social
      | Survey
      | Text
      | TgPost
      | ToDo
      | Video
      | Activity;
    data: Widget['data'] | any;
    additionalData?: Widget['additionalData'];
  }[] = [
    {
      type: 'audio',
      label: 'Аудио',
      widget: Audio,
      data: {
        type: 'audio',
        platform: 'Yandex',
        link: 'https://music.yandex.ru/track/38887766?utm_source=web&utm_medium=copy_link',
      },
    },
    {
      type: 'steam_game',
      label: 'Время в игре',
      widget: Game,
      data: {
        type: 'steam_game',
        hours_played: 245,
        icon: DuckSoulsIcon,
        title: 'Duck Souls',
      },
    },
    {
      type: 'photo',
      label: 'Изображение',
      widget: Photo,
      files: [DuckHello, DuckUmbrella],
      data: {
        type: 'photo',
      },
    },
    {
      type: 'progress_bar',
      label: 'Прогресс',
      widget: Progress,
      data: {
        type: 'progress_bar',
        description: 'Проплыть 100 км',
        currentProgress: 76,
        maxProgress: 100,
      },
    },
    {
      type: 'social_media',
      label: 'Социальная сеть',
      widget: Social,
      data: {
        type: 'social_media',
        platform: 'Twitch',
        link: 'https://www.twitch.tv/twitch',
      },
      additionalData: {
        type: 'Twitch',
        title: 'Twitch',
        followers: 2412740,
      },
    },
    {
      type: 'survey',
      label: 'Опрос',
      widget: Survey,
      data: {
        type: 'survey',
        question: 'Какой хлеб лучше',
        options: [{ description: 'Белый' }, { description: 'Чёрный' }],
      },
    },
    {
      type: 'text',
      label: 'Текст',
      widget: Text,
      data: {
        type: 'text',
        text: 'Если тонешь — всплывёшь. Ты же утка!',
      },
    },
    {
      type: 'todo',
      label: 'Список задач',
      widget: ToDo,
      data: {
        type: 'todo',
        title: '',
        tasks: [
          { description: 'Поставить цель на эту неделю', isCompleted: true },
          { description: 'Спланировать вечер', isCompleted: false },
          {
            description: 'Поздороваться с одной новой уткой',
            isCompleted: true,
          },
        ],
      },
    },
    {
      type: 'video',
      label: 'Видео',
      widget: Video,
      data: {
        type: 'video',
        link: 'https://rutube.ru/video/c6cc4d620b1d4338901770a44b3e82f4/?r=wd',
        platform: 'Rutube',
      },
    },
    {
      type: 'post',
      label: 'Телеграм пост',
      widget: TgPost,
      data: {
        type: 'post',
        link: 'durov/337',
      },
    },
    {
      type: 'activity',
      label: 'Деятельность',
      widget: Activity,
      data: {
        status: {
          isStudent: true,
          isWorker: true,
        },
        education: {
          type: 'university',
          educationStage: 'bachelor',
          course: 2,
        },
        work: {
          company: 'Ducko',
          job_title: 'full-stack',
          experience: 2,
        },
      },
    },
  ];
  $inspect(widgetFilter);
</script>

<Sheet.Root
  bind:open
  onOpenChange={(state) => {
    if (!state) {
      onClick();
      window.Telegram.WebApp.MainButton.show();
    }
  }}
>
  <Sheet.Content side="bottom" class="max-h-[calc(100vh-75px)]">
    <Sheet.Header>
      <Sheet.Title class="text-center text-[16px] font-medium mb-6"
        >Новый виджет</Sheet.Title
      >
      <Sheet.Description class="text-text-color mb-6">
        <Input placeholder="Поиск по виджетам" bind:value={widgetFilter} />
      </Sheet.Description>
    </Sheet.Header>
    <div class="overflow-y-scroll h-[calc(100vh-195px)] pt-2">
      <!--      <button onclick={() => onClick('text')}>Test Text</button>-->
      {#each WIDGETS as widget (widget.type)}
        {#if widget.label.toLowerCase().includes(widgetFilter.toLowerCase())}
          {@const Component = widget.widget}
          <div
            onclick={(e) => {
              e.stopPropagation();
              onClick(widget.type);
            }}
            class="border-2 border-dashed border-white/25 p-2 rounded-xl relative mb-6 font-semibold
"
          >
            <span
              class="absolute -top-3.75 left-1/2 transform -translate-x-1/2 z-5"
            >
              {widget.label}
            </span>
            {#if widget.type === 'photo'}
              <div class="pointer-events-none">
                <Component urls={widget.files} isTestImage={true} />
              </div>
            {:else if widget.type === 'survey'}
              <div class="pointer-events-none">
                <Component
                  data={widget.data}
                  id={1}
                  survey={{
                    type: 'survey',
                    stats: [0, 1],
                    myVote: 1,
                  }}
                />
              </div>
            {:else}
              <div class="pointer-events-none">
                <Component
                  data={widget.data}
                  socialMediaData={widget.additionalData}
                />
              </div>
            {/if}
          </div>
        {/if}
      {/each}
    </div>
    <Sheet.Footer>
      <Sheet.Close class={buttonVariants({ variant: 'outline' })}
        >Save changes</Sheet.Close
      >
    </Sheet.Footer>
  </Sheet.Content>
</Sheet.Root>

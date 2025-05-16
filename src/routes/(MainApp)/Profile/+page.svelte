<script lang="ts">
  import {
    createWidget,
    deleteWidget,
    updateWidget,
    changeWidgetPostion,
    updateWidgetsOrder,
  } from '$lib/components/widgetConstructors/widgetsConstructor';

  import { Star, Plus } from '@lucide/svelte';
  import GameWidget from '$lib/components/widgets/GameWidget.svelte';
  import VideoWidget from '$lib/components/widgets/VideoWidget.svelte';
  import TextWidget from '$lib/components/widgets/TextWidget.svelte';
  import SocialWidget from '$lib/components/widgets/SocialWidget.svelte';
  import StickerWidget from '$lib/components/widgets/StickerWidget.svelte';
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import ProgressWidget from '$lib/components/widgets/ProgressWidget.svelte';
  import InterestsWidget from '$lib/components/widgets/InterestsWidget.svelte';
  import ToDoWidget from '$lib/components/widgets/ToDoWidget.svelte';
  import SurveyWidget from '$lib/components/widgets/SurveyWidget.svelte';
  import PhotoWidget from '$lib/components/widgets/PhotoWidget.svelte';
  import type { BasicWidget } from '$lib/widgetsTypes/widgetsTypes';

  import EditTgPost from '$lib/components/editor/EditTgPost.svelte';
  import EditTgAccount from '$lib/components/editor/EditTgAccount.svelte';
  import EditSurvey from '$lib/components/editor/EditSurvey.svelte';
  import EditText from '$lib/components/editor/EditText.svelte';
  import EditVideo from '$lib/components/editor/EditVideo.svelte';
  import type { Option } from '$lib/widgetTypes/widgetTypes';
  import Edit from '$lib/components/editor/Edit.svelte';
  import WidgetsListMenu from '$lib/components/editor/WidgetsListMenu.svelte';
  import { Input } from '$lib/components/ui/input';

  const widgets: BasicWidget[] = [
    {
      telegram_id: '123',
      order: 1,
      data: {
        type: 'text',
        text: '„Люблю играть в Valorant. Часто говорят что выгляжу как будто сгенерирована нейросетью“',
      },
    },
    {
      telegram_id: '123',
      order: 2,
      data: {
        game: 'Valorant',
        type: 'steam_game',
        steam_user_id: 'https://steamcommunity.com/profiles/76561198295231108/',
        hours_played: 1489,
        game_icon: 'Valorant.png',
      },
    },
    {
      telegram_id: '123',
      order: 3,
      type: 'video',
      data: {
        type: 'video',
        link: 'https://www.tiktok.com/player/v1/7492189973605551366', // Нужно чтобы приходили последнии цифры
        platform: 'TikTok',
      },
    },
    // {
    //   telegram_id: '123',
    //   order: 4,
    //   data: {
    //     type: 'video',
    //     link: 'https://www.youtube.com/embed/rs6Y4kZ8qtw?si=gtMMU8kaUHylKI32', // Нужно чтобы приходили последнии цифры
    //     platform: 'YouTube',
    //   },
    // },
    {
      telegram_id: '123',
      order: 5,
      data: {
        type: 'video',
        link: 'https://rutube.ru/play/embed/af22ea10424506d17d00fd15ab79804f/', // Нужно чтобы приходили последнии цифры
        platform: 'Rutube',
      },
    },
    {
      telegram_id: '123',
      order: 6,
      data: {
        type: 'social_media',
        platform: 'Youtube',
        username: '@lowderplay',
        link: 'https://www.youtube.com/@lowderplay/videos',
        subscribers: 11,
      },
    },
    {
      telegram_id: '123',
      order: 7,
      data: {
        type: 'social_media',
        platform: 'Steam',
        username: "I'M SCHIZOPHRENIC AND HAVE A GUN",
        link: 'https://steamcommunity.com/profiles/76561198295231108/',
        subscribers: 1,
      },
    },
    {
      telegram_id: '123',
      order: 8,
      data: {
        type: 'social_media',
        platform: 'Twitch',
        username: 'shadowkekw',
        link: 'https://www.twitch.tv/shadowkekw',
        subscribers: 42,
      },
    },
    {
      telegram_id: '123',
      order: 9,
      data: {
        type: 'social_media',
        platform: 'Twitter',
        username: 'elonmusk',
        link: 'https://x.com/elonmusk',
        subscribers: 1,
      },
    },
    {
      telegram_id: '123',
      order: 10,
      data: {
        type: 'social_media',
        platform: 'VK',
        username: 'Илья Спицын',
        link: 'https://vk.com/id378488092',
        subscribers: 34,
      },
    },
    {
      telegram_id: '123',
      order: 11,
      data: {
        type: 'progress_bar',
        description: 'Количество тренировок',
        currentProgress: 65,
        maxProgress: 100,
      },
    },
    {
      telegram_id: '123',
      order: 12,
      data: {
        type: 'photo',
        photos: [
          {
            name: 'https://images.pexels.com/photos/1525041/pexels-photo-1525041.jpeg?cs=srgb&dl=pexels-francesco-ungaro-1525041.jpg&fm=jpg',
            width: 150,
            height: 200,
            size: 25,
          },
          {
            name: 'https://learn.zoner.com/wp-content/uploads/2018/08/how-to-place-one-photo-inside-another-create-your-own-photo-collage-using-layers.jpg',
            width: 350,
            height: 200,
            size: 50,
          },
        ],
      },
    },
    {
      telegram_id: '123',
      order: 13,
      data: {
        type: 'todo',
        title: 'Список дней на день',
        tasks: [
          {
            order: 1,
            description: 'Проснуться',
            isCompleted: true,
          },
          {
            order: 1,
            description: 'Сходить на пары',
            isCompleted: false,
          },
          {
            order: 1,
            description: 'Позаниматься',
            isCompleted: true,
          },
          {
            order: 1,
            description: 'Покушать :)',
            isCompleted: true,
          },
        ],
      },
    },
    {
      telegram_id: '123',
      order: 14,
      data: {
        type: 'survey',
        question: 'Следующее видео',
        options: [
          {
            description: 'Моя Косметичка',
            votes: 125,
          },
          {
            description: 'Диф. уравнения',
            votes: 25,
          },
        ],
        summuryVotes: 150,
      },
    },
  ];
  const sortedWidgets = [...widgets].sort((a, b) => a.order - b.order);

  let changeMode = $state(false);
  $effect(() => {
    window.Telegram.WebApp.MainButton.setParams({
      text: changeMode ? 'Сохранить' : 'Изменить анкету',
      is_visible: true,
    });
  });
  window.Telegram.WebApp.MainButton.onClick(() => (changeMode = !changeMode));

  let showWidgetMenu = $state(false);
  let addedWidget = $state('');
  let nextStage = $state(false); //dakdoawdkpowadkwapodkwapodkwadaopd)))))

  let height = $state(380);
  $effect(() => {
    window.addEventListener(
      'scroll',
      () => (height = Math.max(300, 380 - window.scrollY * 0.5)),
    );
  });

  let someTest = [
    {
      widget: {
        telegram_id: '123',
        order: 1,
        data: {
          type: 'text',
          text: '„Люблю играть в Valorant. Часто говорят что выгляжу как будто сгенерирована нейросетью“',
        },
      },
      deleteStatus: false,
      changeStatus: false,
      additionalData: {
        socialMeidaData: 25,
      },
    },
    {
      widget: {
        telegram_id: '123',
        order: 11,
        data: {
          type: 'progress_bar',
          description: 'Количество тренировок',
          currentProgress: 65,
          maxProgress: 100,
        },
      },
      deleteStatus: false,
      changeStatus: false,
      additionalData: {
        socialMeidaData: 25,
      },
    },
  ];
</script>

{#if height - 90 <= 210}
  <div class="w-full h-16 bg-background fixed flex z-2">
    <img src="Girl.png" class="w-16 h-16 rounded-full p-2" alt="UserPhoto" />
    <div class="container font-[Inter] p-[4px]">
      <p class="font-extrabold text-[16px] flex items-center">
        Илона Абудаби, 18<Star />
      </p>
      <p class="font-semibold text-[16px]"><span>Махачкала, Россия</span></p>
    </div>
  </div>
{/if}
<img
  src="Girl.png"
  alt="person"
  class="w-full h-95 rounded-b-[24px] object-cover"
  style="height: {height}px"
/>
<div class="container font-[Inter] p-[16px] w-full max-w-full relative">
  <p class="font-extrabold text-[32px] flex items-center">
    Илона Абудаби, 18<Star />
  </p>
  <p class="font-semibold text-[20px]"><span>Екатеринбург, Россия</span></p>
  <InterestsWidget
    interests={[
      ['🎮', 'Игры'],
      ['🚗', 'Развлечения'],
      ['💅', 'Красота'],
      ['🧮', 'Математика'],
      ['🎮', 'Кодинг'],
      ['🎮', 'игры'],
      ['🎮', 'игры'],
      ['🎮', 'игры'],
    ]}
  />

  {#if changeMode}
    <button
      class="my-2 w-full h-12 bg-[#34C759] rounded-xl flex justify-center items-center gap-3"
      onclick={() => (showWidgetMenu = !showWidgetMenu)}
    >
      <Plus class="size-5" />
      <p class="font-medium">Добавить виджет</p>
    </button>
  {/if}

  {#if showWidgetMenu}
    <WidgetsListMenu bind:showWidgetMenu bind:addedWidget bind:nextStage />
  {/if}
  {#if addedWidget === 'Видео'}
    <EditVideo bind:nextStage bind:addedWidget />
  {:else if addedWidget === 'Текст'}
    <EditText bind:nextStage bind:addedWidget />
  {:else if addedWidget === 'Опрос'}
    <EditSurvey bind:nextStage bind:addedWidget />
  {:else if addedWidget === 'Телеграм Аккаунт'}
    <EditTgAccount bind:nextStage bind:addedWidget />
  {:else if addedWidget === 'Телеграм Пост'}
    <EditTgPost bind:nextStage bind:addedWidget />
  {/if}

  {#each sortedWidgets as widget}
    <div class="relative">
      {#if changeMode}
        <Edit widgetType={widget.data.type} />
      {/if}
      {#if widget.data.type === 'text'}
        <TextWidget data={widget.data} />
      {:else if widget.data.type === 'steam_game'}
        <GameWidget data={widget.data} />
      {:else if widget.data.type === 'video'}
        <VideoWidget data={widget.data} />
      {:else if widget.data.type === 'social_media'}
        <SocialWidget data={widget.data} socialMediaData = {widget.additionalData.socialMediaData} />
      {:else if widget.data.type === 'progress_bar'}
        <ProgressWidget data={widget.data} />
        <!--{:else if widget.data.type === 'photo'}-->
        <!--  <ProgressWidget data={widget.data} />-->
      {:else if widget.data.type === 'todo'}
        <ToDoWidget data={widget.data} />
      {:else if widget.data.type === 'survey'}
        <SurveyWidget data={widget.data} />
      {:else if widget.data.type === 'photo'}
        <PhotoWidget data={widget.data} />
      {/if}
    </div>
  {/each}
  <!--    <button onclick={() => (open = true)}>fsdfsd</button>-->
  <!--    <Video bind:open bind:addedWidget />-->
  <!--    <StickerWidget image="Stickers/Fire.svg" left={285} top={185} rotate={0} />-->
  <!--    <StickerWidget image="Stickers/Eyes.svg" left={15} top={315} rotate={-30} />-->
  <!--    &lt;!&ndash; Сделать чтобы при True комнонент помечал как выполненное задание &ndash;&gt;-->
  <!--    <ToDoWidget-->
  <!--    title="Что хочу сделать:)"-->
  <!--    tasks={[-->
  <!--      ['hello', true],-->
  <!--      ['world', true],-->
  <!--      [25, true],-->
  <!--      [50, true],-->
  <!--      [75, true],-->
  <!--      [100, false],-->
  <!--    ]}-->
  <!--  />-->
  <!--  <SurveyWidget-->
  <!--    tasks={[-->
  <!--      ['hello', 100],-->
  <!--      ['world', 25],-->
  <!--    ]}-->
  <!--    votes={125}-->
  <!--  />-->
</div>

<style>
  .container {
    margin-bottom: 8px;
  }
</style>

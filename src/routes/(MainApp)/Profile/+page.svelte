<script lang="ts">
  import { Star } from '@lucide/svelte';
  import GameWidget from '$lib/components/Widgets/GameWidget.svelte';
  import VideoWidget from '$lib/components/Widgets/VideoWidget.svelte';
  import TextWidget from '$lib/components/Widgets/TextWidget.svelte';
  import SocialWidget from '$lib/components/Widgets/SocialWidget.svelte';
  import StickerWidget from '$lib/components/Widgets/StickerWidget.svelte';
  import ProgressWidget from '$lib/components/Widgets/ProgressWidget.svelte';
  import InterestsWidget from '$lib/components/Widgets/InterestsWidget.svelte';
  import ToDoWidget from '$lib/components/Widgets/ToDoWidget.svelte';
  import SurveyWidget from '$lib/components/Widgets/SurveyWidget.svelte';
  import PhotoWidget from '$lib/components/Widgets/PhotoWidget.svelte';
  import type { BasicWidget } from '$lib/widgetsTypes/widgetsTypes';
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
    {
      telegram_id: '123',
      order: 4,
      data: {
        type: 'video',
        link: 'https://www.youtube.com/embed/rs6Y4kZ8qtw?si=gtMMU8kaUHylKI32', // Нужно чтобы приходили последнии цифры
        platform: 'YouTube',
      },
    },
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
        description: 'Анжумания',
        currentProgress: 62,
        maxProgress: 150,
      },
    },
  ];
  const sortedWidgets = [...widgets].sort((a, b) => a.order - b.order);

  let height = $state(380);
  $effect(() => {
    window.addEventListener(
      'scroll',
      () => (height = Math.max(300, 380 - window.scrollY * 0.5)),
    );
  });
</script>

{#if height - 90 <= 210}
  <div class="w-full h-16 bg-background fixed flex z-2">
    <img src="Girl.png" class="w-16 h-16 rounded-full p-2" />
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
      ['🎮', 'игры'],
      ['🚗', 'развлечение'],
      ['✔️', 'АХАХАХАХАХАХ'],
      ['🎮', 'игры'],
      ['🎮', 'игры'],
      ['🎮', 'игры'],
      ['🎮', 'игры'],
      ['🎮', 'игры'],
    ]}
  />
  {#each sortedWidgets as widget}
    {#if widget.data.type === 'text'}
      <TextWidget data={widget.data} />
    {:else if widget.data.type === 'steam_game'}
      <GameWidget data={widget.data} />
    {:else if widget.data.type === 'video'}
      <VideoWidget data={widget.data} />
    {:else if widget.data.type === 'social_media'}
      <SocialWidget data={widget.data} />
    {:else if widget.data.type === 'progress_bar'}
      <ProgressWidget data={widget.data} />
    {/if}
  {/each}
  <!--  <StickerWidget image="Stickers/Fire.svg" left={285} top={185} rotate={0} />-->
  <!--  <StickerWidget image="Stickers/Eyes.svg" left={15} top={315} rotate={-30} />-->
  <!--  &lt;!&ndash; Сделать чтобы при True комнонент помечал как выполненное задание &ndash;&gt;-->
  <!--  <ToDoWidget-->
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
  <!--  <PhotoWidget />-->
</div>

<style>
  .container * {
    margin-bottom: 8px;
  }
</style>

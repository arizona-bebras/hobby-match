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
      order: 123,
      data: {
        type: 'Video',
        link: '1321',
        platform: 'YouTube',
      },
    },
    {
      telegram_id: '123',
      order: 123,
      data: {
        type: 'Video',
        link: '1321',
        platform: 'YouTube',
      },
    },
  ];

  let height = $state(380);
  $effect(() => {
    window.addEventListener(
      'scroll',
      () => (height = Math.max(300, 380 - window.scrollY * 0.5)),
    );
    console.log(window.scrollY);
  });
</script>

{#if height - 90 <= 210}
  <div class="w-full h-16 bg-white fixed flex z-2">
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
  <TextWidget
    text="„Люблю играть в Valorant. Часто говорят что выгляжу как будто сгенерирована нейросетью“"
  />
  <!-- Что делать с часами -->
  <GameWidget img="Valorant.png" game="Valorant" hour={1445} />
  <VideoWidget link="ValorantEdit.mp4" platform="" />
  <!-- Что делать с количеством подписчиков -->
  <SocialWidget
    platform="Youtube"
    url="https://www.youtube.com/@lowderplay/videos"
    username="@lowderplay"
    amountSubscribers={11}
  />
  <SocialWidget
    platform="VK"
    url="https://vk.com/id378488092"
    username="Илья Спицын"
    amountSubscribers={34}
  />
  <SocialWidget
    platform="Steam"
    url="https://steamcommunity.com/profiles/76561198295231108/"
    username="I'M SCHIZOPHRENIC AND HAVE A GUN "
    amountSubscribers={1}
  />
  <SocialWidget
    platform="Twitch"
    url="https://www.twitch.tv/shadowkekw"
    username="shadowkekw"
    amountSubscribers={42}
  />
  <SocialWidget
    platform="Twitter"
    url="https://x.com/elonmusk"
    username="elonmusk"
    amountSubscribers={12431}
  />
  <VideoWidget link="ValorantEdit2.mp4" platform="" />
  <!-- АА да -->
  <StickerWidget image="Stickers/Fire.svg" left={285} top={185} rotate={0} />
  <StickerWidget image="Stickers/Eyes.svg" left={15} top={315} rotate={-30} />
  <ProgressWidget label="Эйсы в Valorant" progress={62} max={100} />
  <!-- Сделать чтобы при True комнонент помечал как выполненное задание -->
  <ToDoWidget
    title="Что хочу сделать:)"
    tasks={[
      ['hello', true],
      ['world', true],
      [25, true],
      [50, true],
      [75, true],
      [100, false],
    ]}
  />
  <SurveyWidget
    tasks={[
      ['hello', 100],
      ['world', 25],
    ]}
    votes={125}
  />
  <PhotoWidget />
</div>

<style>
  .container * {
    margin-bottom: 8px;
  }
</style>

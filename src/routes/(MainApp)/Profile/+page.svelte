<script>
  import { Star } from '@lucide/svelte';
  import GameWidget from "./Widgets/GameWidget.svelte";
  import VideoWidget from "./Widgets/VideoWidget.svelte";
  import TextWidget from "./Widgets/TextWidget.svelte";
  import SocialWidget from "./Widgets/SocialWidget.svelte";
  import StickerWidget from "./Widgets/StickerWidget.svelte";
  import ProgressWidget from "./Widgets/ProgressWidget.svelte";
  import InterestsWidget from "./Widgets/InterestsWidget.svelte";
  import ToDoWidget from "./Widgets/ToDoWidget.svelte";

  let height = $state(380);
  $effect(() => {
      window.addEventListener('scroll', () => height = Math.max(300, 380 - window.scrollY * 0.5));
      console.log(window.scrollY);
  })
</script>


  {#if height - 90 <= 210}
    <div  class="w-full h-16 bg-white fixed flex z-2">
      <img src="Girl.png" class="w-16 h-16 rounded-full p-2">
      <div class="container font-[Inter] p-[4px]">
        <p class="font-extrabold text-[16px] flex items-center">Илона Абудаби, 18<Star /></p>
        <p class="font-semibold text-[16px]"><span>Махачкала, Россия</span></p>
      </div>
    </div>
  {/if}
  <img src="Girl.png" alt="person" class="w-full h-95 rounded-b-[24px] object-cover" style="height: {height}px">
  <div class="container font-[Inter] p-[16px] w-full max-w-full relative">
    <p class="font-extrabold text-[32px] flex items-center">Илона Абудаби, 18<Star /></p>
    <p class="font-semibold text-[20px]"><span>Екатеринбург, Россия</span></p>
    <TextWidget text='„Люблю играть в Valorant. Часто говорят что выгляжу как будто сгенерирована нейросетью“'/>
    <GameWidget img="Valorant.png" game="Valorant" hour={1445}/>
    <VideoWidget video="ValorantEdit.mp4"/>
    <SocialWidget type="Youtube" url="twitch.tv/ilonalala" amountSubscribers={11}/>
    <SocialWidget type="VK" url="twitch.tv/ilonalala" amountSubscribers={34}/>
    <SocialWidget type="Steam" url="twitch.tv/ilonalala" amountSubscribers={1}/>
    <SocialWidget type="Twitch" url="twitch.tv/ilonalala" amountSubscribers={42}/>
    <SocialWidget type="Twitter" url="twitch.tv/ilonalala" amountSubscribers={12431}/>
    <VideoWidget video="ValorantEdit2.mp4"/>
    <StickerWidget image="Stickers/Fire.svg" left={285} top={185} rotate={0} />
    <StickerWidget image="Stickers/Eyes.svg" left={15} top={315} rotate={-30}/>
    <ProgressWidget label="Эйсы в Valorant" progress={62} max={100}/>
    <InterestsWidget interests={[["🎮", "игры"], ["🚗", "развлечение"], ["✔️", "АХАХАХАХАХАХ"], ["🎮", "игры"], ["🎮", "игры"],["🎮", "игры"],["🎮", "игры"],["🎮", "игры"]]}/>
    <ToDoWidget tasks={["hello", "world", 25, 50, 75, 100]}/>
  </div>


<style>
  .container *{
      margin-bottom: 8px;
  }
</style>
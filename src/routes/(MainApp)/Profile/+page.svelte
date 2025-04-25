<script>
  import { Star } from '@lucide/svelte';
  import GameWidget from "../../../lib/components/widgets/GameWidget.svelte";
  import VideoWidget from "../../../lib/components/widgets/VideoWidget.svelte";
  import TextWidget from "../../../lib/components/widgets/TextWidget.svelte";
  import SocialWidget from "../../../lib/components/widgets/SocialWidget.svelte";
  import StickerWidget from "../../../lib/components/widgets/StickerWidget.svelte";
  import ProgressWidget from "../../../lib/components/widgets/ProgressWidget.svelte";
  import InterestsWidget from "../../../lib/components/widgets/InterestsWidget.svelte";
  import ToDoWidget from "../../../lib/components/widgets/ToDoWidget.svelte";
  import QuestionWidget from "../../../lib/components/widgets/QuestionWidget.svelte"; 

  let height = $state(380);
  $effect(() => {
      window.addEventListener('scroll', () => height = Math.max(300, 380 - window.scrollY * 0.5));
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
    <SocialWidget type="Youtube" url="https://www.youtube.com/@lowderplay/videos" userName="@lowderplay" amountSubscribers={11}/>
    <SocialWidget type="VK" url="https://vk.com/id378488092" userName="Илья Спицын" amountSubscribers={34}/>
    <SocialWidget type="Steam" url="https://steamcommunity.com/profiles/76561198295231108/" userName="I'M SCHIZOPHRENIC AND HAVE A GUN " amountSubscribers={1}/>
    <SocialWidget type="Twitch" url="https://www.twitch.tv/shadowkekw" userName="shadowkekw" amountSubscribers={42}/>
    <SocialWidget type="Twitter" url="https://x.com/elonmusk" userName="elonmusk" amountSubscribers={12431}/>
    <VideoWidget video="ValorantEdit2.mp4"/>
    <StickerWidget image="Stickers/Fire.svg" left={285} top={185} rotate={0} />
    <StickerWidget image="Stickers/Eyes.svg" left={15} top={315} rotate={-30}/>
    <ProgressWidget label="Эйсы в Valorant" progress={62} max={100}/>
    <InterestsWidget interests={[["🎮", "игры"], ["🚗", "развлечение"], ["✔️", "АХАХАХАХАХАХ"], ["🎮", "игры"], ["🎮", "игры"],["🎮", "игры"],["🎮", "игры"],["🎮", "игры"]]}/>
    <ToDoWidget tasks={["hello", "world", 25, 50, 75, 100]}/>
    <QuestionWidget tasks={[["hello", 100], ["world", 25]]} votes={125}/>
  </div>


<style>
  .container *{
      margin-bottom: 8px;
  }
</style>
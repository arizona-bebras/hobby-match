<script>
  import { Star } from '@lucide/svelte';
  import GameWidget from "./Widgets/GameWidget.svelte";
  import VideoWidget from "./Widgets/VideoWidget.svelte";
  import TextWidget from "./Widgets/TextWidget.svelte";
  import LinkWidget from "./Widgets/LinkWidget.svelte";
  import EmojiWidget from "./Widgets/EmojiWidget.svelte";
  import ProgressWidget from "./Widgets/ProgressWidget.svelte";
  let height = $state(380);
  $effect(() => {
      window.addEventListener('scroll', () => height = Math.max(300, 380 - window.scrollY * 0.5));
      console.log(window.scrollY);
  })
</script>


  {#if height - 90 <= 210}
    <div  class="w-[94%] h-16 bg-white fixed flex z-2">
      <img src="images.jfif" class="w-16 h-16 rounded-full">
      <div class="container font-[Inter] p-[4px]">
        <p class="font-extrabold text-[16px] flex items-center">Илона Абудаби, 18<Star /></p>
        <p class="font-semibold text-[16px]"><span>Махачкала, Россия</span></p>
      </div>
    </div>
  {/if}
  <img src="Herro.jpg" alt="person" class="w-full h-95 rounded-b-[24px] object-cover" style="height: {height}px">
  <div class="container font-[Inter] p-[16px] w-full max-w-full relative">
    <p class="font-extrabold text-[32px] flex items-center">Илона Абудаби, 18<Star /></p>
    <p class="font-semibold text-[20px]"><span>Махачкала, Россия</span></p>
    <TextWidget text='„Не тот велик, кто никогда не падал, а тот велик — кто падал и вставал!“'/>
    <GameWidget/>
    <GameWidget/>
    <VideoWidget video="upal-clip.mp4"/>
    <LinkWidget/>
    <VideoWidget video="NowYou.mp4"/>
    <EmojiWidget image="Fire.png" left={125} top={25} rotate={90} />
    <EmojiWidget image="Eyes.png" left={25} top={150} rotate={5}/>
    <ProgressWidget progress={62} max={100}/>
  </div>


<style>
  .container *{
      margin-bottom: 8px;
  }
</style>
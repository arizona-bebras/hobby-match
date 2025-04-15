<script>
  import {Star} from "@lucide/svelte";
  import { fade } from 'svelte/transition';

  let height = $state(300);
  $effect(() => {
      window.addEventListener('scroll', () => height = Math.max(235, 300 - window.scrollY * 0.5));
      console.log(height);
  })
  let scale = 1.0
  let border = 0;
</script>

<!--  <header class="magic transition-all">-->
<!--    <img src="testImage.jpg" alt="person" class="w-full h-95 rounded-b-[24px] object-cover justify-self-center">-->
<!--    <div class="container font-[Inter] p-[16px]">-->
<!--      <p class="font-extrabold text-[32px] flex items-center">Илона Абудаби, 18<Star /></p>-->
<!--      <p class="font-semibold text-[20px]"><span>Махачкала, Россия</span></p>-->
<!--  </header>-->

<!--  <header class="flex">-->
<!--    <img src="person.jpg" class="w-20 h-20 rounded-xl">-->
<!--    <div class="flex flex-col gap-2">-->
<!--      <p>hello</p>-->
<!--      <p>world</p>-->
<!--    </div>-->

<!--  </header>-->
<!--  {#if height <= 240}-->
<!--  {scale = 0.3}-->
<!--  {border = 100}-->
<!--  {:else if height >= 240}-->
<!--  {scale = 1}-->
<!--  {border = 0}-->
<!--  {:else}-->
<!--  {/if}-->
<!--  <div>-->

<!--  </div>-->

<img src="testImage.jpg" alt="person" class="w-full h-95 object-cover transition-[transform, border-radius] duration-500 ease-in-out top-0"
     style="height: {height}px;"/>
{#if height <= 235}
  <div transition:fade class="w-full h-16 bg-black fixed flex">
    <img src="images.jfif" class="w-16 h-16 rounded-full">
    <p>hello world</p>
  </div>
{/if}
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>
<p>hello</p>

<style>
.magic{
    display: flex;
    flex-direction: row;
    align-items: center;
}
.magic img{
    width: 50px;
    height: 50px;

}
</style>
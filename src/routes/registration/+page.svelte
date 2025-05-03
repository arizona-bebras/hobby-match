<script lang="ts">
  import Information from '$lib/components/registration/Information.svelte';
  import Photo from '$lib/components/registration/Photo.svelte';
  import Interests from '$lib/components/registration/Interests.svelte';
  import type { PageProps } from '../../../.svelte-kit/types/src/routes/registration/$types';
  let stages: string[][] = [
    ['information', 'Информация'],
    ['photo', 'Фото'],
    ['interests', 'Интересы'],
  ];
  let complitedStages: string[] = [];
  let currentStage: string = $state('information');
  window.Telegram.WebApp.MainButton.setParams({
    has_shine_effect: true,
    is_active: false,
    is_visible: true,
    text: 'Продолжить',
    color: '#808080',
  });
  let { data }: PageProps = $props();
</script>

<div class="p-4 w-full">
  <div class="flex">
    {#each stages as stage}
      <button
        onclick={() => {
          if (complitedStages.includes(stage[0])) currentStage = stage[0];
        }}
        class="border-t-2 w-30.75 p-2.5 {complitedStages.includes(stage[0])
          ? 'text-accent/50'
          : currentStage === stage[0]
            ? 'text-accent'
            : 'text-[#A7A7A7]'}"
        class:font-bold={currentStage === stage[0]}>{stage[1]}</button
      >
    {/each}
  </div>
  {#if currentStage === 'information'}
    <!--    <Information bind:currentStage />-->
    <Information form={data.information} />
    {window.Telegram.WebApp.MainButton.onClick(() => {
      currentStage = 'photo';
      complitedStages.push('information');
    })}
  {:else if currentStage === 'photo'}
    <Photo form={data.photo} />
    {window.Telegram.WebApp.MainButton.onClick(() => {
      currentStage = 'interests';
      complitedStages.push('photo');
    })}
  {:else if currentStage === 'interests'}
    <Interests form={data.interests} />
  {/if}
  <button onclick={() => console.log(complitedStages)}>ComplitedStages</button>
</div>

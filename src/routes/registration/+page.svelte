<script lang="ts">
  import Information from '$lib/components/registration/Information.svelte';
  import Photo from '$lib/components/registration/Photo.svelte';
  import Interests from '$lib/components/registration/Interests.svelte';
  import type { PageProps } from '../../../.svelte-kit/types/src/routes/registration/$types';
  import { pb } from '$lib/index';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  let stages: string[][] = [
    ['information', 'Информация'],
    ['photo', 'Фото'],
    ['interests', 'Интересы'],
  ];
  let completedStages: string[] = $state([]);
  let currentStage: string = $state('information');
  window.Telegram.WebApp.MainButton.setParams({
    has_shine_effect: true,
    is_active: false,
    is_visible: true,
    text: 'Продолжить',
    color: '#808080',
  });
  let { data }: PageProps = $props();

  function nextStage() {
    switch (currentStage) {
      case 'information':
        currentStage = 'photo';
        break;
      case 'photo':
        currentStage = 'interests';
        break;
      case 'interests':
        break;
    }

    if (!completedStages.includes(currentStage)) {
      completedStages.push(currentStage);
    }
  }
  $effect(() => {
    if (pb.authStore.record !== null) {
      completedStages = ['information', 'photo', 'interests'];
    }
  });
  $inspect(completedStages);
</script>

<div class="p-4 w-full">
  <div class="flex">
    {#each stages as stage}
      <button
        onclick={() => {
          if (completedStages.includes(stage[0])) currentStage = stage[0];
        }}
        class="border-t-2 w-30.75 p-2.5 {completedStages.includes(stage[0])
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
    <Information form={data.information} {nextStage} />
  {:else if currentStage === 'photo'}
    <Photo form={data.photo} {nextStage} />
  {:else if currentStage === 'interests'}
    <Interests form={data.interests} />
  {/if}
  <!--  <button onclick={() => console.log(complitedStages)}>ComplitedStages</button>-->
</div>

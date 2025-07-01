<script lang="ts">
  import Information from '$lib/components/registration/Information.svelte';
  import Photo from '$lib/components/registration/Photo.svelte';
  import Interests from '$lib/components/registration/Interests.svelte';
  import type { PageProps } from '../../../.svelte-kit/types/src/routes/registration/$types';
  import { pb } from '$lib/index';
  import { onMount } from 'svelte';
  import { Check } from '@lucide/svelte';
  import { goto } from '$app/navigation';

  let information: Information | undefined = $state();
  let photo: Photo | undefined = $state();
  let interests: Interests | undefined = $state();
  let stages: string[][] = [
    ['information', 'Информация'],
    ['photo', 'Фото'],
    ['interests', 'Интересы'],
  ];
  let completedStages: string[] = $state([]);
  let currentStage: string = $state('information');
  window.Telegram.WebApp.MainButton.setParams({
    is_active: false,
    is_visible: true,
    text: 'Продолжить',
    color: '#808080',
  });
  let { data }: PageProps = $props();

  function nextStage() {
    if (!completedStages.includes(currentStage)) {
      completedStages.push(currentStage);
    }
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
  }
  onMount(() => {
    if (
      pb.authStore.record?.miniapp_name &&
      pb.authStore.record?.gender &&
      pb.authStore.record?.birth_date &&
      pb.authStore.record?.location &&
      pb.authStore.record?.user_info
    ) {
      completedStages.push('information');
    }
    if (pb.authStore.record?.user_photo) {
      completedStages.push('photo');
    }
    if (pb.authStore.record?.interests.length >= 1) {
      completedStages.push('interests');
    }
    console.log(pb.authStore.record?.user_photo);
    console.log(pb.authStore.record?.interests);
  });
  $inspect(currentStage);
  $inspect(completedStages);
</script>

<div class="p-4 w-full">
  <div class="flex">
    {#each stages as stage}
      <button
        onclick={() => {
          if (completedStages.includes(stage[0])) currentStage = stage[0];
        }}
        class="border-t-2 w-1/3 p-2.5 {completedStages.includes(stage[0])
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
    <Information form={data.information} bind:this={information} {nextStage} />
  {:else if currentStage === 'photo'}
    <Photo form={data.photo} bind:this={photo} {nextStage} />
  {:else if currentStage === 'interests'}
    <Interests form={data.interests} bind:this={interests} />
  {/if}
  {#if completedStages.includes('information') && completedStages.includes('photo') && completedStages.includes('interests')}
    <button
      class="bg-accent size-12.5 fixed right-4 bottom-4 z-2 flex items-center justify-center rounded-xl"
      onclick={async () => {
        if (currentStage === 'information') await information?.save();
        else if (currentStage === 'photo') await photo?.save();
        else if (currentStage === 'interests') await interests?.save();
        await goto('/Profile');
      }}
    >
      <Check class="size-6 text-text-color" />
    </button>
  {/if}
  <!--  <button onclick={() => console.log(complitedStages)}>ComplitedStages</button>-->
</div>

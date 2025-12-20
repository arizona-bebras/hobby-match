<script lang="ts">
  import {
    Information,
    Photo,
    Interests,
    RegisterStages,
    Test,
    type Stages,
  } from '$lib/components/registration/index';

  import type { PageProps } from '../../../.svelte-kit/types/src/routes/registration/$types';
  import { pb } from '$lib';
  import { onMount } from 'svelte';
  import { Check } from '@lucide/svelte';
  import { goto } from '$app/navigation';
  import client from '$lib/api/client';

  let information: Information | undefined = $state();
  let photo: Photo | undefined = $state();
  let test: Test | undefined = $state();
  let interests: Interests | undefined = $state();

  window.Telegram.WebApp.MainButton.setParams({
    is_active: false,
    is_visible: true,
    text: 'Продолжить',
    color: '#808080',
  });
  let { data }: PageProps = $props();

  onMount(async () => {
    const { data } = await client.GET('/api/me');
    if (
      data?.miniapp_name &&
      data?.gender &&
      data?.birth_date &&
      data?.location &&
      data?.user_info
    ) {
      markStageComplete('Информация');
    }
    if (data?.user_photo) {
      markStageComplete('Фото');
    }
    if (data?.personality_test) {
      markStageComplete('Тест');
    }
    if (data?.interests) {
      markStageComplete('Интересы');
    }

    // if (
    //   pb.authStore.record?.miniapp_name &&
    //   pb.authStore.record?.gender &&
    //   pb.authStore.record?.birth_date &&
    //   pb.authStore.record?.location &&
    //   pb.authStore.record?.user_info
    // ) {
    //   markStageComplete('Информация');
    // }
    // if (pb.authStore.record?.user_photo) {
    //   markStageComplete('Фото');
    // }
    // if (pb.authStore.record?.interests.length >= 1) {
    //   markStageComplete('Интересы');
    // })
  });

  let stages = $state([
    {
      title: 'Информация',
      isCurrentStage: true,
      isComplete: false,
    },
    {
      title: 'Фото',
      isCurrentStage: false,
      isComplete: false,
    },
    {
      title: 'Тест',
      isCurrentStage: false,
      isComplete: false,
    },
    {
      title: 'Интересы',
      isCurrentStage: false,
      isComplete: false,
    },
  ]);

  function getCurrentStage() {
    return stages.find((element) => element.isCurrentStage);
  }

  function setCurrentStage(newTitle: string) {
    for (const stage of stages) {
      stage.isCurrentStage = stage.title === newTitle;
    }
  }

  function markStageComplete(stage: Stages) {
    if (stage === 'Информация') {
      stages[0].isComplete = true;
    } else if (stage === 'Фото') {
      stages[1].isComplete = true;
    } else if (stage === 'Тест') {
      stages[2].isComplete = true;
    } else {
      stages[3].isComplete = true;
    }
  }
  $inspect(stages);
</script>

<div class="p-4 w-full min-h-screen bg-background text-text-color">
  <RegisterStages bind:stages />
  {#if getCurrentStage()!.title === 'Информация'}
    <Information
      form={data.information}
      {setCurrentStage}
      {markStageComplete}
      bind:this={information}
    />
  {:else if getCurrentStage()!.title === 'Фото'}
    <Photo
      form={data.photo}
      {setCurrentStage}
      {markStageComplete}
      bind:this={photo}
    />
  {:else if getCurrentStage()!.title === 'Тест'}
    <Test {setCurrentStage} {markStageComplete} bind:this={test} />
  {:else}
    <Interests form={data.interests} {setCurrentStage} bind:this={interests} />
  {/if}

  {#if stages.every((obj) => obj.isComplete)}
    <button
      class="bg-accent size-12.5 fixed right-4 bottom-4 z-2 flex items-center justify-center rounded-xl"
      onclick={async () => {
        let currentStage = getCurrentStage()!.title;
        if (currentStage === 'Информация') await information?.save();
        else if (currentStage === 'Фото') await photo?.save();
        else if (currentStage === 'Интересы') await interests?.save();
        else if (currentStage === 'Тест') await test?.save();
        await goto('/profile');
      }}
    >
      <Check class="size-6 text-white" />
    </button>
  {/if}
</div>

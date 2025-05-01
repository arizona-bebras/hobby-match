<script lang="ts">
  import Information from '$lib/components/registration/Information.svelte';
  import Photo from '$lib/components/registration/Photo.svelte';
  import Interests from '$lib/components/registration/Interests.svelte';
  import type { PageProps } from '../../../.svelte-kit/types/src/routes/registration/$types';
  import { goto } from '$app/navigation';
  let stages: string[] = ['information', 'photo', 'interests'];
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
    <!--{#each stages as stage}-->
    <!--  <button-->
    <!--    class="border-t-2 w-30.75 p-2.5 {complitedStages.includes(stage)-->
    <!--      ? 'text-accent/50'-->
    <!--      : currentStage === stage-->
    <!--        ? 'text-accent'-->
    <!--        : 'text-[#A7A7A7]'}"-->
    <!--    class:font-bold={currentStage === stage}>Информация</button-->
    <!--  >-->
    <!--{/each}-->
    <button
      class="border-t-2 w-30.75 p-2.5 {complitedStages.includes('information')
        ? 'text-accent/50'
        : currentStage === 'information'
          ? 'text-accent'
          : 'text-[#A7A7A7]'}"
      class:font-bold={currentStage === 'information'}>Информация</button
    >
    <button
      class="border-t-2 w-30.75 p-2.5 {complitedStages.includes('photo')
        ? 'text-accent/50'
        : currentStage === 'photo'
          ? 'text-accent'
          : 'text-[#A7A7A7]'}"
      class:font-bold={currentStage === 'photo'}>Фото</button
    >
    <button
      class="border-t-2 w-30.75 p-2.5 {complitedStages.includes('interests')
        ? 'text-accent/50'
        : currentStage === 'interests'
          ? 'text-accent'
          : 'text-[#A7A7A7]'}"
      class:font-bold={currentStage === 'interests'}>Интересы</button
    >
  </div>
  {#if currentStage === 'information'}
    <!--    <Information bind:currentStage />-->
    <Information form={data.information} />
    {window.Telegram.WebApp.MainButton.onClick(() => {
      currentStage = 'photo';
    })}
  {:else if currentStage === 'photo'}
    <Photo form={data.photo} />
    {window.Telegram.WebApp.MainButton.onClick(() => {
      currentStage = 'interests';
    })}
    {complitedStages.push('information')}
  {:else if currentStage === 'interests'}
    <Interests form={data.interests} />
    {complitedStages.push('photo')}
  {/if}
</div>

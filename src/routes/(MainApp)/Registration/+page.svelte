<script lang="ts">
  import Information from '$lib/components/registration/Information.svelte';
  import Photo from '$lib/components/registration/Photo.svelte';
  import Interests from '$lib/components/registration/Interests.svelte';
  let stages: string[] = ['information', 'photo', 'interests'];
  let complitedStages: string[] = [];
  let currentStage: string = $state('information');
</script>

<div class="p-4 w-full">
  <div class="flex">
    <button
      class="border-t-2 w-30.75 p-2.5 {complitedStages.includes('information')
        ? 'text-accent/50'
        : currentStage === 'information'
          ? 'text-accent'
          : 'text-[#A7A7A7]'}"
      class:font-bold={currentStage === 'information'}>Информация</button
    >

    <!--    if (complitedStages.includes('information')) = text-accent/25-->
    <!--    elif currentStage  === 'information') = text-accent-->
    <!--    else -->

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
    <Information bind:currentStage />
  {:else if currentStage === 'photo'}
    <Photo bind:currentStage />
    {complitedStages.push('information')}
  {:else if currentStage === 'interests'}
    <Interests />
    {complitedStages.push('photo')}
  {/if}
  <button
    onclick={() => {
      console.log(currentStage);
    }}>checkComplitedStages</button
  >
</div>

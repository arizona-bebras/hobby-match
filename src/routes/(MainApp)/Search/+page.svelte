<script lang="ts">
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  import { goto } from '$app/navigation';
  import Questionnaire from '$lib/components/profile/Questionnaire.svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import { onMount } from 'svelte';
  import { ArrowUp } from '@lucide/svelte';
  import { pb } from '$lib';

  let isMousePress = $state(false);
  let container: HTMLDivElement | undefined = $state();
  let offeredProfiles: PageData[] = $state([]);
  let currentProfile = $state(0);
  useTelegramButton(() => goto('/Profile'));

  // let offeredProfiles

  onMount(() => {
    window.Telegram.WebApp.MainButton.setText(
      // changeMode ? 'Сохранить' : 'Изменить виджеты',
      'Моя анкета',
    );
    // getProfiles();
  });
  let elementSize = $state(0);
  let isProfileEnd = $state(false);
  let transitionScroll = $state(0);
  // let scrollY = $derived(container.scrollTop);
  // $inspect(scrollY);
  $effect(() => {
    if (elementSize >= 80 && !isMousePress) {
      elementSize = 0;
      container?.scrollTo(0, 0);
      console.log('Опа! Загружаем новую страницу');
      currentProfile += 1;
      offeredProfiles.shift();
    }
  });
  $effect(() => {
    if (offeredProfiles.length <= 2) {
      getProfiles().then(
        (response) => (offeredProfiles = [...offeredProfiles, ...response]),
      );
    }

    // if (amountViewedProfiles >= 2) {
    //   amountViewedProfiles = 0;
    //   console.log(4455);
    // }
  });
  async function getProfiles(): Promise<PageData[]> {
    const response = await pb.send('/worker/feed', {
      method: 'GET',
    });
    // console.log(response);
    // offeredProfiles = response;
    console.log(response);
    return response;
  }

  $inspect(isProfileEnd);
</script>

<div
  class="overflow-y-auto"
  bind:this={container}
  onmousedown={() => (isMousePress = true)}
  onmouseup={() => (isMousePress = false)}
  ontouchstart={() => (isMousePress = true)}
  ontouchend={() => (isMousePress = false)}
  role="button"
  tabindex="0"
  onscroll={() => {
    // console.log('СКРОЛ)');
    if (isProfileEnd && isMousePress) {
      // Нужно какое-то условие, которое выполнит этот код 1 раз
      if (transitionScroll === 0) {
        transitionScroll = container.scrollTop;
        //console.log('ВЫПОЛНИЛОСЬ!');
      }
      // console.log((container.scrollTop - doubleScroll) * 1.2, elementSize);
      elementSize = Math.max(0, (container.scrollTop - transitionScroll) * 1.2);
      console.log(container.scrollTop - transitionScroll);

      // elementSize += ;
      // elementSize = Math.min(80, elementSize);

      //   if (!isMousePress) {
      //     console.log(4455);
      //   }
      // } else {
      //   elementSize = 0;
      // }
      // else {
      //     doubleScroll = container.scrollTop;
    } else {
      console.log('Обнуление размеров компонента');
      transitionScroll = 0;
    }

    isProfileEnd =
      container?.children[0].offsetHeight - document.body.offsetHeight <=
      container?.scrollTop - 8;

    // console.log(elementSize);
    // console.log(
    //   container.children[0].offsetHeight - document.body.offsetHeight,
    //   isProfileEnd,
    //   container.scrollTop - 8,
    // );
  }}
>
  <!--  <button-->
  <!--    class="bg-purple-400 fixed"-->
  <!--    onclick={() => {-->
  <!--      doubleScroll = container.scrollTop;-->
  <!--    }}-->
  <!--  >-->
  <!--    Начать отслеживание-->
  <!--  </button>-->
  {#if offeredProfiles.length <= 0}
    <p>Загрузка</p>
  {:else}
    <p>Загрузилось {offeredProfiles}</p>
    <Questionnaire data={offeredProfiles[0]} />

    {#if isProfileEnd && isMousePress}
      <div
        class="bg-accent/25 max-w-15 max-h-27 rounded-full mx-auto"
        style:width="{elementSize}px"
        style:height="{elementSize}px"
      >
        <p style:font-size="min(30px, {elementSize}px)" class="text-center">
          &#8593;
        </p>
        <img
          class="rounded-full p-1 aspect-square object-cover"
          src={pb.buildURL(
            `/api/files/_pb_users_auth_/${offeredProfiles[1].id}/${offeredProfiles[1].user_photo}`,
          )}
          alt="userImage"
        />
        <!--        <p style:font-size="min(20px, {elementSize}px)">Testdasdasdsadasdas</p>-->
      </div>
    {/if}
  {/if}
  <!--  <div>-->
  <!--    <Questionnaire data={offeredProfiles[0]} />-->
  <!--  </div>-->
</div>

<!--<div-->
<!--  class="snap-y snap-mandatory w-full h-full overflow-y-auto gap-y-40"-->
<!--  bind:this={container}-->
<!--  onscroll={() => {-->
<!--    updateCurrentSection();-->
<!--  }}-->
<!--&gt;-->
<!--  <button-->
<!--    onclick={() => {-->
<!--      profiles.splice(0, 1);-->
<!--      profiles.push(thirdProfile);-->
<!--      console.log(profiles);-->
<!--    }}-->
<!--    class="fixed size-25 bg-purple-500">123</button-->
<!--  >-->
<!--  <button-->
<!--    onclick={() => {-->
<!--      // profiles.push(offeredProfiles.shift());-->
<!--      // profiles.splice(0, 1);-->
<!--      // currentProfile = 0;-->
<!--      // container.scrollTop = 0;-->
<!--      console.log(profiles);-->
<!--      // profileRef.forEach((element: HTMLDivElement) =>-->
<!--      //   console.log(element.offsetHeight),-->
<!--      // );-->
<!--    }}-->
<!--    class="fixed right-5 size-25 bg-purple-500 z-100 flex">Check state</button-->
<!--  >-->
<!--  <div class="fixed size-25 bg-purple-500">-->
<!--    Текущий профиль: {currentProfile}-->
<!--  </div>-->

<!--{#each profiles as profile, i}-->
<!--  <div class="snap-start" bind:this={profileRef[i]}>-->
<!--    <Questionnaire data={profile} />-->
<!--  </div>-->
<!--  &lt;!&ndash;    <div class="w-40 h-40 bg-purple-500 snap-always"></div>&ndash;&gt;-->
<!--{/each}-->
<!--</div>-->

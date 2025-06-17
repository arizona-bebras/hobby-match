<script lang="ts">
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  import { goto } from '$app/navigation';
  import Questionnaire from '$lib/components/profile/Questionnaire.svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import { onMount } from 'svelte';
  import { ArrowUp } from '@lucide/svelte';
  import { pb } from '$lib';

  let isMousePress = $state(false);
  let profileContainer: HTMLDivElement | undefined = $state();
  let screenContainer: HTMLDivElement | undefined = $state();
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
  let touchStartPosition: { x: number; y: number } | null = $state(null);
  let transitionScroll = $state(0);
  // let scrollY = $derived(container.scrollTop);
  // $inspect(scrollY);
  $effect(() => {
    if (touchStartPosition === null) {
      if (elementSize >= 108) {
        screenContainer?.scrollTo(0, 0);
        console.log('Опа! Загружаем новую страницу');
        currentProfile += 1;
        offeredProfiles.shift();
      }
      elementSize = 0;
    }
  });
  $effect(() => {
    if (offeredProfiles.length <= 2) {
      getProfiles().then(
        (response) => (offeredProfiles = [...offeredProfiles, ...response]),
      );
    }
  });
  $effect(() => {
    console.log('AAAAAAAAAAAAAAAAAAAAAAAAAAAA');
    if (
      profileContainer &&
      screenContainer &&
      profileContainer.offsetHeight <= screenContainer.offsetHeight
    ) {
      isProfileEnd = true;
    }
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
  $inspect(touchStartPosition);
</script>

<div
  class="overflow-y-auto w-full"
  bind:this={screenContainer}
  ontouchstart={(e) => {
    console.log(e, isProfileEnd);
    if (isProfileEnd) {
      touchStartPosition = {
        x: e.changedTouches[0].clientX,
        y: e.changedTouches[0].clientY,
      };
    }
  }}
  ontouchend={() => {
    touchStartPosition = null;
  }}
  ontouchmove={(e) => {
    if (touchStartPosition) {
      const delta = (touchStartPosition?.y - e.changedTouches[0].clientY) * 0.5;
      if (delta > 0) {
        elementSize = delta;
        screenContainer?.scrollTo(0, screenContainer?.scrollHeight);
      }
    }
  }}
  onwheel={(e) => {
    if (isProfileEnd) {
      if (e.deltaY > 0) {
        elementSize += e.deltaY / 10;
        if (elementSize > 108) {
          touchStartPosition = null;
        } else {
          touchStartPosition = { x: 0, y: 0 };
        }
      } else {
        elementSize = 0;
        touchStartPosition = null;
      }

      screenContainer?.scrollTo(0, screenContainer?.scrollHeight);
      console.log(elementSize, e.deltaY / 10);
    } else {
      elementSize = 0;
      touchStartPosition = null;
    }
  }}
  onscroll={() => {
    if (!screenContainer || !profileContainer) {
      return;
    }
    console.log(111);
    isProfileEnd =
      screenContainer.scrollTop + screenContainer.offsetHeight >=
      profileContainer.offsetHeight;
  }}
>
  {#if offeredProfiles.length <= 0}
    <p>Загрузка</p>
  {:else}
    <div bind:this={profileContainer}>
      <Questionnaire data={offeredProfiles[0]} />
    </div>
    {#if isProfileEnd}
      <div
        class="bg-accent/25 max-w-15 max-h-27 rounded-full mx-auto overflow-hidden"
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
      </div>
    {/if}
  {/if}
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

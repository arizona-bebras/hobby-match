<script lang="ts">
  import { ProgressRing } from '@skeletonlabs/skeleton-svelte';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  import { goto } from '$app/navigation';
  import Questionnaire from '$lib/components/profile/Questionnaire.svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import { onMount } from 'svelte';
  import { ArrowUp } from '@lucide/svelte';
  import { pb } from '$lib';
  import { fade, scale, slide, fly, draw, blur } from 'svelte/transition';
  import { Tween } from 'svelte/motion';
  import { cubicOut, quintOut } from 'svelte/easing';
  import { elasticOut } from 'svelte/easing';
  import { Plus, Pencil, Save, Heart, HeartOff } from '@lucide/svelte';
  import LikeButton from '$lib/components/search/LikeButton.svelte';

  let isMousePress = $state(false);
  let profileContainer: HTMLDivElement | undefined = $state();
  let screenContainer: HTMLDivElement | undefined = $state();
  let offeredProfiles: PageData[] = $state([]);
  let currentProfile = $state(0);
  let liked: boolean = $state(false);
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
  // let elementSize = new Tween(0, {
  //   duration: 400,
  //   easing: cubicOut,
  // });
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
  let hapticAvailable = $state(true);
  let hapticDisable = $state(false);
  $effect(() => {
    if (elementSize >= 108 && hapticAvailable) {
      window.Telegram.WebApp.HapticFeedback.impactOccurred('light');
      hapticAvailable = false;
      hapticDisable = false;
    } else if (elementSize <= 105 && elementSize >= 20 && !hapticDisable) {
      window.Telegram.WebApp.HapticFeedback.impactOccurred('light');
      hapticAvailable = true;
      hapticDisable = true;
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
    if (
      profileContainer &&
      screenContainer &&
      profileContainer.offsetHeight <= screenContainer.offsetHeight
    ) {
      isProfileEnd = true;
    }
  });
  // $effect(() => {
  //   if (liked && touchStartPosition !== null) {
  //     console.log("LIKEDLIKEDLIKEDLIKED")
  //     pb.collection('likes').create({
  //       user: pb.authStore.record!.id,
  //       liked_user: offeredProfiles[0].id
  //     })
  //   }
  // });
  async function getProfiles(): Promise<PageData[]> {
    const response = await pb.send('/worker/feed', {
      method: 'GET',
    });
    // console.log(response);
    // offeredProfiles = response;
    console.log(response);
    return response;
  }
  let isMoving = $state(false);
  // $inspect(elementSize);
  $inspect(elementSize);

  function close(node: HTMLDivElement, { duration }: { duration: number }) {
    const startWidth = node.offsetWidth;
    const startHeight = node.offsetHeight;

    return {
      duration,
      css: (t) => {
        const eased = quintOut(t);
        return `
          width: ${startWidth * t}px;
          height: ${startHeight * t}px;
          opacity: ${t};
          overflow: hidden;
        `;
      },
    };
  }
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
    isMoving = false;
    touchStartPosition = null;
  }}
  ontouchmove={(e) => {
    isMoving = true;
    if (touchStartPosition) {
      const delta = (touchStartPosition?.y - e.changedTouches[0].clientY) * 0.5;
      if (delta > 0) {
        elementSize = delta;
        screenContainer?.scrollTo(0, screenContainer?.scrollHeight);
      }
    }
  }}
  onwheel={(e) => {
    isMoving = true;
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
    <div
      class="flex flex-col items-center justify-center w-screen h-screen font-[Inter] text-lg font-medium"
    >
      <ProgressRing
        value={null}
        meterStroke="stroke-accent"
        trackStroke="stroke-accent/25"
      ></ProgressRing>
      <p class="pt-4 text-text-color">Ищем подходящие профили...</p>
    </div>
  {:else}
    <button
      onclick={() => {
        if (!liked) {
          pb.collection('likes').create({
            user: pb.authStore.record!.id,
            liked_user: offeredProfiles[0].id
          })
        }
        liked = !liked
        console.log("LIKEDLIKEDLIKEDLIKED")
        }}
      class="bg-accent size-12.5 fixed right-6.5 bottom-5 z-2 flex items-center justify-center rounded-xl"
    >
      {#if !liked}
        <Heart class="size-6 text-text-color" />
      {:else}
        <HeartOff class="size-6 text-text-color" />
      {/if} 
    </button>
   
    <div bind:this={profileContainer}>
      <Questionnaire data={offeredProfiles[0]} />
    </div>
    {#if isProfileEnd && isMoving}
      <div
        out:close={{ duration: 200 }}
        class="bg-accent/25 max-w-15 max-h-27 rounded-full mx-auto overflow-hidden"
        style:width="{elementSize}px"
        style:height="{elementSize}px"
      >
        <p style:font-size="min(30px, {elementSize}px)" class="text-center">
          &#8593;
        </p>
        {#if elementSize >= 35}
          <img
            class="rounded-full p-1 aspect-square object-cover max-w-15 max-h-15 mx-auto"
            style:width="{elementSize - 35}px"
            style:height="{elementSize - 35}px"
            src={pb.buildURL(
              `/api/files/_pb_users_auth_/${offeredProfiles[1].id}/${offeredProfiles[1].user_photo}`,
            )}
            alt="userImage"
          />
        {/if}
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

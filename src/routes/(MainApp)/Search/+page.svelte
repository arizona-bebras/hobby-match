<script lang="ts">
  import { ProgressRing } from '@skeletonlabs/skeleton-svelte';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  import { goto } from '$app/navigation';
  import Questionnaire from '$lib/components/profile/Questionnaire.svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import { onMount } from 'svelte';
  import { pb } from '$lib';
  import { fly } from 'svelte/transition';
  import { Heart, LoaderCircle } from '@lucide/svelte';
  import { plausible } from '../../../hooks.client';

  let { data }: { data: { page?: PageData } } = $props();

  let profileContainer: HTMLDivElement | undefined = $state();
  let screenContainer: HTMLDivElement | undefined = $state();
  let offeredProfiles: PageData[] = $state(data.page ? [data.page] : []);
  let currentProfile = $state(0);
  useTelegramButton(async () => {
    window.Telegram.WebApp.MainButton.showProgress();
    await goto('/Profile');
    window.Telegram.WebApp.MainButton.hideProgress();
  });
  let liked: boolean = $state(false);
  useTelegramButton(() => goto('/Profile'));

  onMount(() => {
    window.Telegram.WebApp.MainButton.setText(
      // changeMode ? 'Сохранить' : 'Изменить виджеты',
      'Моя анкета',
    );
    window.Telegram.WebApp.MainButton.show();
    // getProfiles();
  });
  let elementSize = $state(0);
  // let elementSize = new Tween(0, {
  //   duration: 400,
  //   easing: cubicOut,
  // });
  let isProfileEnd = $state(false);
  let touchStartPosition: { x: number; y: number } | null = $state(null);
  // let scrollY = $derived(container.scrollTop);
  // $inspect(scrollY);
  $effect(() => {
    if (touchStartPosition === null) {
      if (elementSize >= 108) {
        screenContainer?.scrollTo(0, 0);
        console.log('Опа! Загружаем новую страницу');
        currentProfile += 1;
        offeredProfiles.shift();
        liked = false;
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
    if (
      profileContainer &&
      screenContainer &&
      profileContainer.offsetHeight <= screenContainer.offsetHeight
    ) {
      isProfileEnd = true;
    }
  });

  $effect(() => {
    const timeout = setTimeout(
      () => {
        if (offeredProfiles.length > 2) return;
        pb.send('/worker/feed', {
          method: 'GET',
        }).then(
          (response) => (offeredProfiles = [...offeredProfiles, ...response]),
        );
      },
      offeredProfiles.length == 0 ? 0 : 800,
    );

    return () => clearTimeout(timeout);
  });
  let isMoving = $state(false);
  // $inspect(elementSize);
  $inspect(elementSize);

  function close(node: HTMLDivElement, { duration }: { duration: number }) {
    const startWidth = node.offsetWidth;
    const startHeight = node.offsetHeight;

    return {
      duration,
      css: (t: number) => {
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
  class="overflow-y-auto w-full min-h-full"
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
    {#key currentProfile}
      <div
        in:fly={{ duration: 500, y: 200 }}
        bind:this={profileContainer}
        class="min-h-full"
      >
        <Questionnaire data={offeredProfiles[0]} />
      </div>
    {/key}
    <button
      onclick={() => {
        if (!liked) {
          pb.collection('likes').create({
            user: pb.authStore.record?.id,
            liked_user: offeredProfiles[0].id,
          });
          plausible.trackEvent('liked');
        }
        liked = !liked;
        console.log('LIKE');
      }}
      disabled={liked}
      class="bg-accent size-12.5 fixed right-6.5 bottom-5 z-2 flex items-center justify-center rounded-xl"
    >
      {#if !liked}
        <Heart class="size-6 text-white" />
      {:else}
        <Heart
          fill="#fff"
          strokeWidth={0}
          class="size-6 text-white animate-ping"
          style="animation-iteration-count: 2; animation-direction: alternate; animation-duration: 400ms"
        />
      {/if}
    </button>

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
          {#if offeredProfiles.length > 1}
            <img
              class="rounded-full p-1 aspect-square object-cover max-w-15 max-h-15 mx-auto"
              style:width="{elementSize - 35}px"
              style:height="{elementSize - 35}px"
              src={pb.buildURL(
                `/api/files/_pb_users_auth_/${offeredProfiles[1]?.id}/${offeredProfiles[1]?.user_photo}?thumb=350x0`,
              )}
              alt="userImage"
            />
          {:else}
            <LoaderCircle class="animate-spin w-full" />
          {/if}
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

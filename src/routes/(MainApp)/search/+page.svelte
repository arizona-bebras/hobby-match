<script lang="ts">
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  import { goto } from '$app/navigation';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import { onMount } from 'svelte';
  import { db } from '$lib'
  import LoadingScreen from '$lib/components/search/LoadingScreen.svelte';
  import TransitionBlock from '$lib/components/search/TransitionBlock.svelte';
  import UserProfile from '$lib/components/search/UserProfile.svelte';

  let { data }: { data: { page?: PageData } } = $props();

  let profileContainer: HTMLDivElement | undefined = $state();
  let screenContainer: HTMLDivElement | undefined = $state();
  let offeredProfiles: PageData[] = $state(data.page ? [data.page] : []);
  let currentProfile = $state(0);
  useTelegramButton(async () => {
    window.Telegram.WebApp.MainButton.showProgress();
    await goto('/profile');
    window.Telegram.WebApp.MainButton.hideProgress();
  });
  let liked: boolean = $state(false);
  useTelegramButton(() => goto('/profile'));

  const authHeader: HeadersInit = new Headers()
  authHeader.set('Authorization', `Bearer ${window.localStorage.getItem("access_token")}`)

  onMount(() => {
    window.Telegram.WebApp.MainButton.setText('Моя анкета');
    window.Telegram.WebApp.MainButton.show();
  });
  let elementSize = $state(0);
  let isProfileEnd = $state(false);
  let touchStartPosition: { x: number; y: number } | null = $state(null);
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
        fetch (`${db}/api/namespace/00000000-0000-0000-0000-000000000000`, {
          method: "GET",
          headers: authHeader,
        }).then(
            (response) => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            }
        ).then(
            (data) => {
                offeredProfiles = [...offeredProfiles, ...data];
            }
        ).catch(
            (error) => {
                console.error("Fetch error:", error);
            }
        );
        // pb.send('/worker/feed', {
        //   method: 'GET',
        // }).then(
        //   (response) => (offeredProfiles = [...offeredProfiles, ...response]),
        // );
      },
      offeredProfiles.length == 0 ? 0 : 800,
    );

    return () => clearTimeout(timeout);
  });
  let isMoving = $state(false);
</script>

<div
  class="overflow-y-auto w-full min-h-full"
  bind:this={screenContainer}
  ontouchstart={(e) => {
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
    } else {
      elementSize = 0;
      touchStartPosition = null;
    }
  }}
  onscroll={() => {
    if (!screenContainer || !profileContainer) {
      return;
    }
    isProfileEnd =
      screenContainer.scrollTop + screenContainer.offsetHeight >=
      profileContainer.offsetHeight;
  }}
>
  {#if offeredProfiles.length <= 0}
    <LoadingScreen />
  {:else}
    <UserProfile
      {currentProfile}
      {profileContainer}
      {offeredProfiles}
      bind:liked
    />

    {#if isProfileEnd && isMoving}
      <TransitionBlock {elementSize} {offeredProfiles} />
    {/if}
  {/if}
</div>

<!--<div-->
<!-- class="snap-y snap-mandatory w-full h-full overflow-y-auto gap-y-40"-->
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
<!--      // currentProfile = 0; -->
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

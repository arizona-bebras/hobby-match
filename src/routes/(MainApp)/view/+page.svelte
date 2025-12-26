<script lang="ts">
  import { page } from '$app/state';
  import { createQuery } from '@tanstack/svelte-query';
  import client from '$lib/api/client';
  import UserProfile from '$lib/components/search/UserProfile.svelte';
  import LoadingScreen from '$lib/components/search/LoadingScreen.svelte';
  import TransitionBlock from '$lib/components/search/TransitionBlock.svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes.ts';
  import { goto } from '$app/navigation';
  import { untrack } from 'svelte';

  let namespaceId = page.url.searchParams.get('namespaceId');
  const profileData = createQuery(() => ({
    queryKey: ['profileData1'],
    queryFn: async () =>
      await client.GET('/api/namespace/{namespace_id}/feed', {
        params: {
          path: {
            namespace_id: namespaceId,
          },
        },
      }),
    select: (data) => data.data,
    gcTime: 0,
    staleTime: 0,
  }));
  let profileContainer: HTMLDivElement | undefined = $state();

  let screenContainer: HTMLDivElement | undefined = $state();
  let offeredProfiles: PageData[] = $state([]);
  let currentProfile = $state(0);

  let liked: boolean = $state(false);

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
  $effect(() => {
    if (profileData.isSuccess) {
      untrack(() => offeredProfiles.push(...profileData.data));
    }
  });
  let hapticAvailable = $state(true);
  let hapticDisable = $state(false);
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
    if (offeredProfiles.length < 2) {
      profileData.refetch();
    }
    console.log(offeredProfiles);
    currentProfile;
  });
  // $effect(() => {
  //   const timeout = setTimeout(
  //     () => {
  //       if (offeredProfiles.length > 2) return;
  //       fetch(`${db}/api/namespace/${namespaceId}/feed`, {
  //         method: 'GET',
  //         headers: authHeader,
  //       })
  //         .then((response) => {
  //           if (!response.ok) {
  //             throw new Error(`HTTP error! status: ${response.status}`);
  //           }
  //           return response.json();
  //         })
  //         .then((data) => {
  //           offeredProfiles = [...offeredProfiles, ...data];
  //         })
  //         .catch((error) => {
  //           console.error('Fetch error:', error);
  //         });
  //     },
  //     offeredProfiles.length == 0 ? 0 : 800,
  //   );
  //
  //   return () => clearTimeout(timeout);
  // });
  $inspect(isProfileEnd);
  let isMoving = $state(false);
</script>

{#if profileData.isSuccess}
  <div
    class="overflow-y-auto w-full"
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
        const delta =
          (touchStartPosition?.y - e.changedTouches[0].clientY) * 0.5;
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
        bind:profileContainer
        {offeredProfiles}
        bind:liked
      />

      {#if isProfileEnd && isMoving}
        <TransitionBlock {elementSize} {offeredProfiles} />
      {/if}
    {/if}
  </div>
{/if}

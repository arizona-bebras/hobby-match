<script lang="ts">
  import { fly } from 'svelte/transition';
  import Questionnaire from '$lib/components/profile/Questionnaire.svelte';
  import { Heart } from '@lucide/svelte';
  import { pb } from '$lib';
  import { plausible } from '../../../hooks.client.js';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';

  let {
    currentProfile,
    profileContainer,
    offeredProfiles,
    liked = $bindable(),
  }: {
    currentProfile: number;
    profileContainer: HTMLDivElement;
    offeredProfiles: PageData[];
    liked: boolean;
  } = $props();
</script>

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

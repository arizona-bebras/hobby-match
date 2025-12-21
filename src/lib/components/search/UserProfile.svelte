<script lang="ts">
  import { fly } from 'svelte/transition';
  import Questionnaire from '$lib/components/profile/Questionnaire.svelte';
  import { Heart } from '@lucide/svelte';
  import { pb } from '$lib';
  import { plausible } from '../../../hooks.client.js';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';

  let {
    currentProfile,
    profileContainer = $bindable(),
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

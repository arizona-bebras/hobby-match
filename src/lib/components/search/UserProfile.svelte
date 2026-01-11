<script lang="ts">
  import { fly } from 'svelte/transition';
  import Questionnaire from '$lib/components/profile/Questionnaire.svelte';
  import { Heart } from '@lucide/svelte';
  import { plausible } from '../../../hooks.client.js';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import WriteUserBtn from '$lib/components/search/WriteUserBtn.svelte';
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
  let profile = $derived.by(() => {
    let pageData: PageData = JSON.parse(JSON.stringify(offeredProfiles[0]!));
    for (const widget of pageData.widgets) {
      widget.data = JSON.parse(widget.data!);
      if (widget.additionalData != '') {
        widget.additionalData = JSON.parse(widget.additionalData!);
      }
    }
    pageData!.widgets = pageData?.widgets?.sort((a, b) => {
      return a.order! - b.order!;
    });
    return pageData;
  });
</script>

{#key currentProfile}
  <div
    in:fly={{ duration: 500, y: 200 }}
    bind:this={profileContainer}
    class="min-h-full"
  >
    <Questionnaire data={profile} isNamespaceProfile={true} />
  </div>
  <WriteUserBtn username={profile.username} />
{/key}

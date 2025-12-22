<script lang="ts">
  import InterestsWidget from '$lib/components/widgets/InterestsWidget.svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import ReportButton from '$lib/components/profile/ReportButton.svelte';
  import Zodiac from '$lib/components/profile/Zodiac.svelte';
  import type { ScrollState } from 'runed';
  import Similarity from '$lib/components/profile/Similarity.svelte';
  import TestResultBtn from '$lib/components/profile/TestResultBtn.svelte';

  let {
    data,
    changeMode = false,
    isNamespaceProfile = false,
    scroll,
  }: {
    data: PageData;
    changeMode?: boolean;
    isNamespaceProfile?: boolean;
    scroll?: ScrollState;
  } = $props();

  // let userDescription: Text = {
  //   type: 'text',
  //   text: pb.authStore.record!.user_info,
  // };
  console.log(data);
</script>

{#if !changeMode}
  <div class="font-extrabold text-[32px] mt-2 wrap-anywhere flex gap-1">
    <span>{data.miniapp_name}, {data.age}</span>
    <!--    <Zodiac />-->
    <!-- TODO: fix id -->
    {#if '' !== data.id}
      <ReportButton offender={data.id} />
    {/if}
  </div>
  <p class="font-semibold text-[20px] break-words">
    <span>{data.location}</span>
  </p>
  {#if isNamespaceProfile}
    <TestResultBtn userSimilarity={50} interests={data.interests} />
  {/if}
{/if}

<InterestsWidget interests={data.interests} />
{#if !isNamespaceProfile}
  <div class="TextBox text-container my-2">
    <b>О себе: </b>
    <p class="break-words italic">{data.user_info}</p>
  </div>
{/if}

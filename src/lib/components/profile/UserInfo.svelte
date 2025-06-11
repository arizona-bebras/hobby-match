<script lang="ts">
  import InterestsWidget from '$lib/components/widgets/InterestsWidget.svelte';
  import { pb } from '$lib';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  let { data, changeMode }: { data: PageData; changeMode: boolean } = $props();
  // let userDescription: Text = {
  //   type: 'text',
  //   text: pb.authStore.record!.user_info,
  // };
</script>

{#if !changeMode}
  <p class="font-extrabold text-[32px] flex items-center mt-2">
    {data.miniapp_name}, {data.age}
  </p>
  <p class="font-semibold text-[20px]"><span>{data.location}</span></p>
{/if}
<InterestsWidget interests={data.interests} />
<div class="TextBox text-container my-2">
  <b>О себе: </b>
  <p class="break-words italic">{pb.authStore.record?.user_info}</p>
</div>

<script lang="ts">
  import InterestsWidget from '$lib/components/widgets/InterestsWidget.svelte';
  import TextWidget from '$lib/components/widgets/TextWidget.svelte';
  import type { Text } from '$lib/widgetTypes/widgetTypes';
  import { pb } from '$lib';
  let { data, changeMode }: { data: any; changeMode: boolean } = $props();
  let userDescription: Text = {
    type: 'text',
    text: pb.authStore.record!.user_info,
  };
</script>

{#if !changeMode}
  <p class="font-extrabold text-[32px] flex items-center">
    {data.miniapp_name}, {data.age}
  </p>
  <p class="font-semibold text-[20px]"><span>{data.location}</span></p>
  <InterestsWidget interests={data.interests.expand?.interests} />
  <TextWidget data={userDescription} />
{/if}

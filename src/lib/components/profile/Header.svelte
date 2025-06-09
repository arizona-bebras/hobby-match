<script lang="ts">
  import ProfileTopLayer from '$lib/components/profile/ProfileTopLayer.svelte';
  import { pb } from '$lib';
  import InterestsWidget from '$lib/components/widgets/InterestsWidget.svelte';
  let {
    height,
    data,
    changeMode,
  }: { height: number; data: any; changeMode: boolean } = $props();
</script>

{#if height - 90 <= 210 || changeMode}
  <ProfileTopLayer {data} fixed={!changeMode} />
{/if}
{#if !changeMode}
  <img
    src={pb.files.getURL(pb.authStore.record ?? {}, data.user_photo)}
    alt="person"
    class="w-full h-95 rounded-b-[24px] object-cover"
    style="height: {height}px"
  />
{/if}

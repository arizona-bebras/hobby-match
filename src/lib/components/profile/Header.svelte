<script lang="ts">
  import ProfileTopLayer from '$lib/components/profile/ProfileTopLayer.svelte';
  import { pb } from '$lib';
  import { scrollY } from 'svelte/reactivity/window';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  let { data, changeMode }: { data: PageData; changeMode: boolean } = $props();

  let image: HTMLImageElement | undefined = $state();
  let width = $derived(image?.width ?? 300);
  let height = $derived(
    Math.max(width - 100, width - (scrollY.current ?? 0) * 0.5),
  );
  let topLayer: boolean = $derived((scrollY.current ?? 0) > height);
  let topLayerContainer: HTMLDivElement | undefined = $state();
</script>

{#if topLayer || changeMode}
  <ProfileTopLayer
    {data}
    shadow={!changeMode}
    bind:container={topLayerContainer}
  />
{/if}
{#if changeMode}
  <div
    class="transition-[height] duration-500"
    style="height: {changeMode ? topLayerContainer?.clientHeight : 0}px"
  ></div>
{/if}
<img
  src={pb.files.getURL(pb.authStore.record ?? {}, data.user_photo)}
  alt="person"
  class="w-full rounded-b-3xl object-cover transition-[height]"
  style="height: {!changeMode ? height : 0}px"
  bind:this={image}
/>

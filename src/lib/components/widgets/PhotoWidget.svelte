<script lang="ts">
  import { Splide, SplideSlide } from '@splidejs/svelte-splide';
  import '@splidejs/svelte-splide/css';
  import type { PhotoData } from '$lib/widgetTypes/widgetTypes';
  import { pb } from '$lib';
  let { additionalData }: { additionalData: PhotoData } = $props();

  let urls = $derived.by(() => {
    console.log(additionalData);
    return additionalData.urls.map((url) => pb.buildURL(`/api/files/${url}`));
  });
  $inspect(urls);
</script>

{#key urls}
  <Splide
    arrows={false}
    options={{
      arrows: false,
      rewind: true,
      classes: { page: 'opacity-100! splide__pagination__page ' },
    }}
  >
    {#each urls as src, i}
      <SplideSlide class="flex justify-center items-center">
        <img
          {src}
          class="w-full aspect-video object-contain"
          alt={`image ${i}`}
        />
      </SplideSlide>
    {/each}
  </Splide>
{/key}

{#if urls.length <= 0}
  <i>&lt;нет картинок&gt;</i>
{/if}

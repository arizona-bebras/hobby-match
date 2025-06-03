<script lang="ts">
  import { Splide, SplideSlide } from '@splidejs/svelte-splide';
  import '@splidejs/svelte-splide/css';
  import type { PhotoData, Photos, Widget } from '$lib/widgetTypes/widgetTypes';
  import { pb } from '$lib';
  import type { RecordModel } from 'pocketbase';
  let {
    data,
    additionalData,
  }: { data: RecordModel; additionalData: PhotoData } = $props();
  // console.log(
  //   'ABOBAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABBBBBBBBBBBB',
  //   additionalData,
  // );
  let urls = $derived(
    additionalData.urls.map((url) => pb.buildURL(`/api/files/${url}`)),
  );
  // let test = $derived(
  //   additionalData.urls.map((url) => pb.files.getURL(data, url)),
  // );
  console.log(additionalData.urls);
</script>

<Splide
  arrows={false}
  options={{
    arrows: false,
    classes: { page: 'opacity-100! splide__pagination__page ' },
  }}
  aria-labelledby="My Favorite Images"
>
  {#each urls as src}
    <SplideSlide class="flex justify-center items-center">
      <img {src} class="w-full aspect-video object-contain" alt="Image 1" />
    </SplideSlide>
  {/each}
</Splide>

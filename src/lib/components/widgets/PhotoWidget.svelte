<script lang="ts">
  import { Splide, SplideSlide } from '@splidejs/svelte-splide';
  import '@splidejs/svelte-splide/css';
  let { urls, isTestImage }: { urls: string[], isTestImage: boolean } = $props();

  console.log($state.snapshot(urls))
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
  {#if isTestImage}
    {#each urls as src, i}
      <SplideSlide class="flex justify-center items-center">
        <img
          src = {`${src}`}
          class="w-full aspect-video object-contain"
          alt={`image ${i}`}
        />
      </SplideSlide>
    {/each}
  {:else}
    {#each urls as src, i}
      <SplideSlide class="flex justify-center items-center">
        <img
          src = {`data:image/png;base64,${src}`}
          class="w-full aspect-video object-contain"
          alt={`image ${i}`}
        />
      </SplideSlide>
    {/each}
  {/if}
  </Splide>
{/key}

{#if urls.length <= 0}
  <i>&lt;нет картинок&gt;</i>
{/if}

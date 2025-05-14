<script lang="ts">
  import { Splide, SplideSlide } from '@splidejs/svelte-splide';
  import '@splidejs/svelte-splide/css';
  let { data } = $props();
</script>

<Splide
  arrows={false}
  options={{ arrows: false }}
  aria-labelledby="My Favorite Images"
>
  {#each data.photos as photo}
    <SplideSlide class="flex justify-center items-center">
      <img
        style:max-width="100%"
        style:max-height="100%"
        src={photo.name}
        alt="Image 1"
      />
    </SplideSlide>
  {/each}
</Splide>

<script lang="ts">
  import Siema from 'siema';
  import { onMount } from 'svelte';

  let { data } = $props();
  console.log(data)

  let slider: Siema;
  let select = $state(0);


  onMount(() => {
    slider = new Siema({
      selector: '.siema',
      duration: 200,
      easing: 'ease-in-out',
      perPage: 1,
      startIndex: 0,
      draggable: true,
      multipleDrag: true,
      threshold: 20,
      loop: false,
      rtl: false,
      onChange: () => {
        // Обновляем select при изменении слайда (включая drag)
        select = slider.currentSlide;
      },
    });

    // Инициализируем начальное положение
    select = slider.currentSlide;
  });

  const prev = () => {
    slider.prev();
  };

  const next = () => {
    slider.next();
  };

  const goTo = (index: number) => {
    slider.goTo(index);
  };
</script>

<div class="relative">
  <div class="siema" style="max-width:300px">
    {#each data as d}
      <img src={d} class="siema-img" />
    {/each}
  </div>

  <div class="absolute bottom-4 flex justify-center w-full gap-2">
    {#each data as _, i}
      <input
        type="radio"
        name="slider-radio"
        value={i}
        checked={select === i}
        on:click={() => goTo(i)}
        class="size-2.5"
      />
    {/each}
  </div>
</div>

<style>
  input {
    appearance: none;
    border-radius: 50%;
    background-color: lightgrey;
    transition: 0.2s all linear;
    cursor: pointer;
  }

  .siema-img {
    height: 250px;
    width: 250px;
  }

  input:checked {
    background-color: grey;
    transform: scale(1.2);
  }
</style>

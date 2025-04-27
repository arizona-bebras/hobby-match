<script lang="ts">
  import Siema from 'siema';
  import { onMount } from 'svelte';

  let slider: Siema;
  let select = 0;

  const data = [
    'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQPrTKoiNrYalIuLLSaFMro_QVvrmOD0MTDFQ&s',
    'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQPrTKoiNrYalIuLLSaFMro_QVvrmOD0MTDFQ&s',
  ];

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
  <div class="siema">
    {#each data as d}
      <img src={d} class="w-full" />
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

  input:checked {
    background-color: grey;
    transform: scale(1.2);
  }
</style>

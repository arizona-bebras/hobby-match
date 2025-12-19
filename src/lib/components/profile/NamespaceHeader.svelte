<script lang="ts">
  let {
    scrollDirection,
    scrollYPos,
  }: {
    scrollDirection: 'left' | 'right' | 'top' | 'bottom' | undefined;
    scrollYPos: number;
  } = $props();

  let isHidden = $state(false);
  let startScrollY = $state(0);

  $effect(() => {
    if (scrollDirection === 'top') {
      isHidden = false;
      startScrollY = 0;
    } else if (scrollDirection === 'bottom') {
      if (startScrollY === 0) {
        startScrollY = scrollYPos;
      }
      const scrolledDistance = scrollYPos - startScrollY;
      if (scrolledDistance > 75) {
        isHidden = true;
      }
    }
  });
  $inspect(scrollDirection, scrollYPos);
</script>

<div
  class="left-4 right-4 px-3 rounded-[8px] py-2.5 bg-accent/50 fixed transition-all duration-300 translate-0 z-5 border-2 border-accent-foreground"
  class:translate-y-[-100%]={isHidden}
>
  <p class="mb-1 text-sm text-text-color/50">Анкета из неймспейса</p>
  <div class="flex items-center gap-2">
    <img
      src="https://www.soyuz.ru/public/uploads/files/2/7480281/20220315190534af66e2c5d3.jpg"
      alt="лого неймспейса"
      class="size-4 rounded-full"
    />
    <p>Hello World!</p>
  </div>
</div>

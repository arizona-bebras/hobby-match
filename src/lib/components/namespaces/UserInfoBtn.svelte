<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { ArrowRight, AlarmClock, CircleX } from '@lucide/svelte';
  import { goto } from '$app/navigation';

  let testData = {
    img: 'https://www.soyuz.ru/public/uploads/files/2/7480281/20220315190534af66e2c5d3.jpg',
    username: 'Valerka',
    register_date: '11.09.2006',
  };
</script>

<Sheet.Root>
  <Sheet.Trigger><ArrowRight class="size-5" /></Sheet.Trigger>
  <Sheet.Content side="bottom">
    <div class="flex flex-col justify-center items-center mb-10 mt-8">
      <img
        alt="Картинка неймспейса"
        src={testData.img}
        class="size-24 rounded-full object-cover"
      />
      <div class="text-center mt-2.5">
        <p class="font-medium">{testData.username}</p>
        <p class="text-inactive text-sm">С нами с {testData.register_date}</p>
      </div>
    </div>
    <div class="first:rounded-t-xl last:rounded-b-xl">
      <button
        onclick={() => goto(`/view?type=single&userId=${410782566}`)}
        class="flex justify-between border-2 border-text-color/25 rounded-t-xl border-b-0 px-4 py-3 items-center w-full"
      >
        <p class="">Смотреть анкету</p>
        <ArrowRight class="size-5" />
      </button>
      <div
        class="flex justify-between border-2 border-text-color/25 px-4 py-3 items-center"
      >
        <p class="w-full">Временно отстранить</p>
        <AlarmClock class="size-5" />
      </div>
      <div
        class="flex justify-between border-2 border-text-color/25 rounded-b-xl border-t-0 px-4 py-3 items-center"
      >
        <p class="w-full">Исключить</p>
        <CircleX class="size-5" />
      </div>
    </div>
  </Sheet.Content>
</Sheet.Root>

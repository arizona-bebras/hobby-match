<script lang="ts">
  import { Input } from '$lib/components/ui/input/index.js';
  import { Textarea } from '$lib/components/ui/textarea/index.js';
  import { Mars, Venus } from '@lucide/svelte';
  let gender = $state('');
  let { currentStage = $bindable() } = $props();

  function getAppleEmogi(emoji: string) {
    let code: string[] = [];
    for (const codePoint of emoji) {
      code.push(codePoint.codePointAt(0).toString(16));
    }
    return `https://cdnjs.cloudflare.com/ajax/libs/emoji-datasource-apple/15.1.2/img/apple/64/${code.join('-')}.png`;
  }

  import CalendarIcon from '@lucide/svelte/icons/calendar';
  import {
    type DateValue,
    DateFormatter,
    getLocalTimeZone,
  } from '@internationalized/date';
  import { cn } from '$lib/utils.js';
  import { Button } from '$lib/components/ui/button/index.js';
  import { Calendar } from '$lib/components/ui/calendar/index.js';
  import * as Popover from '$lib/components/ui/popover/index.js';

  const df = new DateFormatter('en-US', {
    dateStyle: 'long',
  });

  let value = $state<DateValue>();
</script>

<div class="h-21">
  <div class="flex w-full">
    <img src={getAppleEmogi('👋')} alt="emoji" class="size-6" />
    <p class="text-xl">Привет! Я Shumi</p>
  </div>
  <p class="font-medium text-2xl">Давай познакомимся!</p>
</div>

<div class="flex flex-col gap-y-4">
  <div>
    <div class="flex w-full mb-2 items-center">
      <img src={getAppleEmogi('😶‍🌫️')} alt="emoji" class="size-4" />
      <p>Как тебя зовут?</p>
    </div>
    <Input type="text" placeholder="Введи своё имя" />
  </div>
  <div>
    <div class="flex w-full mb-2 items-center">
      <img src={getAppleEmogi('👾')} alt="emoji" class="size-4" />
      <p>Какого ты пола?</p>
    </div>
    <div class="flex">
      <button
        class="w-1/2 h-16.75 rounded-l-xl {gender === 'male'
          ? 'bg-[#00A6FF]/50'
          : 'bg-[#00A6FF]/20'}"
        onclick={() => (gender = 'male')}
      >
        <Mars class="text-[#0CB9F8] mx-auto" />
        <p class="text-[#40A7E3]">Мужской</p>
      </button>
      <button
        class="w-1/2 h-16.75 rounded-r-xl {gender === 'female'
          ? 'bg-[#F80CC9]/50'
          : 'bg-[#F80CC9]/15'}"
        onclick={() => (gender = 'female')}
      >
        <Venus class="text-[#F80CC9] mx-auto" />
        <p class="text-[#F80CC9]">Женский</p>
      </button>
    </div>
  </div>
  <div>
    <div class="flex w-full mb-2 items-center">
      <img src={getAppleEmogi('📅')} alt="emoji" class="size-4" />
      <p>Когда ты родился?</p>
    </div>
    <Popover.Root>
      <Popover.Trigger>
        {#snippet child({ props })}
          <Button
            variant="outline"
            class={cn(
              'w-[280px] justify-start text-left font-normal',
              !value && 'text-muted-foreground',
            )}
            {...props}
          >
            <CalendarIcon class="mr-2 size-4" />
            {value
              ? df.format(value.toDate(getLocalTimeZone()))
              : 'Выберите дату'}
          </Button>
        {/snippet}
      </Popover.Trigger>
      <Popover.Content class="w-auto p-0">
        <Calendar bind:value type="single" initialFocus />
      </Popover.Content>
    </Popover.Root>
  </div>
  <div>
    <div class="flex w-full mb-2 items-center">
      <img src={getAppleEmogi('🌍')} alt="emoji" class="size-4" />
      <p>Где ты живёшь?</p>
    </div>
    <Input type="text" placeholder="Начни вводить название города" />
  </div>
  <div>
    <div class="flex w-full mb-2 items-center">
      <img src={getAppleEmogi('💫')} alt="emoji" class="size-4" />
      <p>Расскажи о себе</p>
    </div>
    <Textarea placeholder="Я люблю рисовать и ищу напарника для..." />
  </div>
</div>
<button
  class="w-full h-12 bg-accent rounded-xl mt-12 text-white font-medium"
  onclick={() => (currentStage = 'photo')}>Продолжить</button
>

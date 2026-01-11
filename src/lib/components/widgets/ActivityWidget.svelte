<script lang="ts">
  import Emoji from '$lib/components/ui/emoji/emogi.svelte';
  import { type Activity } from '$lib/widgetTypes/widgetTypes';
  import { getCorrectForm } from '$lib/utils';
  let TestData: Activity = {
    type: 'activity',
    status: {
      isStudent: true,
      isWorker: true,
    },
    education: {
      type: 'university',
      educationStage: 'bachelor',
      course: 2,
    },
    work: {
      company: 'Yep',
      job_title: 'Программист',
      experience: 1,
    },
  };
  let { data }: { data: Activity } = $props();

  let educationStages = {
    bachelor: 'Бакалавриат',
    master: 'Магистратура',
    postgraduate: 'Аспирантура',
  };
  console.log(data);
</script>

<!--<div class="p-2 bg-accent/25 rounded-[12px]">-->
<!--  <div class="mb-4">-->
<!--    <p class="text-xl font-semibold">Ваш статус</p>-->
<!--    <p class="">Выберите один или несколько вариантов вашей занятости</p>-->
<!--  </div>-->
<!--  <p class="font-medium text-gray-400">Чем вы занимаетесь?</p>-->
<!--  <div class="flex">-->
<!--    <button class="flex flex-col flex-1/2 items-center justify-center p-6">-->
<!--      <Emoji symbol="🎓" />-->
<!--      <span>Учусь</span>-->
<!--    </button>-->
<!--    <button class="flex flex-col flex-1/2 items-center justify-center p-6">-->
<!--      <Emoji symbol="💼" />-->
<!--      <span>Работаю</span>-->
<!--    </button>-->
<!--  </div>-->
<!--</div>-->
<div class="p-2.5 bg-accent/25 rounded-[12px]">
  <p class="font-semibold text-lg mb-4">Статус занятости</p>
  {#if data.status.isStudent}
    <div
      class="flex gap-x-4.5 items-center bg-accent/25 p-2 rounded-[12px] border-1 border-accent mb-2 box-border"
    >
      <div class="bg-accent rounded-[12px] flex items-center">
        <p class="scale-175 m-2.5">🎓</p>
      </div>
      <div class="flex flex-col min-w-0">
        <p class="text-xs font-semibold text-accent">ОБУЧЕНИЕ</p>
        {#if data.education.type === 'school'}
          <p class="font-semibold text-[14px] truncate">
            Школа №{data.education.place},
            <span>{data.education.class} класс</span>
          </p>
        {:else}
          <p class="font-semibold text-[14px] truncate">
            {data.education.place}, {educationStages[
              data.education.educationStage
            ]},
            <span>{data.education.course} курс</span>
          </p>
        {/if}
      </div>
    </div>
  {/if}
  {#if data.status.isWorker}
    <div
      class="flex justify-between items-center bg-accent/25 p-2 rounded-[12px] border-1 border-accent"
    >
      <div class="flex gap-x-4.5 min-w-0">
        <div class="bg-accent rounded-[12px] flex items-center">
          <p class="scale-175 m-2.5">💼</p>
        </div>
        <div class="flex flex-col min-w-0">
          <p class="text-xs font-semibold text-accent">РАБОТА</p>
          <p class="font-semibold text-[14px] truncate">
            {data.work.job_title} в {data.work.company}
          </p>
        </div>
      </div>
      {#if data.work.experience}
        <p
          class="text-xs self-start p-1 bg-accent/25 rounded-[8px] border-1 border-accent/75 font-semibold text-[12px] shrink-0"
        >
          Стаж: {data.work.experience}
          {getCorrectForm(data.work.experience, ['год', 'года', 'лет'])}
        </p>
      {/if}
    </div>
  {/if}
</div>

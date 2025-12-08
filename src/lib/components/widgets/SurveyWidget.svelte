<script lang="ts">
  import { Check } from '@lucide/svelte';
  import { Progress } from '@skeletonlabs/skeleton-svelte';
  import { chooseOption } from '$lib/components/widgetConstructors/widgetsConstructor';
  import type { SurveyData, Survey } from '$lib/widgetTypes/widgetTypes';
  import { string } from 'zod';

  let { data, id, survey }: { data: Survey; id: string; survey: SurveyData } =
    $props();

  let selected = $state(survey.myVote);
  let votesCount = $derived(survey.stats.reduce((a, b) => a + b, 0));
  let fakeVote = $derived(survey.myVote === null ? 1 : 0);

  $effect(() => {
    selected = survey.myVote;
  });
  $effect(() => {
    console.log(survey, selected);
    if (selected !== null && survey.myVote === null)
      chooseOption(id, selected!).catch((e) => console.error(e));
  });

  function getCorrectForm(count: number): string {
    const remainder10 = count % 10;
    const remainder100 = count % 100;

    let wordForm: string;

    if (remainder100 >= 11 && remainder100 <= 14) {
      wordForm = 'голосов';
    } else if (remainder10 === 1) {
      wordForm = 'голос';
    } else if (remainder10 >= 2 && remainder10 <= 4) {
      wordForm = 'голоса';
    } else {
      wordForm = 'голосов';
    }

    return wordForm;
  }
</script>

<div class="ToDoBox">
  <div class="flex">
    <p class="pb-2.25 font-bold">{data.question}</p>
    {#if selected !== null}
      <p class="ml-auto">
        {votesCount + fakeVote}
        {getCorrectForm(votesCount + fakeVote)}
      </p>
    {/if}
  </div>
  <!--  <div class="flex gap-2">-->
  <!--    <button class="rounded-full size-5.25 border-2 border-[#D9D9D9]" onclick={isSelected = true}-->
  <!--            class:bg-accent={isSelected} class:border-accent={isSelected}></button>-->
  <!--    <p class:text-accent={isSelected}>Да</p>-->
  <!--  </div>-->
  <!--  <div class="flex gap-2">-->
  <!--    <button class="rounded-full size-5.25 border-2 border-[#D9D9D9]" onclick={isSelected = true}-->
  <!--            class:bg-accent={isSelected} class:border-accent={isSelected}></button>-->
  <!--    <p class:text-accent={isSelected}>Нет</p>-->
  <!--  </div>-->
  {#each data.options as task, i}
    <label
      class="flex items-center space-x-2 relative"
      onclick={(e) => e.stopPropagation()}
    >
      <input
        type="radio"
        name={id}
        value={i}
        bind:group={selected}
        class="appearance-none rounded-full size-5.25 border-2 border-[#D9D9D9]
       cursor-pointer checked:bg-accent checked:border-accent peer opacity-100"
        disabled={selected !== null}
      />
      <Check
        class="absolute size-3.5 left-1 stroke-white invisible peer-checked:visible"
      />
      <p class="peer-checked:text-accent">{task.description}</p>
      {#if selected !== null}
        {@const vote = survey.stats[i] + (selected === i ? fakeVote : 0)}
        <p class="ml-auto">
          {vote}
          {getCorrectForm(vote)}
        </p>
      {/if}
    </label>
    {#if selected !== null}
      <div class="ml-7.25">
        <Progress
          value={survey.stats[i] + (selected === i ? fakeVote : 0)}
          max={votesCount + fakeVote}
          height="h-1"
          meterBg={selected === i ? 'bg-accent' : 'bg-black/40'}
          trackBg="bg-[#D9D9D9]"
        />
      </div>
    {/if}
  {/each}
</div>

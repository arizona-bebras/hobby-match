<script lang="ts">
  import { Check } from '@lucide/svelte';
  import { Progress } from '@skeletonlabs/skeleton-svelte';
  import { chooseOption } from '$lib/components/widgetConstructors/widgetsConstructor'
  let { data, surveyId } = $props();
  let selected = $state('');

  $effect(() => {
    console.log(selected)
    if (selected != '') 
      chooseOption(surveyId, selected)
  })
</script>

<div class="ToDoBox">
  <div class="flex">
    <p class="pb-2.25 font-bold">{data.question}</p>
    {#if selected != ''}
      <p class="ml-auto">{data.summuryVotes} проголосовали</p>
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
  {#each data.options as task}
    {console.log(task)}
    <label class="flex items-center space-x-2">
      <input
        type="radio"
        name="radio-direct"
        bind:group={selected}
        value={task.description}
        class="appearance-none rounded-full size-5.25 border-2 border-[#D9D9D9]
       cursor-pointer checked:bg-accent checked:border-accent peer opacity-100"
        disabled={selected != ''}
      />
      <Check
        class="absolute size-3.5 left-5.5 stroke-white invisible peer-checked:visible"
      />
      <p class="peer-checked:text-accent">{task.description}</p>
      {#if selected != ''}
        <p class="ml-auto">{task.votes} голосов</p>
      {/if}
    </label>
    {#if selected != ''}
      <div class="ml-7.25">
        <Progress
          value={task.votes}
          max={data.summuryVotes}
          height="h-1"
          meterBg={selected === task.description ? 'bg-accent' : 'bg-black/40'}
          trackBg="bg-[#D9D9D9]"
        />
      </div>
    {/if}
  {/each}
</div>

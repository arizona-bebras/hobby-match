<script lang="ts">
  import { Check } from '@lucide/svelte';
  import { Progress } from '@skeletonlabs/skeleton-svelte';
  let { tasks, votes } = $props();
  let selected = $state('');
</script>

<div class="ToDoBox">
  <div class="flex">
    <p class="pb-2.25 font-bold">Да или нет</p>
    {#if selected != ''}
      <p class="ml-auto">{votes} проголосовали</p>
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
  {#each tasks as task}
    {console.log(task)}
    <label class="flex items-center space-x-2">
      <input
        type="radio"
        name="radio-direct"
        bind:group={selected}
        value={task[0]}
        class="appearance-none rounded-full size-5.25 border-2 border-[#D9D9D9]
       cursor-pointer checked:bg-accent checked:border-accent peer opacity-100"
        disabled={selected != ''}
      />
      <Check
        class="absolute size-3.5 left-9.5 stroke-white invisible peer-checked:visible"
      />
      <p class="peer-checked:text-accent">{task[0]}</p>
      {#if selected != ''}
        <p class="ml-auto">{task[1]} голосов</p>
      {/if}
    </label>
    {#if selected != ''}
      <div class="ml-7.25">
        <Progress
          value={task[1]}
          max={votes}
          height="h-1"
          meterBg={selected === task[0] ? 'bg-accent' : 'bg-black/40'}
          trackBg="bg-[#D9D9D9]"
        />
      </div>
    {/if}
  {/each}
</div>

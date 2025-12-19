<script lang="ts">
  let {
    stages = $bindable(),
  }: {
    stages: { title: string; isCurrentStage: boolean; isComplete: boolean }[];
  } = $props();

  function setCurrentStage(newTitle) {
    for (const stage of stages) {
      stage.isCurrentStage = stage.title === newTitle;
    }
  }
</script>

<div class="flex">
  {#each stages as stage (stage)}
    <button
      onclick={() => {
        if (stage.isComplete) {
          setCurrentStage(stage.title);
        }
      }}
      class="border-t-2 w-1/3 py-2.5 {stage.isCurrentStage
        ? 'text-accent font-bold'
        : stage.isComplete
          ? 'text-accent/50'
          : ''} ">{stage.title}</button
    >
  {/each}
</div>
<!--<div class="flex">-->
<!--    {#each stages as stage}-->
<!--      <button-->
<!--        onclick={() => {-->
<!--          if (completedStages.includes(stage[0])) currentStage = stage[0];-->
<!--        }}-->
<!--        class="border-t-2 w-1/3 p-2.5 {completedStages.includes(stage[0])-->
<!--          ? 'text-accent/50'-->
<!--          : currentStage === stage[0]-->
<!--            ? 'text-accent'-->
<!--            : 'text-[#A7A7A7]'}"-->
<!--        class:font-bold={currentStage === stage[0]}>{stage[1]}</button-->
<!--      >-->
<!--    {/each}-->
<!--</div>-->

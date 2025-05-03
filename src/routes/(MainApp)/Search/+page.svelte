<script lang="ts">
  import { Progress } from '@skeletonlabs/skeleton-svelte';
  import { Check } from '@lucide/svelte';
  let tasks = [1, 2, 3, 4];
  let selected = $state('');
</script>

<p>Search page</p>
<div class="group">
  {#each tasks as task}
    <label class="flex items-center space-x-2">
      <input
        type="radio"
        name="radio-direct"
        bind:group={selected}
        value={task}
        class="appearance-none rounded-full size-5.25 border-2 border-[#D9D9D9]
         cursor-pointer checked:bg-accent checked:border-accent peer"
      />
      <p
        class="peer-checked:text-accent"
        class:bg-red-200={selected != task && selected != ''}
      >
        {task}
      </p>
      {console.log(selected)}
    </label>
  {/each}
  <Progress value={null} />
  <Progress value={50} max={100} height="h-4" meterBg="bg-tertiary-500"
    >25</Progress
  >
  <div
    class="size-5 bg-red-200 rounded-full place-items-center place-content-center flex"
  >
    <Check class="size-3.5"></Check>
  </div>
</div>

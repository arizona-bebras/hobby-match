<script lang="ts">
  let { tasks }= $props();
  import { Check } from '@lucide/svelte'; //
</script>


<div class="ToDoBox">
  <p class="pb-2.25 font-bold">Список целей</p>
  <form class="space-y-2.25">
  {#each tasks as task}
    <label class="flex items-center space-x-2">
      <input type="checkbox" class="appearance-none rounded-full size-5.25 border-2 border-[#D9D9D9]
     cursor-pointer checked:bg-accent checked:border-accent peer"/>
      <p class="peer-checked:text-accent peer-checked:line-through peer-checked: decoration-2">{task}</p>
    </label>
  {/each}
  </form>
</div>
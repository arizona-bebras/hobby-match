<script lang="ts">
  import { pb } from '$lib/index'
  let { title, tasks, widgetId } = $props();

  let newTasks = $state.snapshot(tasks);

  function changeTaskState(taskOrder: number) {
    newTasks[taskOrder-1].isCompleted = !newTasks[taskOrder-1].isCompleted;
    changeTaskStateInDB(widgetId);
  }

  async function changeTaskStateInDB(widgetId: string) {
    console.log(newTasks)
    await pb.collection('widgets').update(widgetId, 
    {
      data:
      {
        "title": title,
        "tasks": newTasks
      }
    })
  }
  
  import { Check } from '@lucide/svelte'; //
</script>

<div class="ToDoBox">
  <p class="pb-2.25 font-bold">{title}</p>
  <form class="space-y-2.25">
    {#each tasks as task}
      <label class="flex items-center space-x-2">
        <input
          type="checkbox"
          class="appearance-none rounded-full size-5.25 border-2 border-[#D9D9D9]
     cursor-pointer checked:bg-accent checked:border-accent peer"
          checked={task.isCompleted}
          onchange={() => changeTaskState(task.order)}
        />
        <Check
          class="absolute size-3.5 left-9.5 stroke-white invisible peer-checked:visible"
        />
        <p
          class="peer-checked:text-accent peer-checked:line-through peer-checked: decoration-2"
        >
          {task.description}
        </p>
      </label>
    {/each}
  </form>
</div>

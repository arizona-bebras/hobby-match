<script lang="ts">
  import { updateWidget } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { Check } from '@lucide/svelte';
  import type { Todo } from '$lib/widgetTypes/widgetTypes';
  let {
    data,
    widgetId,
    isViewingMode,
  }: { data: Todo; widgetId: string; isViewingMode: boolean } = $props();
  console.log('Данные Todo:', data);
</script>

<div class="ToDoBox">
  <p class="pb-2.25 font-bold">{data.title}</p>
  <form class="space-y-2.25">
    {#each data.tasks as task}
      <label class="flex items-center space-x-2">
        <input
          type="checkbox"
          class="appearance-none rounded-full size-5.25 border-2 border-[#D9D9D9]
     cursor-pointer checked:bg-accent checked:border-accent peer"
          disabled={isViewingMode}
          checked={task.isCompleted}
          onchange={async (e) => {
            task.isCompleted = e.currentTarget.checked;
            await updateWidget(widgetId, data);
          }}
        />
        <Check
          class="absolute size-3.5 left-5.5 stroke-white invisible peer-checked:visible"
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

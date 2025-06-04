<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { Textarea } from '$lib/components/ui/textarea';
  import {
    createWidget,
    deleteWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input/index.js';
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import { toDoScheme } from '$lib/components/editor/schemes/toDoScheme';
  import { pb } from '$lib';
  import { invalidate } from '$app/navigation';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import { X } from '@lucide/svelte';
  import type { Task } from '$lib/widgetTypes/widgetTypes';
  let {
    nextStage: open = $bindable(),
    numberOfWidgets,
    widgetId,
  }: {
    nextStage: boolean;
    numberOfWidgets: number;
    widgetId?: string;
  } = $props();
  function onOpenChange() {
    setTimeout(() => {
      document.body.style.cssText = '';
      console.log('Компонент уничтожен');
      window.Telegram.WebApp.MainButton.show();
    }, 10);
  }

  const form = superForm(defaults(zod(toDoScheme)), {
    SPA: true,
    validators: zodClient(toDoScheme),
    onSubmit: async ({ formData }) => {
      const formValues = {
        type: 'todo' as const,
        title: formData.get('title'),
        tasks: transformTasksToAPI(formData.getAll('options')),
      };
      if (widgetId != undefined) {
        await updateWidget(widgetId, formValues);
      } else {
        await createWidget(formValues, numberOfWidgets + 1);
      }
    },
  });

  function transformTasksToAPI(tasks: FormDataEntryValue[]): Task[] {
    let result: Task[] = $state([]);
    for (const element of tasks) {
      result.push({
        description: element,
        isCompleted: false,
      });
    }
    return result;
  }

  const { form: formData, enhance, validateForm } = form;

  if (widgetId !== undefined) {
    pb.collection('widgets')
      .getOne(widgetId)
      .then((result) => {
        $formData.title = result.data.title;
        result.data.tasks.forEach(
          (task: { description: string; isCompleted: boolean }) => {
            console.log(task.description);
            $formData.tasks = [...$formData.tasks, task.description];
          },
        );
      });
  }
  let isButtonActive = $state(false);
  $effect(() => {
    validateForm().then((response) => {
      isButtonActive = response.valid;
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });
</script>

<Sheet.Root bind:open {onOpenChange}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance>
        <p class="text-accent-foreground font-medium pb-4.5">
          Виджет "Список задач"
        </p>
        <Form.Field {form} name="title">
          <Form.Control>
            {#snippet children({ props })}
              <Input
                {...props}
                bind:value={$formData.title}
                placeholder="Напишите что-нибудь, предположим, о себе"
                class="mb-6.75"
              />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
        {#each $formData.tasks as element, index}
          <div class="flex justify-center items-center mb-2">
            <Input
              name="options"
              bind:value={$formData.tasks[index]}
              placeholder={`Вариант ${index + 1}`}
              class=""
            />
            <button
              type="button"
              onclick={() =>
                // very tasty))))
                ($formData.tasks = $formData.tasks.toSpliced(index, 1))}
            >
              <X class="size-5 ml-2" />
            </button>
          </div>
        {/each}
        <button
          type="button"
          class="text-accent-foreground font-bold underline underline-offset-3 decoration-2 flex pb-4"
          onclick={() => {
            if ($formData.tasks.length < 8) {
              $formData.tasks[$formData.tasks.length] = '';
            }
          }}
        >
          Добавить вариант
        </button>
        {#if widgetId !== undefined}
          <DeleteButton {widgetId} />
        {/if}

        <SaveButton
          onClick={() => {
            form.submit();
            open = false;
            onOpenChange();
          }}
          {isButtonActive}
        />
        <SuperDebug data={$formData} />
      </form>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>

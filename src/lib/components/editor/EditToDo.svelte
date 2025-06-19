<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input/index.js';
  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import { toDoScheme } from '$lib/components/editor/schemes/toDoScheme';
  import { pb } from '$lib';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import { X } from '@lucide/svelte';
  import type { Todo } from '$lib/widgetTypes/widgetTypes';
  let {
    widgetId,
    onClose,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    onClose: CallableFunction;
  } = $props();

  const form = superForm(defaults(zod(toDoScheme)), {
    SPA: true,
    validators: zodClient(toDoScheme),
    onSubmit: async () => {
      const widget: Todo = {
        type: 'todo',
        title: $formData.title,
        tasks: $formData.tasks,
      };
      if (widgetId) {
        await updateWidget(widgetId, widget);
      } else {
        await createWidget(widget);
      }
      onClose();
    },
  });

  const { form: formData, enhance, validateForm, reset } = form;

  $effect(() => {
    if (widgetId) {
      pb.collection('widgets')
        .getOne(widgetId)
        .then((result) => {
          $formData.title = result.data.title;
          $formData.tasks = result.data.tasks;
        });
    } else {
      reset();
    }
  });

  let isButtonActive = $state(false);
  $effect(() => {
    validateForm().then((response) => {
      isButtonActive = response.valid;
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });
</script>

<Sheet.Root bind:open onOpenChange={(state) => !state && onClose()}>
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
                placeholder="Я хочу..."
              />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
        {#each $formData.tasks as value, index}
          <div class="flex justify-center items-center mb-2">
            <Input
              name="options"
              bind:value={value.description}
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
              $formData.tasks = [
                ...$formData.tasks,
                { description: '', isCompleted: false },
              ];
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
          }}
          {isButtonActive}
        />
        <!--        <SuperDebug data={$formData} />-->
      </form>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>

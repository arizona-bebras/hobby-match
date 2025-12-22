<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import * as Form from '$lib/components/ui/form';
  import { Input } from '$lib/components/ui/input';
  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zod4, zodClient } from 'sveltekit-superforms/adapters';
  import { toDoScheme } from '$lib/components/editor/todo/toDoScheme';
  import { pb } from '$lib';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import { X } from '@lucide/svelte';
  import type { Text, Todo } from '$lib/widgetTypes/widgetTypes';
  let {
    widgetId,
    onClose,
    widgetData,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    widgetData: Todo;
    onClose: CallableFunction;
  } = $props();

  const form = superForm(defaults(zod4(toDoScheme)), {
    SPA: true,
    validators: zod4(toDoScheme),
    onSubmit: async () => {
      isLoading = true;
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
      isLoading = false;
      onClose();
    },
  });

  const { form: formData, enhance, validateForm, reset } = form;

  $effect(() => {
    if (widgetId) {
      $formData.title = widgetData.title;
      $formData.tasks = widgetData.tasks;
    } else {
      reset();
    }
  });

  let isButtonActive = $state(false);
  let isLoading = $state(false);
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
      <form method="POST" use:enhance class="text-text-color">
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
          {isLoading}
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

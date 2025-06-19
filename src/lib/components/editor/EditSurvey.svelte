<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input';
  import { X } from '@lucide/svelte';
  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import { surveyScheme } from '$lib/components/editor/schemes/surveySheme';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { pb } from '$lib';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import type { Survey } from '$lib/widgetTypes/widgetTypes';

  let {
    widgetId,
    onClose,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    onClose: CallableFunction;
  } = $props();
  const form = superForm(defaults(zod(surveyScheme)), {
    SPA: true,
    validators: zodClient(surveyScheme),
    onSubmit: async () => {
      isLoading = true;
      const widget: Survey = {
        type: 'survey',
        question: $formData.question,
        options: $formData.options.map((description) => ({
          description,
        })),
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
  let isButtonActive = $state(false);
  let isLoading = $state(false);
  $effect(() => {
    validateForm().then((response) => {
      isButtonActive = response.valid;
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });

  $inspect($formData.options);
  $effect(() => {
    if (widgetId) {
      pb.collection('widgets')
        .getOne(widgetId)
        .then((result) => {
          console.log(result);
          $formData.question = result.data.question;
          result.data.options.forEach(
            (element: { description: string; votes: number }) => {
              $formData.options = [...$formData.options, element.description];
            },
          );
        });
    } else {
      reset();
    }
  });
</script>

<Sheet.Root bind:open onOpenChange={(state) => !state && onClose()}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance>
        <p class="text-accent-foreground font-medium pb-4.5">Виджет "Опрос"</p>
        <Form.Field {form} name="question">
          <Form.Control>
            {#snippet children({ props })}
              <Input
                placeholder="Напишите какой-нибудь вопрос"
                class="mb-2"
                {...props}
                bind:value={$formData.question}
              />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>

        {#each $formData.options as _, index}
          <div class="flex justify-center items-center mb-2">
            <Input
              name="options"
              bind:value={$formData.options[index]}
              placeholder={`Вариант ${index + 1}`}
              class=""
            />
            <button
              type="button"
              onclick={() =>
                // very tasty))))
                ($formData.options = $formData.options.toSpliced(index, 1))}
            >
              <X class="size-5 ml-2" />
            </button>
          </div>
        {/each}
        <button
          type="button"
          class="text-accent-foreground font-bold underline underline-offset-3 decoration-2 flex pb-4"
          onclick={() => {
            if ($formData.options.length <= 4) {
              $formData.options[$formData.options.length] = '';
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

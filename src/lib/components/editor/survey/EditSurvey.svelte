<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet';
  import * as Form from '$lib/components/ui/form';
  import { Input } from '$lib/components/ui/input';
  import { Info, X } from '@lucide/svelte';
  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zod4, zodClient } from 'sveltekit-superforms/adapters';
  import { surveyScheme } from '$lib/components/editor/survey/surveySheme';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { pb } from '$lib';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import type { Survey, Text } from '$lib/widgetTypes/widgetTypes';
  import { onMount } from 'svelte';

  let {
    widgetId,
    onClose,
    widgetData,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    widgetData: Survey;
    onClose: CallableFunction;
  } = $props();
  const form = superForm(defaults(zod4(surveyScheme)), {
    SPA: true,
    validators: zod4(surveyScheme),
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
  onMount(() => {
    if (widgetId) {
      $formData.question = widgetData.question;
      widgetData.options.forEach((element) => {
        $formData.options = [...$formData.options, element.description];
      });
      // pb.collection('widgets')
      //   .getOne(widgetId)
      //   .then((result) => {
      //     console.log(result);
      //     $formData.question = result.data.question;
      //     result.data.options.forEach(
      //       (element: { description: string; votes: number }) => {
      //         $formData.options = [...$formData.options, element.description];
      //       },
      //     );
      //   });
    } else {
      reset();
    }
  });
</script>

<Sheet.Root bind:open onOpenChange={(state) => !state && onClose()}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance class="text-text-color">
        <p class="text-accent-foreground font-medium pb-4.5">Виджет "Опрос"</p>
        <Form.Field {form} name="question">
          <Form.Control>
            {#snippet children({ props })}
              <Input
                placeholder="Напиши какой-нибудь вопрос"
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
          <div class="flex flex-row pb-2 gap-1 text-gray-400 items-start">
            <Info class="inline-block size-4 mt-1" />
            <span class="align-top">
              При изменении опроса все голоса сбрасываются
            </span>
          </div>
        {/if}
        <SaveButton
          class="mt-2"
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

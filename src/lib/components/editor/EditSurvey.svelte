<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input';
  import { Trash2, X } from '@lucide/svelte';
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import { surveyScheme } from '$lib/components/editor/schemes/surveySheme';
  import {
    createWidget,
    deleteWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { pb } from '$lib';

  let {
    nextStage: open = $bindable(),
    numberOfWidgets,
    widgetId,
  }: {
    nextStage: boolean;
    numberOfWidgets: number;
    widgetId?: string;
  } = $props();
  const form = superForm(defaults(zod(surveyScheme)), {
    SPA: true,
    validators: zodClient(surveyScheme),
    onSubmit: async ({ formData }) => {
      const formValues = {
        question: formData.get('question'),
        options: transformOptionsToAPI(formData.getAll('options')),
        type: 'survey',
        summaryVotes: '0',
      };
      if (widgetId != undefined) {
        await updateWidget(widgetId, formValues);
      } else {
        await createWidget(formValues, numberOfWidgets + 1);
      }
    },
  });

  const { form: formData, enhance, validateForm } = form;
  let isButtonActive = $state(false);
  $effect(() => {
    validateForm().then((response) => {
      isButtonActive = response.valid;
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });

  function onOpenChange() {
    setTimeout(() => {
      document.body.style.cssText = '';
      console.log('Компонент уничтожен');
      window.Telegram.WebApp.MainButton.show();
    }, 10);
  }

  function transformOptionsToAPI(options: FormDataEntryValue[]): object[] {
    let result: object[] = $state([]);
    for (const element of options) {
      result.push({
        description: element,
        votes: 0,
      });
    }
    return result;
  }
  let array: string[] = $state([]);

  $inspect($formData.options);
  if (widgetId !== undefined) {
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
    console.log($formData.options);
  }
</script>

<Sheet.Root bind:open {onOpenChange}>
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

        {#each $formData.options as element, index}
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
          class="text-accent-foreground font-bold underline underline-offset-3 decoration-2 flex"
          onclick={() => {
            if ($formData.options.length <= 4) {
              $formData.options[$formData.options.length] = '';
            }
          }}
        >
          Добавить вариант
        </button>
        {#if widgetId !== undefined}
          <Sheet.Close
            onclick={() => {
              deleteWidget(widgetId.toString());
            }}
            class="ring-offset-background focus:ring-ring data-[state=open]:bg-secondary absolute right-4 top-3.5 rounded-sm opacity-70 transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-offset-2 disabled:pointer-events-none p-2"
          >
            <Trash2 class="size-5 text-destructive" />
            <span class="sr-only">Close</span>
          </Sheet.Close>
        {/if}
        <button
          onclick={() => {
            form.submit();
            open = false;
            onOpenChange();
          }}
          disabled={!isButtonActive}
          class="w-full h-12 {isButtonActive
            ? 'bg-accent'
            : 'bg-inactive'} rounded-xl mt-9">Сохранить</button
        >
        <!--        <SuperDebug data={$formData} />-->
      </form>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>

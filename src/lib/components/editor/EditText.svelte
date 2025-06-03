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
  import { textSchema } from '$lib/components/editor/schemes/textSheme';
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import { Trash2 } from '@lucide/svelte';
  import { pb } from '$lib';
  import { invalidate } from '$app/navigation';
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

  const form = superForm(defaults(zod(textSchema)), {
    SPA: true,
    validators: zodClient(textSchema),
    onSubmit: async ({ formData }) => {
      formData.set('type', 'text');
      const formValues = Object.fromEntries(formData);
      if (widgetId != undefined) {
        await updateWidget(widgetId, formValues);
      } else {
        await createWidget(formValues, numberOfWidgets + 1);
      }
    },
  });

  const { form: formData, enhance, validateForm } = form;

  if (widgetId !== undefined) {
    pb.collection('widgets')
      .getOne(widgetId)
      .then((result) => ($formData.text = result.data.text));
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
        <p class="text-accent-foreground font-medium pb-4.5">Виджет "Текст"</p>
        <Form.Field {form} name="text">
          <Form.Control>
            {#snippet children({ props })}
              <Textarea
                {...props}
                bind:value={$formData.text}
                placeholder="Напишите что-нибудь, предположим, о себе"
                class="mb-6.75"
              />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
        {#if widgetId !== undefined}
          <Sheet.Close
            onclick={async () => {
              await deleteWidget(widgetId.toString());
              await invalidate('user:widgets');
            }}
            class="ring-offset-background focus:ring-ring data-[state=open]:bg-secondary absolute right-4 top-2 rounded-sm opacity-70 transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-offset-2 disabled:pointer-events-none p-2"
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

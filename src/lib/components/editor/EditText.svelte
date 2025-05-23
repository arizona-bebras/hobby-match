<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { Textarea } from '$lib/components/ui/textarea';
  import {
    createWidget,
    deleteWidget,
    updateWidget,
    changeWidgetPosition,
    updateWidgetsOrder,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input/index.js';
  import {
    textSchema,
    type FormSchema,
  } from '$lib/components/editor/schemes/textSheme';
  import SuperDebug, {
    type SuperValidated,
    type Infer,
    superForm,
    defaults,
  } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import type { SubmitFunction } from '@sveltejs/kit';
  import { invalidate } from '$app/navigation';
  let { nextStage: open = $bindable() }: { nextStage: boolean } = $props();

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
      // console.log(formData);
      // const uploadedFiles = formData.getAll('files') as File[];
      // console.log(uploadedFiles);
      const formValues = Object.fromEntries(formData);
      // console.log(formValues);
      await createWidget(formValues, 24);
      // setTimeout(
      //   () =>
      //     invalidate('user:widgets')
      //       .then(() => {
      //         console.log('invalidated');
      //       })
      //       .catch(console.log),
      //   1000,
      // );
    },
  });

  const { form: formData, enhance } = form;
</script>

<Sheet.Root bind:open {onOpenChange}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance>
        <p class="text-accent-foreground font-medium pb-2">Виджет "Текст"</p>
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
        <Form.Button>Submit</Form.Button>
        <button
          onclick={() => {
            open = false;
            onOpenChange();
          }}
          class="w-full h-12 bg-accent rounded-xl">Сохранить</button
        >
        <SuperDebug data={$formData} />/
      </form>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>

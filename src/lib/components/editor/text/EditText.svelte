<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet';
  import { Textarea } from '$lib/components/ui/textarea';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import * as Form from '$lib/components/ui/form';
  import { textSchema } from '$lib/components/editor/text/textSheme';
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zod4, zodClient } from 'sveltekit-superforms/adapters';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import type { Text, Widget } from '$lib/widgetTypes/widgetTypes';
  import client from '$lib/api/client';
  let {
    widgetId,
    onClose,
    widgetData,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    widgetData: Text;
    onClose: CallableFunction;
  } = $props();

  let isLoading = $state(false);
  const form = superForm(defaults(zod4(textSchema)), {
    SPA: true,
    validators: zod4(textSchema),
    onSubmit: async () => {
      isLoading = true;
      const widget: Text = {
        type: 'text',
        text: $formData.text,
      };
      if (widgetId != undefined) {
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
    if (widgetData) {
      $formData.text = widgetData.text;
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
      <form method="POST" use:enhance class="text-text-color">
        <p class="text-accent-foreground font-medium pb-4.5">Виджет "Текст"</p>
        <Form.Field {form} name="text">
          <Form.Control>
            {#snippet children({ props })}
              <Textarea
                {...props}
                bind:value={$formData.text}
                placeholder="Напишите что-нибудь, предположим, о себе"
              />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
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
        {#if import.meta.env.DEV}
          <SuperDebug data={$formData} />
        {/if}
      </form>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>

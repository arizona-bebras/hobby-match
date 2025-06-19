<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { Textarea } from '$lib/components/ui/textarea';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import * as Form from '$lib/components/ui/form/index.js';
  import { textSchema } from '$lib/components/editor/schemes/textSheme';
  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import { pb } from '$lib';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import type { Text } from '$lib/widgetTypes/widgetTypes';
  let {
    widgetId,
    onClose,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    onClose: CallableFunction;
  } = $props();

  let isLoading = $state(false);
  const form = superForm(defaults(zod(textSchema)), {
    SPA: true,
    validators: zodClient(textSchema),
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
    if (widgetId) {
      pb.collection('widgets')
        .getOne(widgetId)
        .then((result) => ($formData.text = result.data.text));
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
        <!--        <SuperDebug data={$formData} />-->
      </form>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>

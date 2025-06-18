<script lang="ts">
  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import { audioScheme } from '$lib/components/editor/schemes/audioScheme';
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { pb } from '$lib';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import type { Audio } from '$lib/widgetTypes/widgetTypes';
  let {
    widgetId,
    onClose,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    onClose: CallableFunction;
  } = $props();
  const form = superForm(defaults(zod(audioScheme)), {
    SPA: true,
    validators: zodClient(audioScheme),
    onSubmit: async () => {
      let url = $formData.link;
      const widget: Audio = {
        type: 'audio',
        link: url,
      };
      if (widgetId != undefined) {
        await updateWidget(widgetId, widget);
      } else {
        await createWidget(widget);
      }
      onClose();
    },
  });
  const { form: formData, enhance, validateForm, reset } = form;
  let isButtonActive = $state(false);
  $effect(() => {
    validateForm().then((response) => {
      isButtonActive = response.valid;
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });
  $effect(() => {
    if (widgetId !== undefined) {
      pb.collection('widgets')
        .getOne(widgetId)
        .then((result) => ($formData.link = result.data.link));
    } else {
      reset();
    }
  });
</script>

<Sheet.Root bind:open onOpenChange={(state) => !state && onClose()}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance>
        <p class="text-accent-foreground font-medium pb-4.5">Виджет "Аудио"</p>
        <p class="pb-2 text-text-color">Введите ссылку на SoundCloud</p>
        <Form.Field {form} name="link">
          <Form.Control>
            {#snippet children({ props })}
              <Input
                placeholder="https://soundcloud.com/..."
                {...props}
                bind:value={$formData.link}
              />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
        {#if widgetId !== undefined}
          <DeleteButton {widgetId} />
        {/if}
        <SaveButton
          onClick={() => {
            form.submit();
          }}
          {isButtonActive}
        />
      </form>
    </Sheet.Header>
    <!--    <SuperDebug data={$formData} />-->
  </Sheet.Content>
</Sheet.Root>

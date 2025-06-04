<script lang="ts">
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import { audioScheme } from '$lib/components/editor/schemes/audioScheme'
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input';
  import {
    createWidget,
    deleteWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { Trash2 } from '@lucide/svelte';
  import { pb } from '$lib';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  let {
    nextStage: open = $bindable(),
    numberOfWidgets,
    widgetId,
  }: {
    nextStage: boolean;
    numberOfWidgets: number;
    widgetId?: string;
  } = $props();
  const form = superForm(defaults(zod(audioScheme)), {
    SPA: true,
    validators: zodClient(audioScheme),
    onSubmit: async () => {
      let url = $formData.link;
      const formValues = {
        type: 'audio' as const,
        link: url,
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
  if (widgetId !== undefined) {
    pb.collection('widgets')
      .getOne(widgetId)
      .then((result) => ($formData.link = `https://t.me/${result.data.link}`));
  }
  function onOpenChange() {
    setTimeout(() => {
      document.body.style.cssText = '';
      console.log('Компонент уничтожен');
      window.Telegram.WebApp.MainButton.show();
    }, 10);
  }
</script>

<Sheet.Root bind:open {onOpenChange}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance>
        <p class="text-accent-foreground font-medium pb-4.5">
          Виджет "Аудио"
        </p>
        <p class="pb-2">Введите ссылку на SoundCloud .</p>
        <Form.Field {form} name="link">
          <Form.Control>
            {#snippet children({ props })}
              <Input
                placeholder="https://t.me/..."
                class="mb-9"
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
            open = false;
            onOpenChange();
          }}
          {isButtonActive}
        />
      </form>
    </Sheet.Header>
    <!--    <SuperDebug data={$formData} />-->
  </Sheet.Content>
</Sheet.Root>

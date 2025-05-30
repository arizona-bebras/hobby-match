<script lang="ts">
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import { postScheme } from '$lib/components/editor/schemes/postSheme';
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
  let {
    nextStage: open = $bindable(),
    numberOfWidgets,
    widgetId,
  }: {
    nextStage: boolean;
    numberOfWidgets: number;
    widgetId?: string;
  } = $props();
  const form = superForm(defaults(zod(postScheme)), {
    SPA: true,
    validators: zodClient(postScheme),
    onSubmit: async ({ formData }) => {
      let url = formData.get('link') as string;
      const formValues = {
        type: 'post' as const,
        link: url.match(/(?<=https:\/\/t\.me\/).*/)![0],
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
          Виджет "Телеграм Пост"
        </p>
        <p class="pb-2">Введите ссылку на пост из публичного канала.</p>
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
          type="button"
          onclick={() => {
            open = false;
            onOpenChange();
            form.submit();
          }}
          disabled={!isButtonActive}
          class="w-full h-12 {isButtonActive
            ? 'bg-accent'
            : 'bg-inactive'} rounded-xl">Сохранить</button
        >
      </form>
    </Sheet.Header>
    <SuperDebug data={$formData} />
  </Sheet.Content>
</Sheet.Root>

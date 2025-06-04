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
  import { progressScheme } from '$lib/components/editor/schemes/progressScheme';
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import { Trash2 } from '@lucide/svelte';
  import { pb } from '$lib';
  import { invalidate } from '$app/navigation';
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
  function onOpenChange() {
    setTimeout(() => {
      document.body.style.cssText = '';
      console.log('Компонент уничтожен');
      window.Telegram.WebApp.MainButton.show();
    }, 10);
  }

  const form = superForm(defaults(zod(progressScheme)), {
    SPA: true,
    validators: zodClient(progressScheme),
    onSubmit: async ({ formData }) => {
      formData.set('type', 'progress_bar');
      const formValues = Object.fromEntries(formData);
      if (widgetId != undefined) {
        await updateWidget(widgetId, formValues);
      } else {
        await createWidget(formValues, numberOfWidgets + 1);
      }
    },
  });

  const { form: formData, enhance, validateForm, allErrors } = form;

  if (widgetId !== undefined) {
    pb.collection('widgets')
      .getOne(widgetId)
      .then((result) => {
        $formData.description = result.data.description;
        $formData.currentProgress = parseInt(result.data.currentProgress);
        $formData.maxProgress = parseInt(result.data.maxProgress);
      });
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
        <p class="text-accent-foreground font-medium pb-4.5">
          Виджет "Прогресс"
        </p>
        <p class="pb-2">Задача:</p>
        <Form.Field {form} name="description">
          <Form.Control>
            {#snippet children({ props })}
              <Input
                {...props}
                bind:value={$formData.description}
                class="mb-1.75"
              />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
        <div class="flex flex-row justify-between w-full">
          <div class="w-[45%]">
            <p>Текущий:</p>
            <Form.Field {form} name="currentProgress">
              <Form.Control>
                {#snippet children({ props })}
                  <Input
                    {...props}
                    bind:value={$formData.currentProgress}
                    type="number"
                    placeholder="Напишите что-нибудь, предположим, о себе"
                    class=""
                  />
                {/snippet}
              </Form.Control>
              <Form.FieldErrors />
            </Form.Field>
          </div>
          <div class="w-[45%]">
            <p class="">Цель:</p>
            <Form.Field {form} name="maxProgress">
              <Form.Control>
                {#snippet children({ props })}
                  <Input
                    {...props}
                    bind:value={$formData.maxProgress}
                    type="number"
                    placeholder="Напишите что-нибудь, предположим, о себе"
                    class=""
                  />
                {/snippet}
              </Form.Control>
              <Form.FieldErrors />
            </Form.Field>
          </div>
        </div>

        {#if $allErrors.length}
          <ul>
            {#each $allErrors as error}
              <li>
                <p class="text-destructive">{error.messages.join('. ')}</p>
              </li>
            {/each}
          </ul>
        {/if}

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
        <SuperDebug data={$formData} />
      </form>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>

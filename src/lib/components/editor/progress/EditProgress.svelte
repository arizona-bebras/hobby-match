<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import * as Form from '$lib/components/ui/form';
  import { Input } from '$lib/components/ui/input';
  import { progressScheme } from '$lib/components/editor/progress/progressScheme';
  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zod4, zodClient } from 'sveltekit-superforms/adapters';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import type { ProgressBar, Video } from '$lib/widgetTypes/widgetTypes';
  let {
    widgetId,
    onClose,
    widgetData,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    widgetData: ProgressBar;
    onClose: CallableFunction;
  } = $props();

  const form = superForm(defaults(zod4(progressScheme)), {
    SPA: true,
    validators: zod4(progressScheme),
    onSubmit: async () => {
      isLoading = true;
      const widget: ProgressBar = {
        type: 'progress_bar',
        currentProgress: $formData.currentProgress,
        description: $formData.description,
        maxProgress: $formData.maxProgress,
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

  const {
    form: formData,
    enhance,
    validateForm,
    allErrors,
    reset,
    errors,
  } = form;

  let isLoading = $state(false);
  $effect(() => {
    if (widgetId) {
      $formData.description = widgetData.description;
      $formData.currentProgress = widgetData.currentProgress;
      $formData.maxProgress = widgetData.maxProgress;
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
        <p class="text-accent-foreground font-medium pb-4.5">
          Виджет "Прогресс"
        </p>
        <p class="pb-2 text-text-color">Задача:</p>
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
            <p class="text-text-color">Текущий:</p>
            <Form.Field {form} name="currentProgress">
              <Form.Control>
                {#snippet children({ props })}
                  <Input
                    {...props}
                    bind:value={$formData.currentProgress}
                    type="number"
                    class=""
                  />
                {/snippet}
              </Form.Control>
              <Form.FieldErrors />
            </Form.Field>
          </div>
          <div class="w-[45%]">
            <p class="text-text-color">Цель:</p>
            <Form.Field {form} name="maxProgress">
              <Form.Control>
                {#snippet children({ props })}
                  <Input
                    {...props}
                    bind:value={$formData.maxProgress}
                    type="number"
                    class=""
                  />
                {/snippet}
              </Form.Control>
              <Form.FieldErrors />
            </Form.Field>
          </div>
        </div>

        {#if $errors._errors !== undefined}
          <p class="text-destructive">{$errors?._errors}</p>
        {/if}

        <!--{#if $allErrors.length}-->
        <!--  <ul>-->
        <!--    {#each $allErrors as error}-->
        <!--      <li>-->
        <!--        <p class="text-destructive">{error.messages.join('. ')}</p>-->
        <!--      </li>-->
        <!--    {/each}-->
        <!--  </ul>-->
        <!--{/if}-->

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

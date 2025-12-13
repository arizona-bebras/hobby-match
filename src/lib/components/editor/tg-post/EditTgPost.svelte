<script lang="ts">
  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zod4, zodClient } from 'sveltekit-superforms/adapters';
  import { postScheme } from '$lib/components/editor/tg-post/postSheme';
  import * as Sheet from '$lib/components/ui/sheet';
  import * as Form from '$lib/components/ui/form';
  import { Input } from '$lib/components/ui/input';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { pb } from '$lib';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  let {
    widgetId,
    onClose,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    onClose: CallableFunction;
  } = $props();
  const form = superForm(defaults(zod4(postScheme)), {
    SPA: true,
    validators: zod4(postScheme),
    onSubmit: async () => {
      isLoading = true;
      const widget = {
        type: 'post' as const,
        link: $formData.link.match(/(?<=https:\/\/t\.me\/).*/)![0],
      };
      if (widgetId) {
        await updateWidget(widgetId, widget);
      } else {
        await createWidget(widget);
      }
      isLoading = false;
      onClose();
    },
  });
  const { form: formData, enhance, validateForm, reset } = form;
  let isButtonActive = $state(false);
  let isLoading = $state(false);
  $effect(() => {
    validateForm().then((response) => {
      isButtonActive = response.valid;
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });
  $effect(() => {
    if (widgetId) {
      pb.collection('widgets')
        .getOne(widgetId)
        .then(
          (result) => ($formData.link = `https://t.me/${result.data.link}`),
        );
    } else {
      reset();
    }
  });
</script>

<Sheet.Root bind:open onOpenChange={(state) => !state && onClose()}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance class="text-text-color">
        <p class="text-accent-foreground font-medium pb-4.5">
          Виджет "Телеграм Пост"
        </p>
        <p class="pb-2 text-text-color">
          Введите ссылку на пост из публичного канала.
        </p>
        <Form.Field {form} name="link">
          <Form.Control>
            {#snippet children({ props })}
              <Input
                placeholder="https://t.me/..."
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
          {isLoading}
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

<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { Input } from '$lib/components/ui/input';
  let {
    numberOfWidgets,
    widgetId,
    onClose,
  }: {
    numberOfWidgets: number;
    widgetId?: string;
    onClose: CallableFunction;
  } = $props();
  import * as Form from '$lib/components/ui/form/index.js';

  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import { videoSchema } from '$lib/components/editor/schemes/videoSheme';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { pb } from '$lib';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';

  const form = superForm(defaults(zod(videoSchema)), {
    SPA: true,
    validators: zodClient(videoSchema),
    onSubmit: async ({ formData }) => {
      const formValues = {
        type: 'video' as const,
        platform: getPlatformType(formData.get('link') as string)!,
        link: formData.get('link') as string,
      };
      if (widgetId !== undefined) {
        await updateWidget(widgetId, formValues);
      } else {
        await createWidget(formValues, numberOfWidgets + 1);
      }
      onClose();
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
      .then((result) => ($formData.link = result.data.link));
  }

  function getPlatformType(url: string) {
    const domain = url.match(/https?:\/\/([^/]+)/)![1].toLowerCase();
    if (domain.includes('youtube')) {
      return 'YouTube' as const;
    } else if (domain.includes('rutube')) {
      return 'Rutube' as const;
    } else if (domain.includes('tiktok')) {
      return 'TikTok' as const;
    }
  }
</script>

<Sheet.Root open={true}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance>
        <p class="text-accent-foreground font-medium pb-4.5">Виджет "Видео"</p>
        <p class="pb-2">Введите ссылку на видео</p>
        <Form.Field {form} name="link">
          <Form.Control>
            {#snippet children({ props })}
              <Input
                placeholder="https://youtube.com/watch?v=..."
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

<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet';
  import { Input } from '$lib/components/ui/input';
  import * as Form from '$lib/components/ui/form';

  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zod4, zodClient } from 'sveltekit-superforms/adapters';
  import { videoSchema } from '$lib/components/editor/video/videoSheme';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import { Info } from '@lucide/svelte';
  import type { Video } from '$lib/widgetTypes/widgetTypes';
  let {
    widgetId,
    onClose,
    widgetData,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    widgetData: Video;
    onClose: CallableFunction;
  } = $props();

  const form = superForm(defaults(zod4(videoSchema)), {
    SPA: true,
    validators: zodClient(videoSchema),
    onSubmit: async () => {
      isLoading = true;
      const widget = {
        type: 'video' as const,
        platform: getPlatformType($formData.link)!,
        link: $formData.link,
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
      $formData.link = widgetData.link;
    } else {
      reset();
    }
  });

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

<Sheet.Root bind:open onOpenChange={(state) => !state && onClose()}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance class="text-text-color">
        <p class="text-accent-foreground font-medium pb-4.5">Виджет "Видео"</p>
        <p class="pb- text-text-color">Введите ссылку на видео</p>

        <div class="flex flex-row pb-2 gap-1 text-gray-400 items-center">
          <Info class="inline-block size-4" />
          <p class=" ">Платформы: YouTube, RuTube, TikTok</p>
        </div>
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

<script lang="ts">
  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zod4, zodClient } from 'sveltekit-superforms/adapters';
  import { audioScheme } from '$lib/components/editor/audio/audioScheme';
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
  const form = superForm(defaults(zod4(audioScheme)), {
    SPA: true,
    validators: zod4(audioScheme),
    onSubmit: async () => {
      isLoading = true;
      let url = $formData.link;
      const widget: Audio = {
        type: 'audio',
        platform: getPlatformType(url)!,
        link: url,
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
    if (widgetId !== undefined) {
      pb.collection('widgets')
        .getOne(widgetId)
        .then((result) => ($formData.link = result.data.link));
    } else {
      reset();
    }
  });

  function getPlatformType(url: string) {
    const domain = url.match(/https?:\/\/([^/]+)/)![1].toLowerCase();
    if (domain.includes('soundcloud')) {
      return 'SoundCloud' as const;
    } else if (domain.includes('music.yandex')) {
      return 'Yandex' as const;
    } else if (domain.includes('open.spotify')) {
      return 'Spotify' as const;
    }
  }
</script>

<Sheet.Root bind:open onOpenChange={(state) => !state && onClose()}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance class="text-text-color">
        <p class="text-accent-foreground font-medium pb-4.5">Виджет "Аудио"</p>
        <p class="pb-2 text-text-color">
          Введите ссылку на SoundCloud или Яндекс Музыку
        </p>
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

<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { Input } from '$lib/components/ui/input';
  let {
    nextStage: open = $bindable(),
    numberOfWidgets,
    widgetId,
  }: {
    nextStage: boolean;
    numberOfWidgets: number;
    widgetId?: string;
  } = $props();
  import * as Form from '$lib/components/ui/form/index.js';

  // АААААА), он не уничтожаеться ))))))))
  // onDestroy(() => {
  //
  //   addedWidget = '';
  //
  // });
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import {
    videoSchema,
    getYouTubeVideoId,
  } from '$lib/components/editor/schemes/videoSheme';
  import {
    createWidget,
    deleteWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { Trash2 } from '@lucide/svelte';
  import { pb } from '$lib';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';

  const form = superForm(defaults(zod(videoSchema)), {
    SPA: true,
    validators: zodClient(videoSchema),
    onSubmit: async ({ formData }) => {
      const formValues = {
        type: 'video' as const,
        platform: getPlatformType(formData.get('link')!),
        link: formData.get('link'),
      };
      if (widgetId !== undefined) {
        await updateWidget(widgetId, formValues);
      } else {
        await createWidget(formValues, numberOfWidgets + 1);
      }
    },
  });
  let platformType = $state('');
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

  function getPlatformType(url: FormDataEntryValue): string {
    const domain = url.match(/https?:\/\/([^/]+)/)[1].toLowerCase();
    if (domain.includes('youtube')) {
      return 'YouTube';
    } else if (domain.includes('rutube')) {
      return 'Rutube';
    } else if (domain.includes('tiktok')) {
      return 'TikTok';
    } else {
      return 'Undefined';
    }
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
            platformType = getPlatformType($formData.link);
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

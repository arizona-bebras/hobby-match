<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { Input } from '$lib/components/ui/input';
  let { nextStage: open = $bindable(), numberOfWidgets } = $props();
  import * as Form from '$lib/components/ui/form/index.js';

  // АААААА), он не уничтожаеться ))))))))
  // onDestroy(() => {
  //
  //   addedWidget = '';
  //
  // });
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import { videoSchema } from '$lib/components/editor/schemes/videoSheme';
  import { createWidget } from '$lib/components/widgetConstructors/widgetsConstructor';

  const form = superForm(defaults(zod(videoSchema)), {
    SPA: true,
    validators: zodClient(videoSchema),
    onSubmit: async ({ formData }) => {
      formData.set('type', 'video');
      formData.set('platform', platformType);
      const formValues = Object.fromEntries(formData);
      await createWidget(formValues, numberOfWidgets + 1);
    },
  });
  let platformType = $state('');
  const { form: formData, enhance } = form;

  function getPlatformType(url: string): string {
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
        <p class="text-accent-foreground font-medium">Виджет "Видео"</p>
        <p>Введите ссылку на видео</p>
        <Form.Field {form} name="link">
          <Form.Control>
            {#snippet children({ props })}
              <Input
                placeholder="https://youtube.com/watch?v=..."
                class="mb-9"
                {...props}
                bind:value={$formData.link}
              />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
        <button
          onclick={() => {
            // наверное это самую малость не правильно)
            platformType = getPlatformType($formData.link);
            form.submit();
            open = false;
            onOpenChange();
          }}
          class="w-full h-12 bg-accent rounded-xl">Сохранить</button
        >
      </form>
    </Sheet.Header>
    <SuperDebug data={$formData} />
  </Sheet.Content>
</Sheet.Root>

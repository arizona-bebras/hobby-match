<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input';
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod } from 'sveltekit-superforms/adapters';
  import { socialScheme } from '$lib/components/editor/schemes/socialScheme';
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
  let isButtonActive = $state(false);

  const form = superForm(defaults(zod(socialScheme)), {
    SPA: true,
    validators: zod(socialScheme),
    onSubmit: async ({ formData }) => {
      let url = formData.get('link') as string;
      const formValues = {
        type: 'social_media' as const,
        platform: getPlatform(url),
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

  function getUsername(link: string): string {
    const url = new URL(link);
    if (url.hostname === 'steamcommunity.com') {
      return url.pathname.split('/')[2]!;
    } else {
      return url.pathname.slice(1);
    }
  }

  function getPlatform(link: string): string {
    const domain = new URL(link).hostname.toLowerCase();

    if (domain.includes('youtube.com')) return 'YouTube';
    if (domain.includes('twitch.tv')) return 'Twitch';
    if (domain.includes('vk.com')) return 'VK';
    if (domain.includes('steamcommunity.com')) return 'Steam';
    if (domain.includes('x.com')) return 'X';
    if (domain.includes('t.me')) return 'Telegram';
    return 'Unknown';
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
        <p class="text-accent-foreground font-medium pb-2">
          Виджет "Социальная сеть"
        </p>
        <p class="pb-2">Вы можете ввести ссылку на канал или личный аккаунт</p>
        <Form.Field {form} name="link">
          <Form.Control>
            {#snippet children({ props })}
              <Input
                placeholder="https://..."
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

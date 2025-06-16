<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input';
  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod } from 'sveltekit-superforms/adapters';
  import { socialScheme } from '$lib/components/editor/schemes/socialScheme';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { pb } from '$lib';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import type { SocialMediaLink } from '$lib/widgetTypes/widgetTypes';
  let {
    widgetId,
    onClose,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    onClose: CallableFunction;
  } = $props();
  let isButtonActive = $state(false);

  const form = superForm(defaults(zod(socialScheme)), {
    SPA: true,
    validators: zod(socialScheme),
    onSubmit: async () => {
      const widget: SocialMediaLink = {
        type: 'social_media' as const,
        platform: getPlatform($formData.link)!,
        link: $formData.link,
      };
      if (widgetId) {
        await updateWidget(widgetId, widget);
      } else {
        await createWidget(widget);
      }
      onClose();
    },
  });
  const { form: formData, enhance, validateForm, reset } = form;

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
        .then((result) => ($formData.link = result.data.link));
    } else {
      reset();
    }
  });

  function getPlatform(link: string) {
    const domain = new URL(link).hostname.toLowerCase();

    if (domain.includes('youtube.com')) return 'YouTube';
    if (domain.includes('twitch.tv')) return 'Twitch';
    if (domain.includes('vk.com')) return 'VK';
    if (domain.includes('steamcommunity.com')) return 'Steam';
    if (domain.includes('x.com')) return 'X';
    if (domain.includes('t.me')) return 'Telegram';
  }
</script>

<Sheet.Root bind:open onOpenChange={(state) => !state && onClose()}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance>
        <p class="text-accent-foreground font-medium pb-2">
          Виджет "Социальная сеть"
        </p>
        <p class="pb-1">Вы можете ввести ссылку на канал или личный аккаунт</p>
        <p class="pb-2 text-gray-400">
          Платформы: YouTube, Twitch, VK, Steam, Telegram
        </p>
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
          }}
          {isButtonActive}
        />
      </form>
    </Sheet.Header>
    <!--    <SuperDebug data={$formData} />-->
  </Sheet.Content>
</Sheet.Root>

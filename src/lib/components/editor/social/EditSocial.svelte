<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet';
  import * as Form from '$lib/components/ui/form';
  import { Input } from '$lib/components/ui/input';
  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zod4 } from 'sveltekit-superforms/adapters';
  import { socialScheme } from '$lib/components/editor/social/socialScheme';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { pb } from '$lib';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import type {
    SocialMediaData,
    SocialMediaLink,
    Text,
  } from '$lib/widgetTypes/widgetTypes';
  import { Info } from '@lucide/svelte';
  import { onMount } from 'svelte';
  let {
    widgetId,
    onClose,
    widgetData,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    widgetData: SocialMediaLink;
    onClose: CallableFunction;
  } = $props();
  let isButtonActive = $state(false);
  let isLoading = $state(false);

  const form = superForm(defaults(zod4(socialScheme)), {
    SPA: true,
    validators: zod4(socialScheme),
    onSubmit: async () => {
      isLoading = true;
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
      isLoading = false;
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

  onMount(() => {
    if (widgetId) {
      $formData.link = widgetData.link;
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
      <form method="POST" use:enhance class="text-text-color">
        <p class="text-accent-foreground font-medium pb-2">
          Виджет "Социальная сеть"
        </p>
        <p class="pb-1 text-text-color">
          Вы можете ввести ссылку на канал или личный аккаунт
        </p>
        <div class="flex flex-row pb-2 gap-1 text-gray-400 items-start">
          <Info class="inline-block size-4 mt-1" />
          <span class="align-top">
            Платформы: YouTube, Twitch, VK, Steam, Telegram
          </span>
        </div>
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

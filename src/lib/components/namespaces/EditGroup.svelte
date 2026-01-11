<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import * as Form from '$lib/components/ui/form/index.js';
  import Emoji from '$lib/components/ui/emoji/emogi.svelte';
  import { Input } from '$lib/components/ui/input';
  import { Textarea } from '$lib/components/ui/textarea';
  import SuperDebug, { defaults, superForm } from 'sveltekit-superforms';
  import { zod4 } from 'sveltekit-superforms/adapters';
  import { createSchema } from '$lib/components/namespaces/schema';
  import client from '$lib/api/client';
  import { activateTgBtn, disableTgBtn } from '$lib/utils';
  import { onMount } from 'svelte';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';

  type NamespaceData = {
    id: string;
    title: string;
    description: string;
    members_count: number;
    admin: number;
  };

  let {
    open = $bindable(),
    namespaceData,
    onSave,
  }: {
    open: boolean;
    namespaceData: NamespaceData;
    onSave: () => void;
  } = $props();

  window.Telegram.WebApp.MainButton.text = 'Сохранить';
  useTelegramButton(() => {
    form.submit();
    open = false;
    onSave();
  });

  let fileButton = $state<HTMLInputElement>();
  const form = superForm(defaults(zod4(createSchema)), {
    SPA: true,
    validators: zod4(createSchema),
    resetForm: false,
    onUpdate({ form }) {
      if (form.valid) {
        // TODO: Call an external API with form.data, await the result and update form
      }
    },
    async onChange() {
      const title = await form.validate('title', { update: false } as never);
      const photo = await form.validate('photo', { update: false } as never);
      const description = await form.validate('description', {
        update: false,
      } as never);
      if (!title && !photo && !description) {
        activateTgBtn();
      } else {
        disableTgBtn();
      }
    },
    async onSubmit() {
      let requestBody = {
        id: namespaceData.id,
        title: $formData.title,
        description: $formData.description,
      };
      if ($formData.photo) {
        requestBody.photo = $formData.photo;
      }
      await client.PATCH('/api/admin/{namespace_id}', {
        params: {
          path: {
            namespace_id: namespaceData.id,
          },
        },
        body: requestBody,
        bodySerializer(body) {
          const fd = new FormData();
          for (const name in body) {
            //@ts-expect-error i love dockerimage
            fd.append(name, body[name]);
          }
          return fd;
        },
      });
    },
  });
  const { form: formData, enhance } = form;

  onMount(() => {
    $formData.title = namespaceData.title;
    $formData.description = namespaceData.description;
  });

  $effect(() => {
    if (open) {
      window.Telegram.WebApp.MainButton.show();
    } else {
      window.Telegram.WebApp.MainButton.hide();
    }
  });
</script>

<Sheet.Root bind:open>
  <Sheet.Content side="bottom" class="max-h-[calc(100vh-65px)]">
    <Sheet.Header>
      <Sheet.Title>Редактирование неймспейса</Sheet.Title>
    </Sheet.Header>
    <form
      method="POST"
      enctype="multipart/form-data"
      use:enhance
      class="space-y-6 max-h-[calc(100vh-110px)] overflow-y-auto"
    >
      <Form.Field {form} name="title">
        <Form.Control>
          {#snippet children({ props })}
            <Form.Label class="text-[16px] font-medium flex items-center">
              Название
            </Form.Label>
            <Input {...props} bind:value={$formData.title} />
          {/snippet}
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>
      <Form.Field {form} name="photo">
        <Form.Control>
          {#snippet children({ props })}
            <p class="text-[16px] font-medium flex items-center">Фотография</p>
            <input
              type="file"
              {...props}
              bind:this={fileButton}
              oninput={() => {
                $formData.photo = fileButton.files[0];
              }}
              class="hidden"
            />
            <button
              type="button"
              onclick={() => {
                fileButton.click();
              }}
              class="w-full bg-accent py-1.5 rounded-lg">Выбрать файл...</button
            >
          {/snippet}
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>
      <Form.Field {form} name="description">
        <Form.Control>
          {#snippet children({ props })}
            <Form.Label class="text-[16px] font-medium flex items-center">
              Описание
            </Form.Label>
            <Textarea {...props} bind:value={$formData.description} />
          {/snippet}
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>
    </form>
    {#if import.meta.env.DEV}
      <SuperDebug data={$formData} />
    {/if}
  </Sheet.Content>
</Sheet.Root>

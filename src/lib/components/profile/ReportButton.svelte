<script lang="ts">
  import { pb } from '$lib';
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import * as Select from '$lib/components/ui/select/index.js';
  import * as Form from '$lib/components/ui/form/index.js';
  import { toast } from 'svelte-sonner';
  import { TriangleAlert } from '@lucide/svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import { defaults, superForm } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  import { reasons, reportSchema } from '$lib/components/profile/reportSchema';
  import { Input } from '$lib/components/ui/input';
  import Emoji from '$lib/components/ui/emoji/emogi.svelte';

  let { offender }: { offender: string } = $props();
  let open = $state(false);
  $effect(() => {
    if (open) {
      window.Telegram.WebApp.MainButton.hide();
    } else {
      window.Telegram.WebApp.MainButton.show();
    }
  });

  const form = superForm(defaults(zod(reportSchema)), {
    SPA: true,
    validators: zodClient(reportSchema),
    onSubmit: async () => {
      try {
        await pb.collection('reports').create({
          reporter: pb.authStore.record?.id,
          offender,
          reason: $formData.reason,
          info: $formData.info,
        });
        toast.success('Жалоба отправлена. Спасибо!');
      } catch {
        toast.error('Жалоба уже отправлена');
      }
      open = false;
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
</script>

<button class="text-destructive p-2" onclick={() => (open = true)}>
  <TriangleAlert />
</button>

<Sheet.Root bind:open onOpenChange={(state) => (open = state)}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" enctype="multipart/form-data" use:enhance>
        <p class="text-accent-foreground font-medium pb-4">Отправить жалобу</p>

        <Form.Field {form} name="reason">
          <Form.Control>
            {#snippet children({ props })}
              <Form.Label>
                <Emoji class="mr-1" symbol="⚡" />Причина жалобы
              </Form.Label>
              <Select.Root
                type="single"
                bind:value={$formData.reason}
                name="reason"
              >
                <Select.Trigger
                  {...props}
                  class="w-full text-ellipsis overflow-clip whitespace-nowrap mb-2 mt-1"
                  >{$formData.reason}</Select.Trigger
                >
                <Select.Content>
                  {#each reasons as reason}
                    <Select.Item value={reason}>{reason}</Select.Item>
                  {/each}
                </Select.Content>
              </Select.Root>
            {/snippet}
          </Form.Control>
        </Form.Field>
        <Form.Field {form} name="info">
          <Form.Control>
            {#snippet children({ props })}
              <Form.Label>
                <Emoji class="mr-1" symbol="🤔" />Дополнительная информация
              </Form.Label>
              <Input
                {...props}
                bind:value={$formData.info}
                placeholder="Что не так?"
                class="mb-2 mt-1"
              />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
        <SaveButton
          text="Отправить"
          {isButtonActive}
          class="mt-1"
          onClick={() => {
            form.submit();
          }}
        />
      </form>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>

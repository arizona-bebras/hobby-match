<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { Textarea } from '$lib/components/ui/textarea';
  import { createWidget } from '$lib/components/widgetConstructors/widgetsConstructor';
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input/index.js';
  import { textSchema } from '$lib/components/editor/schemes/textSheme';
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zodClient } from 'sveltekit-superforms/adapters';
  let {
    nextStage: open = $bindable(),
    numberOfWidgets,
  }: { nextStage: boolean; numberOfWidgets: number } = $props();

  function onOpenChange() {
    setTimeout(() => {
      document.body.style.cssText = '';
      console.log('Компонент уничтожен');
      window.Telegram.WebApp.MainButton.show();
    }, 10);
  }

  const form = superForm(defaults(zod(textSchema)), {
    SPA: true,
    validators: zodClient(textSchema),
    onSubmit: async ({ formData }) => {
      formData.set('type', 'text');
      const formValues = Object.fromEntries(formData);
      await createWidget(formValues, numberOfWidgets + 1);
    },
  });

  const { form: formData, enhance } = form;
</script>

<Sheet.Root bind:open {onOpenChange}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" use:enhance>
        <p class="text-accent-foreground font-medium pb-2">Виджет "Текст"</p>
        <Form.Field {form} name="text">
          <Form.Control>
            {#snippet children({ props })}
              <Textarea
                {...props}
                bind:value={$formData.text}
                placeholder="Напишите что-нибудь, предположим, о себе"
                class="mb-6.75"
              />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
        <button
          onclick={() => {
            form.submit();
            open = false;
            onOpenChange();
          }}
          class="w-full h-12 bg-accent rounded-xl">Сохранить</button
        >
        <SuperDebug data={$formData} />/
      </form>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>

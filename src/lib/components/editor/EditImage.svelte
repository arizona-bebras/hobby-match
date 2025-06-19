<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { X } from '@lucide/svelte';
  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod } from 'sveltekit-superforms/adapters';
  import { imageScheme } from '$lib/components/editor/schemes/imageSheme';
  import {
    createWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { pb } from '$lib';
  import { invalidate } from '$app/navigation';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';

  let {
    widgetId,
    onClose,
    open = $bindable(false),
  }: {
    open: boolean;
    widgetId?: string;
    onClose: CallableFunction;
  } = $props();

  let photoFiles: FileList | undefined = $state();
  $effect(() => {
    console.log(photoInput, photoFiles);
  });

  let photoInput: HTMLInputElement;

  const form = superForm(defaults(zod(imageScheme)), {
    SPA: true,
    validators: zod(imageScheme),
    onSubmit: async () => {
      const widget = {
        type: 'photo' as const,
      };
      if (!widgetId) {
        await createWidget(widget, $formData.files);
      } else {
        await updateWidget(widgetId, widget, $formData.files);
      }
      onClose();
    },
  });

  let imageUrls: string[] = $state([]);
  $effect(() => {
    photoFiles = undefined;
    if (widgetId) {
      pb.collection('widgets')
        .getOne(widgetId!)
        .then((record) => (imageUrls = record.additionalData.urls));
    } else {
      reset();
    }
  });

  const { form: formData, enhance, validateForm, reset } = form;
  let isButtonActive = $state(false);
  $effect(() => {
    validateForm().then((response) => {
      console.log(response);
      isButtonActive = response.valid;
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });
</script>

<Sheet.Root bind:open onOpenChange={(state) => !state && onClose()}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <form method="POST" enctype="multipart/form-data" use:enhance>
        <p class="text-accent-foreground font-medium pb-4">
          Виджет "Изображения"
        </p>
        <div class="overflow-auto max-h-80">
          {#if widgetId !== undefined}
            {#each imageUrls as url, i}
              <div class="w-full h-auto bg-accent/45 rounded-2xl relative mb-4">
                <img
                  src={pb.buildURL(`/api/files/${url}`)}
                  class="p-4"
                  alt="loadedImage"
                />
                <button
                  class="absolute right-4 top-4"
                  onclick={async () => {
                    await pb.collection('widgets').update(widgetId, {
                      'files-': [url.split('/').pop()],
                    });
                    imageUrls.splice(i, 1);
                    await invalidate('user:widgets');
                    isButtonActive = true;
                  }}
                  type="button"
                >
                  <X class=" text-accent" />
                </button>
              </div>
            {/each}
          {/if}
          {#each $formData.files as element, index}
            <div class="w-full h-auto bg-accent/45 rounded-2xl relative mb-4">
              <img
                src={URL.createObjectURL(element)}
                class="p-4"
                alt="loadedImage"
              />
              <button
                class="absolute right-4 top-4"
                onclick={() =>
                  ($formData.files = $formData.files.toSpliced(index, 1))}
                type="button"
              >
                <X class="text-accent" />
              </button>
            </div>
          {/each}
        </div>
        <button
          onclick={() => {
            photoInput.click();
          }}
          type="button"
          class="w-full h-12 bg-accent rounded-xl mb-2 mt-4"
          >Выбрать изображение</button
        >
        <input
          name="files"
          type="file"
          class="hidden"
          bind:files={photoFiles}
          bind:this={photoInput}
          accept="image/png, image/jpeg"
          oninput={() => {
            console.log(photoInput.files);
            Array.from(photoInput.files ?? []).forEach(
              (element: File) =>
                ($formData.files = [...$formData.files, element]),
            );
            console.log($formData.files);
          }}
        />
        {#if widgetId !== undefined}
          <DeleteButton {widgetId} />
        {/if}
        <button
          type="button"
          onclick={() => {
            form.submit();
          }}
          disabled={!isButtonActive}
          class="w-full h-12 {isButtonActive
            ? 'bg-accent'
            : 'bg-inactive'} rounded-xl mt-2">Сохранить</button
        >
      </form>
    </Sheet.Header>
    <!--      <SuperDebug data={$formData} />-->
  </Sheet.Content>
</Sheet.Root>

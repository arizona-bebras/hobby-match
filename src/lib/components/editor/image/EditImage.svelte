<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet';
  import { X } from '@lucide/svelte';
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zod4 } from 'sveltekit-superforms/adapters';
  import { imageScheme } from '$lib/components/editor/image/imageSheme';
  import {
    createWidget,
    deletePhotoFromWidget,
    updateWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { db } from '$lib';
  import { invalidate } from '$app/navigation';
  import DeleteButton from '$lib/components/editor/DeleteButton.svelte';
  import SaveButton from '$lib/components/editor/SaveButton.svelte';
  import type { PhotoData } from '$lib/widgetTypes/widgetTypes';

  let {
    widgetId,
    onClose,
    open = $bindable(false),
    widgetData,
    files,
  }: {
    open: boolean;
    widgetId?: string;
    widgetData: PhotoData;
    onClose: CallableFunction;
    files: string[];
  } = $props();

  let photoFiles: FileList | undefined = $state();
  $effect(() => {
    console.log(photoInput, photoFiles);
  });

  let photoInput: HTMLInputElement;
  let isLoading = $state(false);
  const form = superForm(defaults(zod4(imageScheme)), {
    SPA: true,
    validators: zod4(imageScheme),
    onSubmit: async () => {
      isLoading = true;
      const widget = {
        type: 'photo' as const,
      };
      if (!widgetId) {
        await createWidget(widget, $formData.files);
      } else {
        await updateWidget(widgetId, widget, $formData.files);
      }
      isLoading = false;
      onClose();
    },
  });

  let imageUrls: string[] = $state([]);
  $effect(() => {
    photoFiles = undefined;
    if (widgetId) {
      // TODO: pocketbase was here
      // pb.collection('widgets')
      //   .getOne(widgetId!)
      //   .then((record) => (imageUrls = record.additionalData.urls));
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
      <form
        method="POST"
        enctype="multipart/form-data"
        use:enhance
        class="text-text-color"
      >
        <p class="text-accent-foreground font-medium pb-4">
          Виджет "Изображения"
        </p>
        <div class="overflow-auto max-h-80">
          {#if widgetId !== undefined}
            {#each files as url, i}
              <div class="w-full h-auto bg-accent/45 rounded-2xl relative mb-4">
                <img
                  src={`data:image/png;base64,${url}`}
                  class="p-4"
                  alt="loadedImage"
                />
                <button
                  class="absolute right-4 top-4"
                  onclick={async () => {
                    await deletePhotoFromWidget(widgetId, i);
                    files.splice(i, 1);
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
          disabled={isLoading}
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
        <SaveButton
          class="mt-1"
          {isLoading}
          onClick={() => {
            form.submit();
          }}
          {isButtonActive}
        />
      </form>
    </Sheet.Header>
    {#if import.meta.env.DEV}
      <SuperDebug data={$formData} />
    {/if}
  </Sheet.Content>
</Sheet.Root>

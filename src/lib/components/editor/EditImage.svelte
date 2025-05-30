<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { Plus, Trash2 } from '@lucide/svelte';
  import { Input } from '$lib/components/ui/input';
  import SuperDebug, {
    superForm,
    defaults,
    filesProxy,
  } from 'sveltekit-superforms';
  import { zod } from 'sveltekit-superforms/adapters';
  import { imageScheme } from '$lib/components/editor/schemes/imageSheme';
  import {
    createWidget,
    deleteWidget,
  } from '$lib/components/widgetConstructors/widgetsConstructor';
  import type { Photo } from '$lib/widgetTypes/widgetTypes';
  import { pb } from '$lib';

  let {
    nextStage: open = $bindable(),
    numberOfWidgets,
    widgetId,
  }: {
    nextStage: boolean;
    numberOfWidgets: number;
    widgetId?: string;
  } = $props();
  let photoInput: HTMLInputElement;
  let images: File[] = $state([]);
  const form = superForm(defaults(zod(imageScheme)), {
    SPA: true,
    //dataType: 'json',
    validators: zod(imageScheme),
    onSubmit: async ({ formData }) => {
      const formValues = {
        type: 'photo' as const,
      };
      await createWidget(formValues, numberOfWidgets + 1, images);
    },
  });

  function onOpenChange() {
    setTimeout(() => {
      document.body.style.cssText = '';
      console.log('Компонент уничтожен');
      window.Telegram.WebApp.MainButton.show();
    }, 10);
  }

  function transformPhotoToApi(images: File[]): Photo[] {
    let result: object[] = [];
    for (const element of images) {
      result.push({
        name: element.name,
        width: 100,
        height: 100,
        size: element.size,
      });
    }
    return result;
  }
  const { form: formData, enhance } = form;
  const files = filesProxy(form, 'files');
  //$inspect($files.item);
</script>

<div class="">
  <Sheet.Root bind:open {onOpenChange}>
    <Sheet.Content side="bottom">
      <Sheet.Header>
        <form method="POST" enctype="multipart/form-data" use:enhance>
          <p class="text-accent-foreground font-medium pb-2">
            Виджет "Изображения"
          </p>

          <div class="overflow-auto h-75">
            {#each images as element, index}
              {console.log(URL.createObjectURL(element), index)}
              <div class="w-full h-auto bg-accent/45 rounded-2xl relative mb-4">
                <img
                  src={URL.createObjectURL(element)}
                  class="p-4"
                  alt="loadedImage"
                />
                <button onclick={() => images.splice(index, 1)} type="button">
                  <Plus class="rotate-45 absolute right-3 top-3 text-black" />
                </button>
              </div>
            {/each}
          </div>
          <button
            onclick={() => {
              photoInput.click();
            }}
            type="button"
            class="w-full h-12 bg-accent rounded-xl mb-2 mt-4">+</button
          >
          <input
            name="files"
            type="file"
            multiple
            class="hidden"
            bind:this={photoInput}
            bind:files={$files}
            accept="image/png, image/jpeg"
            oninput={() => {
              //ts: я подтверждаю, что он итерируемый
              // images.push([...photoInput.files]);
              let convert = [...photoInput.files];
              convert.forEach((element) => images.push(element));
            }}
          />
          {#if widgetId !== undefined}
            <Sheet.Close
              onclick={() => {
                deleteWidget(widgetId.toString());
              }}
              class="ring-offset-background focus:ring-ring data-[state=open]:bg-secondary absolute right-4 top-3.5 rounded-sm opacity-70 transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-offset-2 disabled:pointer-events-none p-2"
            >
              <Trash2 class="size-5 text-destructive" />
              <span class="sr-only">Close</span>
            </Sheet.Close>
          {/if}
          <button
            type="button"
            onclick={() => {
              form.submit();
              images = [];
              open = false;
              onOpenChange();
            }}
            class="w-full h-10 bg-accent rounded-xl">Сохранить</button
          >
        </form>
      </Sheet.Header>
      <!--      <SuperDebug data={$formData} />-->
    </Sheet.Content>
  </Sheet.Root>
</div>

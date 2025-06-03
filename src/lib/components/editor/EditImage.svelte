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
    updateWidget,
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
      if (widgetId != undefined) {
        console.log('ВИДЖЕТ УСПЕШНО ОБНОВЛЁН', formData.get('files'));
        await updateWidget(widgetId, formValues, formData.get('files'));
      } else {
        console.log('ВИДЖЕТ УСПЕШНО СОЗДАН', images);
        await createWidget(formValues, numberOfWidgets + 1, images);
      }
    },
  });

  let imageUrls: string[] = $state([]);
  if (widgetId !== undefined) {
    pb.collection('widgets')
      .getOne(widgetId!)
      .then((record) =>
        record.files.forEach((element) =>
          imageUrls.push(pb.files.getURL(record, element)),
        ),
      );
    // pb.collection('widgets')
    //   .getOne(widgetId)
    //   .then((result) =>
    //     result.files.forEach((element) =>
    //       console.log(pb.buildURL(`/api/files/${element}`)),
    //     ),
    //   );
    // pb.
    //console.log(pb.buildURL(`/api/files/${result.files}`))
    // let test = getRecords();
    //   getRecords();
  }

  // async function getRecords(titleImages: ):{
  //   const record = await pb.collection('widgets').getOne(widgetId!);
  //   console.log(
  //     pb.files.getURL(record, 'hero_falling_left_sprite_list_ipeux570cb.png'),
  //   );
  // }

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
  const { form: formData, enhance, validateForm } = form;
  const files = filesProxy(form, 'files');
  let isButtonActive = $state(false);
  $effect(() => {
    validateForm().then((response) => {
      isButtonActive = response.valid;
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });
  // $inspect(images);
  //$inspect($files.item);
</script>

<div class="">
  <Sheet.Root bind:open {onOpenChange}>
    <Sheet.Content side="bottom">
      <Sheet.Header>
        <form method="POST" enctype="multipart/form-data" use:enhance>
          <p class="text-accent-foreground font-medium pb-4">
            Виджет "Изображения"
          </p>
          <div class="overflow-auto h-75">
            {#if widgetId !== undefined}
              {#each imageUrls as url}
                <div
                  class="w-full h-auto bg-accent/45 rounded-2xl relative mb-4"
                >
                  <img src={url} class="p-4" alt="loadedImage" />
                  <button onclick={() => images.splice(index, 1)} type="button">
                    <Plus class="rotate-45 absolute right-3 top-3 text-black" />
                  </button>
                </div>
              {/each}
            {/if}
            {#each images as element, index}
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

            <!--{#each images as element, index}-->
            <!--  {console.log(URL.createObjectURL(element), index)}-->
            <!--  <div class="w-full h-auto bg-accent/45 rounded-2xl relative mb-4">-->
            <!--    {#if widgetId !== undefined}-->
            <!--      {console.log(454545454)}-->
            <!--      <img src={imageUrls[0]} class="p-4" alt="loadedImage" />-->
            <!--    {:else}-->
            <!--      {console.log(100)}-->
            <!--      <img-->
            <!--        src={URL.createObjectURL(element)}-->
            <!--        class="p-4"-->
            <!--        alt="loadedImage"-->
            <!--      />-->
            <!--    {/if}-->
            <!--    <button onclick={() => images.splice(index, 1)} type="button">-->
            <!--      <Plus class="rotate-45 absolute right-3 top-3 text-black" />-->
            <!--    </button>-->
            <!--  </div>-->
            <!--{/each}-->
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
              open = false;
              onOpenChange();
              if (widgetId !== undefined && images.length === 0) {
                //pass
              } else {
                form.submit();
              }
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
</div>

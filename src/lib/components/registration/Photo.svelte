<script lang="ts">
  import { pb } from '$lib/index';
  import Emoji from '$lib/components/ui/emoji/emogi.svelte';
  let fileInput: HTMLInputElement;
  let tgImage = window.Telegram.WebApp.initDataUnsafe.user?.photo_url;
  console.log(tgImage);
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input/index.js';
  import {
    photoSchema,
    type FormSchema,
  } from '$lib/components/registration/PhotoFormShema';
  import SuperDebug, {
    type SuperValidated,
    type Infer,
    superForm,
    fileProxy,
  } from 'sveltekit-superforms';
  import { zodClient } from 'sveltekit-superforms/adapters';
  import { informationSchema } from '$lib/components/registration/InformationFormShema';
  import { object } from 'zod';
  import { onDestroy } from 'svelte';
  let { form: photo }: { form: SuperValidated<Infer<FormSchema>> } = $props();

  const form = superForm(photo, {
    validators: zodClient(photoSchema),
    dataType: 'json',
  });

  const { form: formData, enhance, validateForm } = form;

  let fetchData = async () => {
    //await console.log(formData);
    //await pb.collection('users').update(pb.authStore.model?.id, $formData);
    form.submit();
    //window.Telegram.WebApp.MainButton.offClick(fetchData);
  };

  window.Telegram.WebApp.MainButton.onClick(async () => {
    form.submit();
    let f = () => this;
    await pb.collection('users').update(pb.authStore.model?.id, $formData);
    window.Telegram.WebApp.MainButton.offClick(f);
  });
  const file = fileProxy(form, 'user_photo');
  $effect(() => {
    validateForm().then((response) => {
      if (response.valid) {
        window.Telegram.WebApp.MainButton.setParams({
          color: window.Telegram.WebApp.themeParams.button_color,
          is_active: true,
          is_visible: true,
        });
      } else {
        window.Telegram.WebApp.MainButton.setParams({
          color: '#808080',
          is_active: false,
          is_visible: true,
        });
      }
    });
    // eslint-disable-next-line @typescript-eslint/no-unused-expressions
    $formData;
  });

  onDestroy(() => {
    window.Telegram.WebApp.MainButton.hide();
  });
</script>

<form method="POST" action="?/photo" enctype="multipart/form-data" use:enhance>
  <div class="flex flex-col gap-y-2">
    <div class="flex w-full mb-2 items-center">
      <Emoji symbol="📷" class="size-6" />
      <p class="font-medium text-2xl h-10.25">Сфоткаемся?</p>
    </div>
    <button
      type="button"
      class="w-full h-12 bg-accent rounded-xl text-white font-medium"
      onclick={() => {
        fileInput.click();
      }}
    >
      Выбрать файл</button
    >
    <input
      type="file"
      class="hidden"
      bind:this={fileInput}
      bind:files={$file}
      accept="image/png, image/jpeg"
    />
    <p class="self-center">или</p>
    <button
      onclick={() => ($formData.user_photo = tgImage ?? '')}
      class="w-full h-12 bg-accent rounded-xl text-white font-medium"
      >Взять текущую фотографию из Telegram</button
    >
  </div>
</form>
<!--<SuperDebug data={$formData} />-->

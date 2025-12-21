<script lang="ts">
  import { pb } from '$lib';
  import Emoji from '$lib/components/ui/emoji/emogi.svelte';
  import {
    photoSchema,
    type FormSchema,
  } from '$lib/components/registration/photo/PhotoFormShema';
  import SuperDebug, {
    type SuperValidated,
    type Infer,
    superForm,
    fileProxy,
  } from 'sveltekit-superforms';
  import { zodClient } from 'sveltekit-superforms/adapters';
  import { onDestroy } from 'svelte';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte.js';
  import { type Stages, updatePhoto } from '$lib/components/registration/';
  // import { BOT_TOKEN } from '$env/static/private';
  let fileInput: HTMLInputElement;
  let tgImage = window.Telegram.WebApp.initDataUnsafe.user?.photo_url;
  console.log(tgImage);
  let {
    form: photo,
    setCurrentStage,
    markStageComplete,
  }: {
    form: SuperValidated<Infer<FormSchema>>;
    setCurrentStage: (stage: Stages) => void;
    markStageComplete: (stage: Stages) => void;
  } = $props();

  export async function save(): Promise<boolean> {
    if (!hasPhoto || formValid) {
      window.Telegram.WebApp.MainButton.showProgress();
      console.log($formData);
      const res = await updatePhoto($formData).finally(
        window.Telegram.WebApp.MainButton.hideProgress,
      );
      if (res != 200) {
        console.log('failed to update user data');
        return false;
      }
      return true;
    } else {
      return false;
    }
  }

  const form = superForm(photo, {
    validators: zodClient(photoSchema),
    dataType: 'json',
  });

  const { form: formData, enhance, validateForm, errors } = form;

  const file = fileProxy(form, 'user_photo');
  async function handleTelegramButtonClick() {
    await save();
    markStageComplete('Фото');
    setCurrentStage('Тест');
  }

  let hasPhoto: boolean = $derived(!!pb.authStore.record!.user_photo);
  let formValid: boolean = $state(false);
  useTelegramButton(handleTelegramButtonClick);
  $effect(() => {
    validateForm().then((response) => {
      formValid = response.valid;
      if (response.valid || hasPhoto) {
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
  let previewSrc = $derived.by(() => {
    if (!hasPhoto || formValid) {
      if (!$formData.user_photo) return null;
      return typeof $formData.user_photo === 'string'
        ? $formData.user_photo
        : URL.createObjectURL($formData.user_photo);
    } else {
      return pb.files.getURL(
        pb.authStore.record ?? {},
        pb.authStore.record?.user_photo,
      );
    }
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
    {#if $errors.user_photo}
      <p class="text-destructive text-sm font-medium">{$errors.user_photo}</p>
    {/if}
    <input
      type="file"
      class="hidden"
      bind:this={fileInput}
      bind:files={$file}
      accept="image/png, image/jpeg, image/svg+xml, image/gif, image/webp"
    />
    <p class="self-center">или</p>
    <button
      type="button"
      onclick={async () => {
        // const tgPhotoFile = await fetch(`https://api.telegram.org/file/bot${BOT_TOKEN}/${tgImage}`)
        $formData.user_photo = tgImage ?? '';
      }}
      class="w-full h-12 bg-accent rounded-xl text-white font-medium"
      >Взять текущую фотографию из Telegram</button
    >
    {#if previewSrc}
      <img
        src={previewSrc}
        class="rounded-3xl object-cover aspect-square my-2"
        alt="аватарка"
      />
    {/if}
  </div>
</form>
<!--<SuperDebug data={$formData} />-->

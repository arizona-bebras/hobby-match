<script lang="ts">
  import { pb } from '$lib/index';
  import Emoji from '$lib/components/ui/emoji/emogi.svelte';
  let fileInput: HTMLInputElement;
  let tgImage = window.Telegram.WebApp.initDataUnsafe.user?.photo_url;
  console.log(tgImage);
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
  import { onDestroy } from 'svelte';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  let {
    form: photo,
    nextStage,
  }: { form: SuperValidated<Infer<FormSchema>>; nextStage: CallableFunction } =
    $props();

  const form = superForm(photo, {
    validators: zodClient(photoSchema),
    dataType: 'json',
    onSubmit: async () => {
      await pb
        .collection('users')
        .update(pb.authStore.record!.id, $formData)
        .finally(window.Telegram.WebApp.MainButton.hideProgress);
      nextStage();
    },
  });

  const { form: formData, enhance, validateForm } = form;

  const file = fileProxy(form, 'user_photo');
  async function handleTelegramButtonClick() {
    if (!hasPhoto) {
      window.Telegram.WebApp.MainButton.showProgress();
      form.submit();
    } else {
      nextStage();
    }
  }

  let hasPhoto: boolean = $derived(!!pb.authStore.record!.user_photo);
  useTelegramButton(handleTelegramButtonClick);
  $effect(() => {
    validateForm().then((response) => {
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
      accept="image/png, image/jpeg, image/svg+xml, image/gif, image/webp"
    />
    <p class="self-center">или</p>
    <button
      onclick={() => ($formData.user_photo = tgImage ?? '')}
      class="w-full h-12 bg-accent rounded-xl text-white font-medium"
      >Взять текущую фотографию из Telegram</button
    >
  </div>
</form>
<SuperDebug data={$formData} />

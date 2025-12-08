<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  import SuperDebug, { superForm, defaults } from 'sveltekit-superforms';
  import { zod4 } from 'sveltekit-superforms/adapters';
  import { createSchema } from '$lib/components/namespaces/schema';
  import * as Form from '$lib/components/ui/form/index.js';
  import { Input } from '$lib/components/ui/input/index.js';
  import { Textarea } from '$lib/components/ui/textarea/index.js';
  import Emoji from '$lib/components/ui/emoji/emogi.svelte';
  import QRCode from '@castlenine/svelte-qrcode';
  import copy from 'copy-to-clipboard';

  let currentStage = $state(1);
  let fileButton: HTMLInputElement = $state();
  let isSheetOpen = $state(false);

  const form = superForm(defaults(zod4(createSchema)), {
    SPA: true,
    validators: zod4(createSchema),
    onUpdate({ form }) {
      if (form.valid) {
        // TODO: Call an external API with form.data, await the result and update form
      }
    },
    async onChange() {
      if (currentStage === 1) {
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
      } else if (currentStage === 2) {
        const secret_word = await form.validate('secret_word', {
          update: false,
        } as never);
        if (!secret_word) {
          activateTgBtn();
        } else {
          disableTgBtn();
        }
      }
    },
  });

  const { form: formData, enhance } = form;

  useTelegramButton(() => {
    currentStage += 1;
    if (currentStage === 2) {
      disableTgBtn();
    }
    if (currentStage === 3) {
      Telegram.WebApp.MainButton.text = 'Закрыть';
    }
    if (currentStage > 3) {
      currentStage = 1;
      isSheetOpen = false;
      form.submit();
    }
  });

  function onSheetOpenHandle() {
    Telegram.WebApp.MainButton.setParams({
      text: 'Далее',
      is_active: false,
      color: Telegram.WebApp.themeParams.hint_color,
    });
  }
  function activateTgBtn() {
    Telegram.WebApp.MainButton.setParams({
      is_active: true,
      color: Telegram.WebApp.themeParams.button_color,
    });
  }

  function disableTgBtn() {
    Telegram.WebApp.MainButton.setParams({
      is_active: false,
      color: Telegram.WebApp.themeParams.hint_color,
    });
  }
  $inspect(isSheetOpen);
  let link = `https://t.me/share/url?url=${encodeURI('https://core.telegram.org/widgets/share')}&text=${encodeURI('hello world')}
           `;
</script>

<Sheet.Root bind:open={isSheetOpen}>
  <Sheet.Trigger
    class="px-6 py-1.5 rounded-[8px] bg-accent text-[14px]"
    onclick={() => {
      onSheetOpenHandle();
    }}>Создать</Sheet.Trigger
  >
  <Sheet.Content side="bottom" class="max-h-[calc(100vh-65px)]">
    <Sheet.Header>
      <Sheet.Title class="text-center text-[16px] font-medium"
        >{currentStage === 1
          ? 'Новый неймспейс'
          : currentStage === 2
            ? 'Тип неймспейса'
            : $formData.title}</Sheet.Title
      >
    </Sheet.Header>
    <form
      method="POST"
      use:enhance
      class="space-y-6 max-h-[calc(100vh-100px)] overflow-y-auto"
    >
      {#if currentStage === 1}
        <Form.Field {form} name="title">
          <Form.Control>
            {#snippet children({ props })}
              <Form.Label class="text-[16px] font-medium flex items-center">
                <Emoji symbol="🤔" class="size-4 mr-1" />
                Как назовём?
              </Form.Label>
              <Input {...props} bind:value={$formData.title} />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
        <Form.Field {form} name="photo">
          <Form.Control>
            {#snippet children({ props })}
              <Form.Label class="text-[16px] font-medium flex items-center">
                <Emoji symbol="📸" class="size-4 mr-1" />
                Украсим фотографией?
              </Form.Label>
              <p class="text-[14px] mb-3">
                Аватарка поможет пользователям быстрее различать твой неймспейс
                в списке доступных. Ну и это красиво
              </p>
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
                class="w-full bg-accent py-1.5 rounded-lg"
                >Выбрать файл...</button
              >
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
        <Form.Field {form} name="description">
          <Form.Control>
            {#snippet children({ props })}
              <Form.Label class="text-[16px] font-medium flex items-center">
                <Emoji symbol="💬" class="size-4 mr-1" />
                А что будет?
              </Form.Label>
              <p class="text-[14px] mb-3">
                Введи описание будущего неймспейса. Это наверняка поможет
                пользователям найти твое сообщество ^^
              </p>
              <Textarea {...props} bind:value={$formData.description} />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
      {:else if currentStage === 2}
        <Form.Field {form} name="secret_word">
          <Form.Control>
            {#snippet children({ props })}
              <Form.Label class="text-[16px] font-medium flex items-center">
                <Emoji symbol="🔗" class="size-4 mr-1" />
                Секретное слово?
              </Form.Label>
              <p class="text-[14px] mb-3">
                Введи описание будущего неймспейса. Это наверняка поможет
                пользователям найти твое сообщество ^^
              </p>
              <Textarea {...props} bind:value={$formData.secret_word} />
            {/snippet}
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
      {:else if currentStage === 3}
        <div class="flex flex-col justify-center items-center mb-6">
          <QRCode data="Hello World!" />
          <div class="text-center mt-6.5">
            <span class="mb-1"
              ><Emoji symbol="🎉" class="size-4 mr-1" /> Всё готово!</span
            >
            <p class="text-gray-400">«{$formData.title}»</p>
            <p class="text-gray-400">готов принимать гостей!</p>
          </div>
        </div>
        <button
          type="button"
          class="py-2 w-full bg-accent/20 mb-2 rounded-xl text-accent-foreground"
          onclick={() => {
            console.log(Telegram.WebApp.version);
            Telegram.WebApp.shareMessage('dasda');
            copy('Hello World Telegram miniapp');
          }}>Скопировать ссылку</button
        >
        <a href={link}>test</a>
        <button type="button" class="py-2 w-full bg-accent rounded-xl mb-4"
          >Поделиться</button
        >
      {/if}
      <!--        <SuperDebug data={formData} />-->
    </form>
  </Sheet.Content>
</Sheet.Root>
<!--<button class="px-6 py-1.5 rounded-[8px] bg-accent text-[14px]">Создать</button>-->

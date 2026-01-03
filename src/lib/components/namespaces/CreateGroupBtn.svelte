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
  import { toSvg } from 'jdenticon';
  import client from '$lib/api/client';
  import { onMount } from 'svelte';
  import { db } from '$lib';

  let {
    currentStage = 1,
    isSheetOpen = $bindable(),
    data = {},
  }: {
    currentStage: number;
    isSheetOpen: boolean;
    data: { id: string; title: string };
  } = $props();

  let fileButton: HTMLInputElement = $state();
  let uploadedPhoto = $state('');
  let namespaceId = $state('');
  let namespaceTitle = $state('');
  $inspect(uploadedPhoto);
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
      }
    },
    async onSubmit() {
      namespaceTitle = $formData.title;
      const response = await client.POST('/api/namespace', {
        body: {
          title: $formData.title,
          photo: $formData.photo,
          description: $formData.description,
        },
        bodySerializer(body) {
          const fd = new FormData();
          for (const name in body) {
            //@ts-expect-error i love dockerimage
            fd.append(name, body[name]);
          }
          return fd;
        },
      });
      namespaceId = response.data!.namespace_id!;
      console.log('DATA:', response.data, namespaceId);
      // client.POST('/api/vote', {
      //   body: {
      //     aboba: 123,
      //   },
      // });
      console.log('Форма отправлена!', $formData.photo);
    },
  });

  const { form: formData, enhance } = form;

  useTelegramButton(async () => {
    //t.me/@botfather?start=9127099d-971f-4b75-9966-b2c24c2dd79a
    currentStage += 1;
    if (currentStage === 2) {
      Telegram.WebApp.MainButton.text = 'Закрыть';
      const svgString = toSvg(generateRandomHash(), 100);
      console.log(svgString);
      if (!$formData.photo) {
        uploadedPhoto = svgXmlToDataURLRobust(svgString);
        await svgToPngBlob(svgString, 100, 100).then((pngBlop) => {
          console.log(pngBlop);
          $formData.photo = pngBlop as Blob;
        });

        // const blob = new Blob([svgString], { type: 'image/svg+xml' });
        // $formData.photo = new File([blob], 'image.svg', {
        //   type: 'image/svg+xml',
        // });
      }
      form.submit();
    }
    if (currentStage > 2) {
      Telegram.WebApp.MainButton.hide();
      currentStage = 1;
      isSheetOpen = false;
    }
  });

  onMount(async () => {
    let photo = await getBase64Image(
      `${db}/api/files/namespaces/${data.id}/undefined`,
    );
    uploadedPhoto = photo;
  });

  function onSheetOpenHandle() {
    Telegram.WebApp.MainButton.setParams({
      text: 'Далее',
      is_active: false,
      color: Telegram.WebApp.themeParams.hint_color,
      is_visible: true,
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

  let link = $derived.by(() => {
    let basePath = 'https://t.me/newshumibot?start=';
    if (data.id) basePath += encodeURI(data.id);
    else basePath += encodeURI(namespaceId);
    console.log(basePath, 'id:', namespaceId);
    return basePath;
  });

  function svgXmlToDataURLRobust(svgXml) {
    const utf8Bytes = new TextEncoder().encode(svgXml);
    const binaryString = String.fromCharCode.apply(null, utf8Bytes);
    const base64 = btoa(binaryString);
    return `data:image/svg+xml;base64,${base64}`;
  }

  async function svgToPngBlob(svgString: string, width = 500, height = 500) {
    return new Promise((resolve, reject) => {
      // 1. Создаем объект Image и Blob из SVG строки
      const img = new Image();
      const svgBlob = new Blob([svgString], {
        type: 'image/svg+xml;charset=utf-8',
      });
      const url = URL.createObjectURL(svgBlob);

      img.onload = () => {
        // 2. Подготавливаем Canvas нужного размера
        const canvas = document.createElement('canvas');
        canvas.width = width;
        canvas.height = height;
        const ctx = canvas.getContext('2d');

        // 3. Рисуем SVG на холст
        ctx.drawImage(img, 0, 0, width, height);

        // 4. Конвертируем Canvas в Blob (формат image/png)
        canvas.toBlob((blob) => {
          URL.revokeObjectURL(url); // Очищаем память
          if (blob) {
            resolve(blob);
          } else {
            reject(new Error('Ошибка при создании Blob'));
          }
        }, 'image/png');
      };

      img.onerror = () => {
        URL.revokeObjectURL(url);
        reject(new Error('Ошибка загрузки SVG'));
      };

      img.src = url;
    });
  }

  function generateRandomHash(length = 24) {
    let randomString =
      Math.random().toString(16).substring(2) +
      Math.random().toString(16).substring(2);
    while (randomString.length < length) {
      randomString += Math.random().toString(16).substring(2);
    }
    return randomString.substring(0, length);
  }
  async function getBase64Image(url) {
    const response = await fetch(url);
    const blob = await response.blob();
    return new Promise((resolve, reject) => {
      const reader = new FileReader();
      reader.onloadend = () => resolve(reader.result);
      reader.onerror = reject;
      reader.readAsDataURL(blob);
    });
  }
</script>

<Sheet.Root
  bind:open={isSheetOpen}
  onOpenChange={(open) => {
    if (!open) {
      Telegram.WebApp.MainButton.hide();
      if (data.title) currentStage = 2;
      else currentStage = 1;
      isSheetOpen = false;
    }
  }}
>
  <Sheet.Trigger
    class="px-6 py-1.5 rounded-[8px] bg-accent text-[14px]"
    onclick={() => {
      onSheetOpenHandle();
    }}
    >Создать
  </Sheet.Trigger>
  <Sheet.Content side="bottom" class="max-h-[calc(100vh-65px)]">
    <Sheet.Header>
      <Sheet.Title class="text-center text-[16px] font-medium"
        >{currentStage === 1 ? 'Новый неймспейс' : $formData.title}</Sheet.Title
      >
    </Sheet.Header>
    <form
      method="POST"
      enctype="multipart/form-data"
      use:enhance
      class="space-y-6 max-h-[calc(100vh-110px)] overflow-y-auto"
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
              <p class="text-[16px] font-medium flex items-center">
                <Emoji symbol="📸" class="size-4 mr-1" />
                Украсим фотографией?
              </p>
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
                  let reader = new FileReader();
                  reader.readAsDataURL(fileButton.files[0]);
                  reader.onload = (ev) => {
                    uploadedPhoto = ev.target.result;
                  };
                }}
                class="hidden"
              />
              <button
                type="button"
                onclick={() => {
                  fileButton.click();
                }}
                class="w-full bg-accent py-1.5 rounded-lg"
                >Выбрать файл...
              </button>
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
        <!--{:else if currentStage === 2}-->
        <!--  <Form.Field {form} name="secret_word">-->
        <!--    <Form.Control>-->
        <!--      {#snippet children({ props })}-->
        <!--        <Form.Label class="text-[16px] font-medium flex items-center">-->
        <!--          <Emoji symbol="🔗" class="size-4 mr-1" />-->
        <!--          Секретное слово?-->
        <!--        </Form.Label>-->
        <!--        <p class="text-[14px] mb-3">-->
        <!--          Введи описание будущего неймспейса. Это наверняка поможет-->
        <!--          пользователям найти твое сообщество ^^-->
        <!--        </p>-->
        <!--        <Textarea {...props} bind:value={$formData.secret_word} />-->
        <!--      {/snippet}-->
        <!--    </Form.Control>-->
        <!--    <Form.FieldErrors />-->
        <!--  </Form.Field>-->
      {:else if currentStage === 2 && (namespaceId || data.title) && uploadedPhoto}
        <div class="flex flex-col justify-center items-center mb-6">
          <QRCode
            data={link}
            logoInBase64={uploadedPhoto}
            logoSize={15}
            logoPadding={1}
          />
          <div class="text-center mt-6.5">
            <span class="mb-1"
              ><Emoji symbol="🎉" class="size-4 mr-1" /> Всё готово!</span
            >
            <p class="text-gray-400">
              «{data.title ? data.title : namespaceTitle}»
            </p>
            <p class="text-gray-400">готов принимать гостей!</p>
          </div>
        </div>
        <button
          type="button"
          class="py-2 w-full bg-accent/20 mb-2 rounded-xl text-accent-foreground"
          onclick={() => {
            copy(link);
          }}
          >Скопировать ссылку
        </button>
        <a
          class="py-2 w-full bg-accent rounded-xl mb-4 block text-center"
          href={`https://t.me/share/url?url=${encodeURI(link)}`}>Поделиться</a
        >
      {/if}
      <SuperDebug data={formData} />
    </form>
  </Sheet.Content>
</Sheet.Root>
<!--<button class="px-6 py-1.5 rounded-[8px] bg-accent text-[14px]">Создать</button>-->

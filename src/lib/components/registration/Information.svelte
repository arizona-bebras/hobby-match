<script lang="ts">
  import { pb } from '$lib/index';
  import {
    DateFormatter,
    type DateValue,
    getLocalTimeZone,
    today,
  } from '@internationalized/date';
  import SuperDebug, {
    type Infer,
    superForm,
    type SuperValidated,
  } from 'sveltekit-superforms';
  import { zodClient } from 'sveltekit-superforms/adapters';
  import {
    type FormSchema,
    informationSchema,
  } from '$lib/components/registration/InformationFormShema';
  import Emoji from '$lib/components/ui/emoji/emogi.svelte';
  import { Input } from '$lib/components/ui/input';
  import { Mars, Venus, CalendarIcon } from '@lucide/svelte';
  import { Textarea } from '$lib/components/ui/textarea';
  import DataPicker from '$lib/components/ui/dataPicker/DataPicker.svelte';
  import * as Form from '$lib/components/ui/form/index.js';
  import { photoSchema } from '$lib/components/registration/PhotoFormShema';
  import { onDestroy } from 'svelte';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  let gender = $state('');

  let value = $state<DateValue>();
  // let dateOfBirth = $derived(
  //   `${value?.day}-${value?.month.toString().padStart(2, '0')}-${value?.year}`,
  // );
  //
  // $inspect(dateOfBirth);
  let {
    form: information,
    nextStage,
  }: { form: SuperValidated<Infer<FormSchema>>; nextStage: CallableFunction } =
    $props();

  const form = superForm(information, {
    validators: zodClient(informationSchema),
    dataType: 'json',
  });

  const { form: formData, enhance, validateForm } = form;

  async function handleTelegramButtonClick() {
    window.Telegram.WebApp.MainButton.showProgress();
    await pb
      .collection('users')
      .update(pb.authStore.record!.id, $formData)
      .finally(window.Telegram.WebApp.MainButton.hideProgress);
    form.submit();
    nextStage();
  }

  useTelegramButton(handleTelegramButtonClick);
  // $effect(() => {
  //   async function onClick() {
  //     form.submit();
  //     await pb.collection('users').update(pb.authStore.model?.id, $formData);
  //   }
  //   window.Telegram.WebApp.MainButton.onClick(onClick);
  //   return () => {
  //     window.Telegram.WebApp.MainButton.offClick(onClick);
  //   };
  // });
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
  $effect(() => {
    $formData.miniapp_name = pb.authStore.record!.miniapp_name;
    $formData.gender = pb.authStore.record!.gender;
    $formData.birth_date = pb.authStore.record?.birth_date
      ? new Date(pb.authStore.record?.birth_date).toISOString()
      : '';
    $formData.location = pb.authStore.record!.location;
    $formData.user_info = pb.authStore.record!.user_info;
  });
  onDestroy(() => {
    window.Telegram.WebApp.MainButton.hide();
  });
</script>

<form method="POST" action="?/information" use:enhance>
  <div class="h-21">
    <div class="flex w-full">
      <Emoji symbol="👋" class="size-6" />
      <p class="text-xl">Привет! Я Shumi</p>
    </div>
    <p class="font-medium text-2xl">Давай познакомимся!</p>
  </div>
  <div class="flex flex-col gap-y-4">
    <div>
      <div class="flex w-full mb-2 items-center">
        <Emoji symbol="😶‍🌫️" />
        <p>Как тебя зовут?</p>
      </div>
      <Form.Field {form} name="miniapp_name">
        <Form.Control>
          {#snippet children({ props })}
            <Input
              {...props}
              placeholder="Введи своё имя"
              bind:value={$formData.miniapp_name}
              name="username"
            />
          {/snippet}
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>
    </div>
    <div>
      <div class="flex w-full mb-2 items-center">
        <Emoji symbol="👾" />
        <p>Какого ты пола?</p>
      </div>
      <div class="flex">
        <button
          class="w-1/2 h-16.75 rounded-l-xl {$formData.gender === 'male'
            ? 'bg-[#00A6FF]/50'
            : 'bg-[#00A6FF]/20'}"
          onclick={(e) => {
            $formData.gender = 'male';
            e.preventDefault();
          }}
        >
          <div class="size-6 bg-background mx-auto rounded-full flex">
            <Mars class="text-[#0CB9F8] m-auto size-3" />
          </div>
          <p class="text-[#40A7E3]">Мужской</p>
        </button>
        <button
          class="w-1/2 h-16.75 rounded-r-xl {$formData.gender === 'female'
            ? 'bg-[#F80CC9]/50'
            : 'bg-[#F80CC9]/15'}"
          onclick={(e) => {
            $formData.gender = 'female';
            e.preventDefault();
          }}
        >
          <div class="size-6 bg-background mx-auto rounded-full flex">
            <Venus class="text-[#F80CC9] m-auto size-3" />
          </div>
          <p class="text-[#F80CC9]">Женский</p>
        </button>
      </div>
    </div>
    <div>
      <div class="flex w-full mb-2 items-center">
        <Emoji symbol="📅" />
        <p>Когда ты родился?</p>
      </div>
      <DataPicker
        bind:value={$formData.birth_date}
        maxValue={today(getLocalTimeZone())}
      />
    </div>
    <div>
      <div class="flex w-full mb-2 items-center">
        <Emoji symbol="🌍" />
        <p>Где ты живёшь?</p>
      </div>
      <Form.Field {form} name="location">
        <Form.Control>
          {#snippet children({ props })}
            <Input
              {...props}
              placeholder="Начни вводить название города"
              bind:value={$formData.location}
            />
          {/snippet}
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>
    </div>
    <div>
      <div class="flex w-full mb-2 items-center">
        <Emoji symbol="💫" />
        <p>Расскажи о себе</p>
      </div>
      <Form.Field {form} name="user_info">
        <Form.Control>
          {#snippet children({ props })}
            <Textarea
              {...props}
              placeholder="Я люблю рисовать и ищу напарника для..."
              bind:value={$formData.user_info}
            />
          {/snippet}
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>
    </div>
  </div>
  <!--  <SuperDebug data={$formData} />-->
</form>

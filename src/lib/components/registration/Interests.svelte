<script lang="ts">
  import { Input } from '$lib/components/ui/input/index.js';
  import { Plus } from '@lucide/svelte';
  import Emoji from '$lib/components/ui/emoji/emoji.svelte';
  import SuperDebug, {
    type Infer,
    superForm,
    type SuperValidated,
  } from 'sveltekit-superforms';
  import {
    type FormSchema,
    interestsScheme,
  } from '$lib/components/registration/InterestsFormShema';
  import { zodClient } from 'sveltekit-superforms/adapters';
  import { onDestroy } from 'svelte';
  import { goto } from '$app/navigation';
  let listOfInterests = $state([
    'Видеоигры',
    'Готовка',
    'Музыка',
    'Фотография',
    'Кино',
    'Авто',
    'Рисование',
    'Дизайн',
    'Кодинг',
  ]);
  let selectedInterests: string[] = $state([]);
  let userInterest: string = $state('');
  let { form: interests }: { form: SuperValidated<Infer<FormSchema>> } =
    $props();

  const form = superForm(interests, {
    validators: zodClient(interestsScheme),
    dataType: 'json',
  });

  const { form: formData, enhance, validateForm } = form;
  window.Telegram.WebApp.MainButton.onClick(() => {
    form.submit();
  });
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

x

<form method="POST" action="?/interests" use:enhance>
  <div class="flex flex-col gap-y-2">
    <div class="flex w-full mb-2 items-center">
      <p class="font-medium text-2xl h-15.25">
        <Emoji symbol="💗" size={6} />Давай узнаем друг-друга поближе?
      </p>
    </div>
    <p>Выберите наиболее интересующие вас темы (максимум 5)</p>
    <div class="flex flex-col gap-y-4">
      <Input
        type="text"
        placeholder="Начните вводить"
        bind:value={userInterest}
        onkeydown={(e) => {
          if (e.key === 'Enter') {
            listOfInterests.push(userInterest);
            console.log(listOfInterests);
          }
        }}
      />
      <div class="flex flex-row gap-2 w-full flex-wrap font-medium">
        {#each listOfInterests as element}
          <button
            onclick={(e) => {
              if (selectedInterests.includes(element)) {
                selectedInterests.splice(selectedInterests.indexOf(element), 1);
              } else {
                selectedInterests.push(element);
              }
              $formData.interests = selectedInterests;
              e.preventDefault();
            }}
            class="{selectedInterests.includes(element)
              ? 'bg-accent'
              : 'bg-accent/25'} rounded-3xl flex flex-row items-center justify-center px-3 py-2 gap-1.5"
          >
            <Plus
              class="{selectedInterests.includes(element)
                ? 'text-white rotate-45'
                : 'text-accent-foreground'} w-4 h-5 stroke-3"
            />
            <p
              class="text-accent-foreground {selectedInterests.includes(element)
                ? 'text-white'
                : 'text-accent-foreground'}"
            >
              {element}
            </p>
          </button>
        {/each}
      </div>
    </div>
  </div>
  <SuperDebug data={$formData} />
</form>

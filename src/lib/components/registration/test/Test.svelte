<script lang="ts">
  import { pb } from '$lib';
  import { db } from '$lib';
  import { Slider } from '$lib/components/ui/slider/index.js';
  import Emoji from '$lib/components/ui/emoji/emogi.svelte';
  import { testSchema } from '$lib/components/registration/test/TestFormShema';
  import { superForm, defaults } from 'sveltekit-superforms';
  import { zod, zod4 } from 'sveltekit-superforms/adapters';
  import { onDestroy } from 'svelte';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte.js';
  import { updateData } from '$lib/components/registration/';
  let {
    setCurrentStage,
    markStageComplete,
  }: {
    setCurrentStage: (stage: string) => void;
    markStageComplete: (
      stage: 'Информация' | 'Фото' | 'Интересы' | 'Тест',
    ) => void;
  } = $props();

  const form = superForm(defaults(zod4(testSchema)), {
    SPA: true,
    validators: zod4(testSchema),
    onUpdate({ form }) {
      if (form.valid) {
        // TODO: Call an external API with form.data, await the result and update form
      }
    },
    onSubmit() {
      console.log('Форма отправлена!');
    },
  });

  const { form: formData, enhance, validateForm } = form;

  export async function save() {
    window.Telegram.WebApp.MainButton.showProgress();
    const res = await updateData($formData).finally(
      window.Telegram.WebApp.MainButton.hideProgress,
    );
    if (res != 200) {
      console.log('failed to update user data');
      return;
    }
    return;
  }

  async function handleTelegramButtonClick() {
    // window.Telegram.WebApp.MainButton.showProgress()
    markStageComplete('Тест');
    setCurrentStage('Интересы');
    // await save();
    form.submit();
  }

  useTelegramButton(handleTelegramButtonClick);
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
  let sliders = $state([
    {
      leftCornerEmoji: '🏠',
      leftCornerText: 'Узкий круг',
      value: 5,
      rightCornerEmoji: '🥳',
      rightCornerText: 'Большие компании',
    },
    {
      leftCornerEmoji: '🔬',
      leftCornerText: 'Теория и анализ',
      value: 5,
      rightCornerEmoji: '🛠️',
      rightCornerText: 'Практика и действия',
    },
    {
      leftCornerEmoji: '📅',
      leftCornerText: 'Планирование и порядок',
      value: 5,
      rightCornerEmoji: '🎨',
      rightCornerText: 'Спонтанность и гибкость',
    },
    {
      leftCornerEmoji: '📢',
      leftCornerText: 'Организатор ',
      value: 5,
      rightCornerEmoji: '🤝',
      rightCornerText: 'Командный игрок',
    },
    {
      leftCornerEmoji: '🌍',
      leftCornerText: 'Разносторонние интересы',
      value: 5,
      rightCornerEmoji: '🎯',
      rightCornerText: 'Фокус на главном',
    },
  ]);
</script>

<form method="POST" use:enhance>
  <div class="flex flex-col gap-y-2">
    <div class="flex w-full items-center">
      <p class="font-medium text-2xl">
        <Emoji symbol="🎨" class="size-6" />Какой ты человек?
      </p>
    </div>
    <p class="mb-5">
      Пройди простой тест - так мы сможем понять тебя намного лучше, чем
      понимали бы без него!
    </p>
  </div>
  {#each sliders as slider (slider.rightCornerText)}
    <div class="mb-6.5">
      <div class="flex items-center mb-0.5">
        <Emoji symbol={slider.leftCornerEmoji} class="mr-1" />
        <p>{slider.leftCornerText}</p>
      </div>
      <Slider
        type="single"
        bind:value={slider.value}
        max={10}
        step={1}
        class="bg-accent-foreground/25 rounded-full"
      />
      <div class=" flex items-center justify-end mt-0.5">
        <Emoji symbol={slider.rightCornerEmoji} class="mr-1" />
        <p>{slider.rightCornerText}</p>
      </div>
    </div>
  {/each}
  <!--  <SuperDebug data={$formData} />-->
</form>

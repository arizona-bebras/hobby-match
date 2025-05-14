<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { Input } from '$lib/components/ui/input';
  import { X } from '@lucide/svelte';
  let { nextStage: open = $bindable(), addedWidget = $bindable() } = $props();

  function onOpenChange() {
    setTimeout(() => {
      document.body.style.cssText = '';
      console.log('Компонент уничтожен');
      window.Telegram.WebApp.MainButton.show();
    }, 10);
  }

  let options: Record<number, string> = $state({});
  let counter: number = $state(0);
  $inspect(options);

  function addOption() {
    counter += 1;
    options[counter] = '';
  }

  function removeOption(key: number) {
    delete options[key];
    options = options;
  }
</script>

<Sheet.Root bind:open {onOpenChange}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <p class="text-accent-foreground font-medium pb-2">Виджет "Опрос"</p>
      <Input placeholder="Напишите какой-нибудь вопрос" class="mb-2" />

      {#each Object.entries(options) as [key, value]}
        <div class="flex justify-center items-center mb-2">
          <Input
            bind:value={options[key]}
            placeholder={`Вариант ${key}`}
            class=""
          />
          <button onclick={() => removeOption(Number(key))}>
            <X class="size-5 ml-2" />
          </button>
        </div>
      {/each}

      <button
        class="text-accent-foreground font-bold underline underline-offset-3 decoration-2 flex"
        onclick={addOption}
      >
        Добавить вариант
      </button>
      <button
        onclick={() => {
          open = false;
          onOpenChange();
        }}
        class="w-full h-12 bg-accent rounded-xl"
      >
        Сохранить
      </button>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>

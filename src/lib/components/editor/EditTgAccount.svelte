<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { Input } from '$lib/components/ui/input';
  let { nextStage: open = $bindable(), addedWidget = $bindable() } = $props();

  function onOpenChange() {
    setTimeout(() => {
      document.body.style.cssText = '';
      console.log('Компонент уничтожен');
      window.Telegram.WebApp.MainButton.show();
    }, 10);
  }
</script>

<Sheet.Root bind:open {onOpenChange}>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <p class="text-accent-foreground font-medium pb-2">
        Виджет "Телеграм Аккаунт"
      </p>
      <p>Вы можете ввести ссылку на канал или личный аккаунт</p>
      <Input placeholder="https://t.me/..." class="mb-9" />
      <button
        onclick={() => {
          open = false;
          onOpenChange();
        }}
        class="w-full h-12 bg-accent rounded-xl">Сохранить</button
      >
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>

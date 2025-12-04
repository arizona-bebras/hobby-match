<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';

  let currentStage = $state(1);

  useTelegramButton(() => {
    currentStage += 1;
    if (currentStage === 3) {
      Telegram.WebApp.MainButton.text = 'Завершить';
    }
  });

  function onSheetOpenHandle() {
    Telegram.WebApp.MainButton.text = 'Далее';
  }

  $inspect(currentStage);
</script>

<Sheet.Root>
  <Sheet.Trigger
    class="px-6 py-1.5 rounded-[8px] bg-accent text-[14px]"
    onclick={() => {
      onSheetOpenHandle();
      console.log(12345);
    }}>Создать</Sheet.Trigger
  >
  <Sheet.Content side="bottom">
    <Sheet.Header>
      {#if currentStage === 1}
        <Sheet.Title>Are you sure absolutely sure?</Sheet.Title>
        <Sheet.Description>
          This action cannot be undone. This will permanently delete your
          account and remove your data from our servers.
        </Sheet.Description>
      {:else if currentStage === 2}
        <Sheet.Title>Тип неймспейса</Sheet.Title>
        <Sheet.Description>
          This action cannot be undone. This will permanently delete your
          account and remove your data from our servers.
        </Sheet.Description>
      {:else if currentStage === 3}
        <Sheet.Title>Страница с qr кодом</Sheet.Title>
        <Sheet.Description>
          This action cannot be undone. This will permanently delete your
          account and remove your data from our servers.
        </Sheet.Description>
      {/if}
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>
<!--<button class="px-6 py-1.5 rounded-[8px] bg-accent text-[14px]">Создать</button>-->

<script lang="ts">
  let { currentStage = $bindable() } = $props();
  import Emoji from '$lib/components/ui/emoji/emoji.svelte';
</script>

<div class="flex flex-col gap-y-2">
  <div class="flex w-full mb-2 items-center">
    <Emoji symbol="📷" />
    <p class="font-medium text-2xl h-10.25">Сфоткаемся?</p>
  </div>
  <button class="w-full h-12 bg-accent rounded-xl text-white font-medium"
    >Выбрать файл</button
  >
  <p class="self-center">или</p>
  <button class="w-full h-12 bg-accent rounded-xl text-white font-medium"
    >Взять текущую фотографию из Telegram</button
  >
  <button
    class="w-full h-12 bg-accent rounded-xl mt-12 text-white font-medium"
    onclick={() => (currentStage = 'interests')}>Продолжить</button
  >
</div>

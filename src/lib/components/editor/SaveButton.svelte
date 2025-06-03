<script lang="ts">
  let {
    onClick,
    isButtonActive,
  }: { onClick: () => void; isButtonActive: boolean } = $props();
</script>

<button
  type="button"
  onclick={onClick}
  disabled={!isButtonActive}
  class="w-full h-12 {isButtonActive ? 'bg-accent' : 'bg-inactive'} rounded-xl"
  >Сохранить</button
>

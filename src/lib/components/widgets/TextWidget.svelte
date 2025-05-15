<script lang="ts">
  import type { Text } from '$lib/widgetTypes/widgetTypes';

  let { data }: { data: Text } = $props();
</script>

<div class="TextBox text-container">
  <p class="mx-[10px] break-words font-[Inter] text-[16px] italic font-medium">
    {data.text}
  </p>
</div>

<style>
  .text-container {
    max-width: 300px;
  }
</style>

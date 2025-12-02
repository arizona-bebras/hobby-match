<script lang="ts">
  import type { Text } from '$lib/widgetTypes/widgetTypes';

  let { data }: { data: Text } = $props();
</script>

<div class="TextBox">
  <p class="break-words font-[Inter] text-[16px] italic font-medium">
    {data.text}
  </p>
</div>

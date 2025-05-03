<script lang="ts">
  let { data } = $props();
  console.log('DATA:', data.text);
</script>

<div class="TextBox text-container">
  <p class="m-[10px] break-words font-[Inter] text-[16px] italic font-medium">
    {data.text}
  </p>
</div>

<style>
  .text-container {
    max-width: 300px;
  }
</style>

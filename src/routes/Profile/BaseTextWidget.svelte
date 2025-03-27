<script lang="ts">
  let { children } = $props();
</script>

<div class="border-l-[2px] w-[100%] mb-[8px]">
  {@render children()}
</div>
<script lang="ts">
  let {open = $bindable()} = $props()
</script>

<main class="flex flex-col bg-red-200 absolute top-[50%] left-[50%] translate-x-[-50%] translate-y-[-50%]">
  <button type="button" class="btn preset-outlined-surface-500 p-5"><a href="#">Media</a></button>
  <button type="button" class="btn preset-outlined-surface-500 p-5"><a href="#">Link</a></button>
  <button type="button" class="btn preset-outlined-surface-500 p-5"><a href="#">Progress</a></button>
  <button type="button" class="btn preset-outlined-surface-500 p-5"><a href="#">Decorative</a></button>
  <button type="button" class="btn preset-outlined-surface-500 p-5" onclick={() => open=false}>Close</button>
</main>
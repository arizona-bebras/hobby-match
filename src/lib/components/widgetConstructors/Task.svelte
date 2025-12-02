<script lang="ts">
  let { onRemove, taskOrder } = $props();
</script>

<div class="flex-column">
  <p>Цель</p>
  <input type="text" name="task{taskOrder}" />
  <button onclick={onRemove}> Remove </button>
</div>

<style>
  input {
    outline: solid;
  }
  .flex-column {
    display: flex;
    flex-direction: column;
  }
</style>

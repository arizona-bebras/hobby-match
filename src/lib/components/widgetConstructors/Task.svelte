<script lang="ts">
    let { onRemove } = $props();
</script>

<div class="flex-column">
    <p>Цель</p>
    <input type="text">
    <button onclick={onRemove}>
		Remove
	</button>
</div>

<style>
    .flex-column {
        display: flex;
        flex-direction: column;
    }
</style>
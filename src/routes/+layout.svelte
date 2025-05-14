<script>
  import { pb } from '$lib';
  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';
  import '../app.css';
  onMount(async () => {
    // if (!pb.authStore.isValid) {
    //   await goto('/authError');
    // }
  });
</script>

<div class="w-screen max-w-full min-h-screen flex flex-col">
  <div class="flex flex-col bg-background flex-1 items-center text-text-color">
    <slot></slot>
  </div>
</div>

<script lang="ts">
  import { goto } from '$app/navigation';
  import { User, MessagesSquare, Eye } from '@lucide/svelte';
  import { page } from '$app/state';
  import NavBar from '$lib/components/NavBar.svelte';

  let { children } = $props();
</script>

<div class="w-screen max-w-full h-screen flex flex-col">
  <div
    class="flex-1 flex flex-col bg-background items-center text-text-color overflow-hidden"
  >
    {@render children()}
  </div>
  <NavBar />
</div>

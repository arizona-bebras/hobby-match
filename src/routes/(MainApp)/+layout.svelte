<script lang="ts">
  import { goto } from '$app/navigation';
  import { User, MessagesSquare, Eye } from '@lucide/svelte';
  import { page } from '$app/state';
  import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';
  import { SvelteQueryDevtools } from '@tanstack/svelte-query-devtools';
  import NavBar from '$lib/components/NavBar.svelte';

  let { children } = $props();
  const queryClient = new QueryClient();
</script>

<QueryClientProvider client={queryClient}>
  <div class="w-screen max-w-full h-screen flex flex-col">
    <div
      class="flex-1 flex flex-col bg-background items-center text-text-color overflow-hidden font-[Inter]"
    >
      {@render children()}
    </div>
    <NavBar />
  </div>
  <SvelteQueryDevtools />
</QueryClientProvider>

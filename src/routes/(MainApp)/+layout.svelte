<script lang="ts">
  import { goto } from '$app/navigation';
  import { User, MessagesSquare, Eye } from '@lucide/svelte';
  import { page } from '$app/state';
  import {
    QueryCache,
    QueryClient,
    QueryClientProvider,
  } from '@tanstack/svelte-query';
  import { SvelteQueryDevtools } from '@tanstack/svelte-query-devtools';
  import NavBar from '$lib/components/NavBar.svelte';
  import { toast } from 'svelte-sonner';

  let { children } = $props();
  const queryClient = new QueryClient({
    queryCache: new QueryCache({
      onSuccess: (data) => {
        if (data.error) {
          console.error('Возникла ошибка!');
          toast.error('Возникла ошибка при получении данных!');
        }
      },
    }),
  });
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

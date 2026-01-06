<script>
  import { Toaster } from '$lib/components/ui/sonner/index.js';
  let { children } = $props();
  import '../app.css';
  import { Search, Settings, UserRoundPen } from '@lucide/svelte';
  // onMount(async () => {
  //   // if (!pb.authStore.isValid) {
  //   //   await goto('/authError');
  //   // }
  // });
</script>

<Toaster position="top-center" richColors />
{@render children()}

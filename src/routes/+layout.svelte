<script>
  let { children } = $props();
  import '../app.css';
  import { Search, Settings, UserRoundPen } from '@lucide/svelte';
  // onMount(async () => {
  //   // if (!pb.authStore.isValid) {
  //   //   await goto('/authError');
  //   // }
  // });
</script>

<slot></slot>

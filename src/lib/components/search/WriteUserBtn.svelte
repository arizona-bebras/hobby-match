<script lang="ts">
  import { UserPen, Mail } from '@lucide/svelte';
  let { username }: { username: string } = $props();
  // window.Telegram.WebApp.openTelegramLink('tg://user?id=5000721686');
  let href = $derived(`https://t.me/${username}`);
</script>

<a
  {href}
  class="fixed bottom-20 right-5.5 bg-accent size-12.5 flex items-center justify-center rounded-xl"
>
  <Mail class="size-6 text-text-color" />
</a>

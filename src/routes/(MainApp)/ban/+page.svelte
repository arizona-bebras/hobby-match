<script>
  import { Ban } from '@lucide/svelte';
  import { page } from '$app/state';
</script>

<div
  class="w-[100vw] h-[100vh] flex flex-col items-center justify-center p-3 text-center gap-2"
>
  <Ban class="size-15" />
  <h1 class="font-bold text-xl">Аккаунт заблокирован!</h1>
  <p>
    Твоя анкета нарушает правила сообщества, поэтому мы её заблокировали. Если
    ты не соглас{'male' === 'female' ? 'на' : 'ен'} с этим решением, обратись на
    почту
    <a href="mailto:ban@shumi.space" class="underline">ban@shumi.space</a>.
  </p>
  {#if page.url.searchParams.has('reason')}
    <p><b>Причина:</b> {page.url.searchParams.get('reason')}</p>
  {/if}
</div>

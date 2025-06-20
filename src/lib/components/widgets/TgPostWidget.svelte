<script lang="ts">
  import type { Post } from '$lib/widgetTypes/widgetTypes';
  let { data }: { data: Post } = $props();
  const telegramIframe = (div: HTMLDivElement) => {
    let textColor = window.Telegram.WebApp.themeParams.accent_text_color;
    let theme = window.Telegram.WebApp.colorScheme;
    div!.innerHTML = '';

    const script = document.createElement('script');
    script.src = 'https://telegram.org/js/telegram-widget.js?22';
    script.setAttribute('async', 'true');
    script.setAttribute('data-telegram-post', data.link);
    script.setAttribute('data-width', '100%');
    script.setAttribute('data-userpic', 'false');
    script.setAttribute('data-color', textColor!);
    script.setAttribute('data-dark-color', textColor!);
    script.setAttribute('data-dark', theme === 'dark' ? '1' : '0');

    div!.appendChild(script);
    script.onload = () => {
      const frames = div.getElementsByTagName('iframe');
      for (const frame of frames) {
        frame.setAttribute('sandbox', '');
      }
    };
  };
</script>

{#key data}
  <button
    onclick={() =>
      window.Telegram.WebApp.openTelegramLink(`https://t.me/${data.link}`)}
    class="w-full"
    aria-label="Пост"
  >
    <div
      use:telegramIframe
      class="w-full overflow-hidden pointer-events-none"
    ></div>
  </button>
{/key}

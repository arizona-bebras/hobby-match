<script lang="ts">
  import { onMount } from 'svelte';
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
  };
</script>

<div style="width: 100%;" use:telegramIframe></div>

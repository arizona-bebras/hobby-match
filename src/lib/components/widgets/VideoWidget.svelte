<script lang="ts">
  import { getYouTubeVideoId } from '$lib/components/editor/schemes/videoSheme';

  let { data } = $props();
  let videId = getYouTubeVideoId(data.link);

  let url;
  switch (data.platform) {
    case 'YouTube':
      url = `https://www.youtube.com/embed/${videId}`;
      break;
    case 'Rutube':
      url = `https://rutube.ru/play/embed/${videId}`;
      break;
    case 'TikTok':
      url = `https://www.tiktok.com/player/v1/${videId}`;
      break;
  }
</script>

<button
  class="TextBox"
  onclick={() => window.Telegram.WebApp.openLink(data.link)}
  aria-label="Видео"
>
  <iframe
    class="w-full aspect-video overflow-hidden pointer-events-none"
    scrolling="no"
    src={url}
    title="Видео"
    frameborder="0"
    referrerpolicy="strict-origin-when-cross-origin"
  ></iframe>
</button>

<script lang="ts">
  import { getYouTubeVideoId } from '$lib/components/editor/video/videoSheme';
  import type { Video } from '$lib/widgetTypes/widgetTypes';
  import test from '$lib/assets/DuckHello.png?h=100';

  let { data, TestVideo = null }: { data: Video; TestVideo?: null | any } =
    $props();
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

{#if TestVideo}
  <video width="320" height="256" controls poster={test}>
    <source src={TestVideo} type="video/mp4" />
    Your browser does not support the video tag.
  </video>
{:else}
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
{/if}

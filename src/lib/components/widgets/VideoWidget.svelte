<script lang="ts">
  import { getYouTubeVideoId } from '$lib/components/editor/schemes/videoSheme';

  let { data } = $props();
  let videId = getYouTubeVideoId(data.link);

  let youtubeURL = `https://www.youtube.com/embed/${videId}`;
  let rutubeURL = `https://rutube.ru/play/embed/${videId}`;
  let tiktokURL = `https://www.tiktok.com/player/v1/${videId}`;
</script>

<div class="TextBox">
  {#if data.platform === 'YouTube'}
    <iframe
      class="w-full aspect-video overflow-hidden"
      scrolling="no"
      src={youtubeURL}
      title="YouTube Video"
      frameborder="0"
      allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
      referrerpolicy="strict-origin-when-cross-origin"
      allowfullscreen
    ></iframe>
  {:else if data.platform === 'Rutube'}
    <iframe
      class="w-full aspect-video overflow-hidden"
      scrolling="no"
      title="Rutube Video"
      src={rutubeURL}
      frameBorder="0"
      allow="clipboard-write; autoplay"
      allowFullScreen
    ></iframe>
  {:else if data.platform === 'TikTok'}
    <iframe
      class="w-full aspect-square overflow-hidden"
      scrolling="no"
      title="TikTok Video"
      src={tiktokURL}
    ></iframe>
  {/if}
</div>

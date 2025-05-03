<script lang="ts">
  let { data } = $props();
</script>

<div class="TextBox">
  {#if data.platform === 'YouTube'}
    <iframe
      class="w-full aspect-video"
      src={data.link}
      title="YouTube Video"
      frameborder="0"
      allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
      referrerpolicy="strict-origin-when-cross-origin"
      allowfullscreen
    ></iframe>
  {:else if data.platform === 'Rutube'}
    <iframe
      class="w-full aspect-video"
      title="Rutube Video"
      src={data.link}
      frameBorder="0"
      allow="clipboard-write; autoplay"
      allowFullScreen
    ></iframe>
  {:else if data.platform === 'TikTok'}
    <iframe class="w-full aspect-square" title="TikTok Video" src={data.link}
    ></iframe>
  {/if}
</div>

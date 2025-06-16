<script lang="ts">
  let { data } = $props();

  async function getEmbedCode(trackUrl: string) {
    const response = await fetch(
      `https://soundcloud.com/oembed?format=json&url=${encodeURIComponent(trackUrl)}`,
    );
    const data = await response.json();
    console.log(data);

    const parser = new DOMParser();
    const doc = parser.parseFromString(data.html, 'text/html');
    const iframe = doc.querySelector('iframe');

    if (!iframe) return data.html;

    iframe.setAttribute('height', '166');
    iframe.style.overflow = 'hidden';
    iframe.style.pointerEvents = 'none';

    const src = new URL(iframe.src);
    src.searchParams.set('hide_related', 'true');
    src.searchParams.set('show_comments', 'false');
    src.searchParams.set('show_teaser', 'false');
    iframe.src = src.toString();
    console.log(iframe);
    return iframe.outerHTML;
  }

  let trackIframePromise = getEmbedCode(data.link);
</script>

{#await trackIframePromise}
  <p></p>
{:then trackIframe}
  {@html trackIframe}
{/await}

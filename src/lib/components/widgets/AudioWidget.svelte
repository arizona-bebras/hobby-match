<script lang="ts">
  let { data } = $props();
  let track = $state();

  async function getScEmbed(trackUrl: string) {
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

  async function getYandexTrack(trackUrl: string) {
    console.log(trackUrl)
    const iframe = document.createElement("iframe")
    iframe.setAttribute('src', `https://music.yandex.ru/iframe/${trackUrl.slice(23)}`)
    iframe.setAttribute('height', '130');
    iframe.setAttribute('width', '500');
    console.log(iframe.outerHTML)
    return iframe.outerHTML;
  }

  async function getSpotifyEmbed(trackUrl: string) {
    const data = await fetch(
      `https://open.spotify.com/oembed?url=${encodeURIComponent(trackUrl)}`
    ).then(response => response.json())
    return data.html!
  }


  switch (data.platform) {
    case 'SoundCloud':
      track = getScEmbed(data.link);
      break
    case 'Yandex':
      track = getYandexTrack(data.link);
      break
    case 'Spotify':
      track = getSpotifyEmbed(data.link);
      break
  }
</script>

<button
  onclick={() => window.Telegram.WebApp.openLink(data.link)}
  class="w-full"
>
  {#await track}
    <p>Загрузка...</p>
  {:then trackIframe}
    <div class="max-w-[350px] overflow-hidden rounded-xl">
      {@html trackIframe}
    </div>
  {/await}
</button>

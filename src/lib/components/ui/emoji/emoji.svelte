<script lang="ts">
  let { symbol } = $props();
  function getAppleEmogi(emoji: string) {
    let code: string[] = [];
    for (const codePoint of emoji) {
      code.push(codePoint.codePointAt(0)!.toString(16));
    }
    return `https://cdnjs.cloudflare.com/ajax/libs/emoji-datasource-apple/15.1.2/img/apple/64/${code.join('-')}.png`;
  }
</script>

<img src={getAppleEmogi(symbol)} alt={symbol} class="size-6 mr-2 inline" />

<script lang="ts">
  import { cn } from '$lib/utils';

  let { symbol, size = 4, class: className = '' } = $props();
  function getAppleEmoji(emoji: string) {
    let code: string[] = [];
    for (const codePoint of emoji) {
      code.push(codePoint.codePointAt(0)!.toString(16));
    }
    return `https://cdnjs.cloudflare.com/ajax/libs/emoji-datasource-apple/15.1.2/img/apple/64/${code.join('-')}.png`;
  }
</script>

<img
  src={getAppleEmoji(symbol)}
  alt={symbol}
  class={cn('mr-2 inline', className)}
  style:width="{size * 4}px"
  style:height="{size * 4}px"
/>

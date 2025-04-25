<script>
  let { image, left, top, rotate } = $props();
</script>

<img src="{image}" alt="SomeImage" class="absolute z-1" style:top={top}px style:left={left}px
     style:transform={`rotate(${rotate}deg)`}>

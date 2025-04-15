<script lang="ts">
  let { img, game, hour } = $props();
</script>

<div class="GameBox box-border">
  <img src="{img}" alt="Valorant Image">
  <div class="flex flex-col font-[Inter] font-medium">
    <p class="text-[20px]">Время в {game}</p>
    <p class="text-[19px]">{hour} часов</p>
  </div>
</div>

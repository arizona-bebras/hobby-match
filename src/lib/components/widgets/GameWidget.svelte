<script lang="ts">
  let { img, game, hour } = $props();
</script>

<div class="GameBox box-border">
  <img src={img} alt="Valorant Image" class="size-10 align-center my-auto" />
  <div class="flex flex-col font-[Inter] font-medium justify-center">
    <p class="text-[16px] font-semibold">Время в {game}</p>
    <p class="text-[14px]">{hour} часов</p>
  </div>
</div>

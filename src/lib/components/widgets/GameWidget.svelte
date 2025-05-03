<script lang="ts">
  let { data } = $props();
</script>

<div class="GameBox box-border">
  <img
    src={data.game_icon}
    alt="Valorant Image"
    class="size-10 align-center my-auto"
  />
  <div class="flex flex-col font-[Inter] font-medium justify-center">
    <p class="text-[16px] font-semibold">Время в {data.game}</p>
    <p class="text-[14px]">{data.hours_played} часов</p>
  </div>
</div>

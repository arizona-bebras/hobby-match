<script lang="ts">
  import type { SteamGameData } from '$lib/widgetTypes/widgetTypes';

  let { data }: { data: SteamGameData } = $props();
  console.log(1111111111111111, data.icon);
</script>

<div class="GameBox box-border">
  <img
    src={data.icon}
    alt="Valorant Image"
    class="size-10 rounded-lg align-center my-auto"
  />
  <div
    class="flex flex-col font-[Inter] font-medium justify-center pl-3 min-w-0"
  >
    <p class="text-[16px] font-semibold truncate">Время в {data.title}</p>
    <p class="text-[14px]">{data.hours_played} часов</p>
  </div>
</div>

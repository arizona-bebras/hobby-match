<script lang="ts">
  import type { SteamGameData } from '$lib/widgetTypes/widgetTypes';
  import { TriangleAlert } from '@lucide/svelte';

  let { data }: { data: SteamGameData } = $props();
  console.log(1111111111111111, data);
</script>

<div class="GameBox box-border">
  {#if Object.hasOwn(data, 'title')}
    <img
      src={data.icon}
      alt={data.title}
      class="size-10 rounded-lg align-center my-auto"
    />
    <div
      class="flex flex-col font-[Inter] font-medium justify-center pl-3 min-w-0"
    >
      <p class="text-[16px] font-semibold truncate">Время в {data.title}</p>
      <p class="text-[14px]">{data.hours_played} часов</p>
    </div>
  {:else}
    <div
      class="size-10 rounded-full bg-background flex items-center justify-center"
    >
      <TriangleAlert class="stroke-destructive" />
    </div>
    <div
      class="flex flex-col font-[Inter] font-medium justify-center pl-3 min-w-0"
    >
      <p class="text-[16px] font-semibold truncate text-destructive">
        Ошибка получения данных
      </p>
      <p class="text-[14px] text-text-color">
        Не удалось получить данные об игре
      </p>
    </div>
  {/if}
</div>

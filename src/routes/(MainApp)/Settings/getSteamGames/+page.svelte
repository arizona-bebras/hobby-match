<script lang="ts">
  let { data } = $props();
  let selectedGame: string = $state('');
</script>

<main>
  <form class="flex flex-col" method="POST" action="?/Settings">
    <label class="label">
      <span class="label-text">Steam link to profile</span>
      <input class="input" type="text" placeholder="Input" name="link" />
      <button class="ig-btn preset-filled">Submit</button>
    </label>
  </form>
  {#if typeof data.games !== 'undefined'}
    <select class="select" bind:value={selectedGame}>
      {#each JSON.parse(data.games) as info}
        <option
          value="{info.name}, {Math.trunc(
            info.playtime_forever / 60,
          )}, {info.appid}, {info.img_icon_url}">{info.name}</option
        >
      {/each}
    </select>
    <p>selected game is: {selectedGame ? selectedGame : 'waiting...'}</p>
    <div class="p-[15px] w-full">
      <div
        class="flex gap-[12px] bg-[#c2c2c2] p-[10px] rounded-[20px] border-[2px] border-solid border-[#7c7c7c] mb-[8px] w-full h-20"
      >
        <img
          src={selectedGame.split(',')[3]}
          alt="Game Image"
          class="w-[50px] h-[50px]"
        />
        <div class="flex flex-col font-[Inter] font-medium">
          <p class="text-[20px] truncate">
            Время в {selectedGame.split(',')[0]}
          </p>
          <p class="text-[19px]">{selectedGame.split(',')[1]} часов</p>
        </div>
      </div>
    </div>
  {/if}
</main>

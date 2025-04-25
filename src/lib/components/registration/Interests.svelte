<script lang="ts">
  import { Input } from '$lib/components/ui/input/index.js';
  import { Plus } from '@lucide/svelte';
  import Emoji from '$lib/components/ui/emoji/emoji.svelte';
  function getAppleEmogi(emoji: string) {
    let code: string[] = [];
    for (const codePoint of emoji) {
      code.push(codePoint.codePointAt(0).toString(16));
    }
    return `https://cdnjs.cloudflare.com/ajax/libs/emoji-datasource-apple/15.1.2/img/apple/64/${code.join('-')}.png`;
  }
  let listOfInterests = $state(['Hello', 'World', 'Yep', 'I', 'LP']);
  let selectedInterests: string[] = $state([]);
  let userInterest: string = $state('');
</script>

<div class="flex flex-col gap-y-2">
  <div class="flex w-full mb-2 items-center">
    <p class="font-medium text-2xl h-15.25">
      <Emoji symbol="💗" />Давай узнаем друг-друга поближе?
    </p>
  </div>
  <p>Выберите наиболее интересующие вас темы (максимум 5)</p>
  <div class="flex flex-col gap-y-4">
    <Input
      type="text"
      placeholder="Начните вводить"
      bind:value={userInterest}
      onkeydown={(e) => {
        if (e.key === 'Enter') {
          listOfInterests.push(userInterest);
          console.log(listOfInterests);
        }
      }}
    />
    <div class="flex flex-row gap-2 w-full flex-wrap font-medium">
      {#each listOfInterests as element}
        <button
          onclick={() => {
            if (selectedInterests.includes(element)) {
              selectedInterests.splice(selectedInterests.indexOf(element), 1);
            } else selectedInterests.push(element);
          }}
          class="{selectedInterests.includes(element)
            ? 'bg-accent'
            : 'bg-accent/25'} rounded-3xl flex flex-row items-center justify-center px-3 py-2 gap-1.5"
        >
          <Plus
            class="{selectedInterests.includes(element)
              ? 'text-white rotate-45'
              : 'text-accent-foreground'} w-4 h-5 stroke-3"
          />
          <p
            class="text-accent-foreground {selectedInterests.includes(element)
              ? 'text-white'
              : 'text-accent-foreground'}"
          >
            {element}
          </p>
        </button>
      {/each}
    </div>
  </div>
  <button
    onclick={() => {
      console.log(selectedInterests);
    }}>Info</button
  >
</div>

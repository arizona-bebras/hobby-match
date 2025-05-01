<script lang="ts">
  import { browser } from '$app/environment';
  import Steam from '$lib/components/SocialMedia/Steam.svelte';
  import Twitch from '$lib/components/SocialMedia/Twitch.svelte';
  import Twitter from '$lib/components/SocialMedia/Twitter.svelte';
  import Vk from '$lib/components/SocialMedia/Vk.svelte';
  import Youtube from '$lib/components/SocialMedia/Youtube.svelte';

  import { ArrowUpRight } from '@lucide/svelte';

  function GetCorrectForm(number: number, words_arr: string[] | string) {
    number = Math.abs(number);
    if (Number.isInteger(number)) {
      let options = [2, 0, 1, 1, 1, 2];
      return words_arr[
        number % 100 > 4 && number % 100 < 20
          ? 2
          : options[number % 10 < 5 ? number % 10 : 5]
      ];
    }
    return words_arr[1];
  }

  let SubscribersType = {
    Youtube: [['подписчик', 'подписчика', 'подписчиков']],
    VK: [['подписчик', 'подписчика', 'подписчиков']],
    Steam: [['уровень', 'уровень', 'уровень']], // :))
    Twitch: [['подписчик', 'подписчика', 'подписчиков']],
    Twitter: [['читатель', 'читателя', 'читателей']],
  };
  const SocialIcons = {
    Youtube: Youtube,
    VK: Vk,
    Steam: Steam,
    Twitch: Twitch,
    Twitter: Twitter,
  };
  let {
    platform,
    url,
    username,
    amountSubscribers,
  }: {
    platform: keyof typeof SubscribersType;
    url: string;
    username: string;
    amountSubscribers: number;
  } = $props();
  const Icon = SocialIcons[platform];
</script>

<button
  class="w-full"
  onclick={() => {
    if (isTMA()) {
      console.log(openLink.isAvailable());
      if (openLink.isAvailable()) {
        console.log(openLink.isAvailable());
        openLink(url, {
          tryBrowser: 'chrome',
          tryInstantView: true,
        });
      }
    }
  }}
>
  <div class="GameBox">
    <div class="size-10 rounded-xl my-auto h-fit fill-accent">
      <!--      <svelte:component this={?Icon}></svelte:component>-->
      <Icon></Icon>
    </div>
    <div class="flex flex-col font-[Inter] font-medium text-start">
      <p class="text-[16px] max-w-60 truncate">{username}</p>
      <p class="text-[14px]">
        {amountSubscribers}
        {GetCorrectForm(amountSubscribers, SubscribersType[platform][0])}
      </p>
    </div>
    <ArrowUpRight class="self-center m-auto mr-0"></ArrowUpRight>
  </div>
</button>

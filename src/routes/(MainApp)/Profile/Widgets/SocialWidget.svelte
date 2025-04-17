<script lang="ts">
    import { openLink } from '@telegram-apps/sdk-svelte';
    import { init } from '@telegram-apps/sdk-svelte';
    import { isTMA } from '@telegram-apps/bridge';
    import { browser } from '$app/environment';

    console.log(browser);
    if (browser) {
        init();
    }


    import { ArrowUpRight } from '@lucide/svelte';

    function GetCorrectForm(number: number, words_arr: string[] | string) {
        number = Math.abs(number);
        if (Number.isInteger(number)) {
            let options = [2, 0, 1, 1, 1, 2];
            return words_arr[(number % 100 > 4 && number % 100 < 20) ? 2 : options[(number % 10 < 5) ? number % 10 : 5]];
        }
        return words_arr[1];
    }

    let SubscribersType = {
        'Youtube': [["подписчик", "подписчика", "подписчиков"], 'SocialMedia/Youtube.svg'],
        'VK': [["подписчик", "подписчика", "подписчиков"], 'SocialMedia/Vk.svg'],
        'Steam': [['уровень', 'уровень', 'уровень'], 'SocialMedia/Steam.svg'], // :))
        'Twitch': [["подписчик", "подписчика", "подписчиков"], 'SocialMedia/Twitch.svg'],
        'Twitter': [["читатель", "читателя", "читателей"], 'SocialMedia/Twitter.svg'],
    }
    let { type, url, userName, amountSubscribers }: {type:keyof typeof SubscribersType, url:string, userName:string, amountSubscribers: number} = $props();
</script>

<button class="w-full" onclick={() => {
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
}}>
  <div class="GameBox">
    <img src="{SubscribersType[type][1]}" alt="Valorant Image" class="size-10 rounded-xl my-auto">
    <div class="flex flex-col font-[Inter] font-medium text-start">
      <p class="text-[16px] max-w-60 truncate">{userName}</p>
      <p class="text-[14px]">{amountSubscribers} {GetCorrectForm(amountSubscribers, SubscribersType[type][0])}</p>
    </div>
      <ArrowUpRight class="self-center m-auto mr-0"></ArrowUpRight>
  </div>
</button>

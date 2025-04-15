<script lang="ts">
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
        'VK': [["подписчик", "подписчика", "подписчиков"], 'SocialMedia/VK.png'],
        'Steam': [['уровень', 'уровень', 'уровень'], 'SocialMedia/Steam.svg'], // :))
        'Twitch': [["подписчик", "подписчика", "подписчиков"], 'SocialMedia/Twitch.png'],
        'Twitter': [["читатель", "читателя", "читателей"], 'SocialMedia/Twitter.svg'],
    }
    let { type, url, amountSubscribers }: {type:keyof typeof SubscribersType, url:string, amountSubscribers: number} = $props();
</script>

<a href="#">
  <div class="GameBox">
    <img src="{SubscribersType[type][1]}" alt="Valorant Image" class="w-14.25 h-14.25 rounded-xl">
    <div class="flex flex-col font-[Inter] font-medium">
      <p class="text-[20px] max-w-60 truncate">{url}</p>
      <p class="text-[19px]">{amountSubscribers} {GetCorrectForm(amountSubscribers, SubscribersType[type][0])}</p>
    </div>
    <button class="self-center m-auto mr-0">
      <ArrowUpRight/>
    </button>
  </div>
</a>

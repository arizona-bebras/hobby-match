<script lang="ts">
  import Steam from '$lib/components/SocialMedia/Steam.svelte';
  import Twitch from '$lib/components/SocialMedia/Twitch.svelte';
  import Twitter from '$lib/components/SocialMedia/Twitter.svelte';
  import Vk from '$lib/components/SocialMedia/Vk.svelte';
  import Youtube from '$lib/components/SocialMedia/Youtube.svelte';
  import TgAccount from '$lib/components/SocialMedia/TgAccount.svelte';

  import { ArrowUpRight } from '@lucide/svelte';
  import type {
    SocialMediaData,
    SocialMediaLink,
  } from '$lib/widgetTypes/widgetTypes';

  function getCorrectForm(number: number, words_arr: string[] | string) {
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

  function formatNumber(num: number): string {
    const absNum = Math.abs(num);

    if (absNum >= 1000000) {
      return Math.round(absNum / 100000) / 10 + ' млн.';
    }
    if (absNum >= 1000) {
      return Math.round(absNum / 100) / 10 + ' тыс.';
    }
    return num.toString();
  }

  let {
    data,
    socialMediaData,
  }: { data: SocialMediaLink; socialMediaData: SocialMediaData } = $props();
</script>

<button
  class="w-full"
  onclick={() => {
    if (data.platform === 'Telegram') {
      window.Telegram.WebApp.openTelegramLink(data.link);
    } else {
      window.Telegram.WebApp.openLink(data.link);
    }
  }}
>
  <div class="GameBox">
    <div class=" w-10 rounded-xl my-auto fill-accent">
      {#if data.platform === 'YouTube'}
        <Youtube />
      {:else if data.platform === 'Twitch'}
        <Twitch />
      {:else if data.platform === 'VK'}
        <Vk />
      {:else if data.platform === 'Steam'}
        <Steam />
      {:else if data.platform === 'X'}
        <Twitter />
      {:else if data.platform === 'Telegram'}
        <TgAccount />
      {/if}
      <!--      <svelte:component this={?Icon}></svelte:component>-->
    </div>
    <div class="flex pl-3 flex-col font-[Inter] font-medium text-start">
      <p class="text-[16px] max-w-60 truncate">
        {#if socialMediaData.type === 'YouTube' || socialMediaData.type === 'Twitch'}
          {socialMediaData.title}
        {:else if socialMediaData.type === 'Steam'}
          {socialMediaData.username}
        {:else}
          dev
        {/if}
      </p>

      <p class="text-[14px]">
        {#if socialMediaData.type === 'VK' || socialMediaData.type === 'Twitch'}
          {socialMediaData.followers}
          {getCorrectForm(socialMediaData.followers ?? 0, [
            'подписчик',
            'подписчика',
            'подписчиков',
          ])}
        {:else if socialMediaData.type === 'Steam'}
          {socialMediaData.level}
          {getCorrectForm(socialMediaData.level ?? 0, [
            'уровень',
            'уровень',
            'уровень',
          ])}
        {:else if socialMediaData.type === 'YouTube'}
          {formatNumber(socialMediaData.subscribers ?? 0)}
          {getCorrectForm(socialMediaData.subscribers ?? 0, [
            'подписчик',
            'подписчика',
            'подписчиков',
          ])}
        {:else}
          dev
        {/if}
      </p>
    </div>
    <ArrowUpRight class="self-center m-auto mr-0"></ArrowUpRight>
  </div>
</button>

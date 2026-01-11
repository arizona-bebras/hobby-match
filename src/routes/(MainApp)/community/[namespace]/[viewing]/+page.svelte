<script lang="ts">
  import Questionnaire from '$lib/components/profile/Questionnaire.svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import { ScrollState } from 'runed';
  import { page } from '$app/state';
  import { createQuery } from '@tanstack/svelte-query';
  import client from '$lib/api/client';
  import WriteUserBtn from '$lib/components/search/WriteUserBtn.svelte';
  import { userData } from '$lib/storage/userData.svelte';
  import { onMount, untrack } from 'svelte';

  let testData: PageData = {
    age: 19,
    gender: 'male',
    id: '621zj74d6h1fjm3',
    interests: [
      'Программирование',
      'Музыка',
      'Звукорежиссура',
      'Литература',
      'Фотография',
    ],
    location: 'Екатеринбург',
    miniapp_name: 'Эмиль',
    user_info: 'Программирую, пишу музыку и занимаюсь копирайтингом.\n',
    user_photo:
      'https://www.soyuz.ru/public/uploads/files/2/7480281/20220315190534af66e2c5d3.jpg',
    widgets: [
      {
        collectionId: '9fxx39w3wb6ulbf',
        collectionName: 'widgets',
        created: '2025-06-20 19:35:02.483Z',
        data: {
          link: 'lnhly/1942',
          type: 'post',
        },
        files: [],
        id: 'y08nut634rhb4sv',
        order: 0,
        updated: '2025-06-20 19:35:02.483Z',
        user: '621zj74d6h1fjm3',
      },
      {
        collectionId: '9fxx39w3wb6ulbf',
        collectionName: 'widgets',
        created: '2025-06-19 20:59:32.013Z',
        data: {
          text: '«Я согласна рисковать, но ради чего-то стоящего»',
          type: 'text',
        },
        files: [],
        id: 'v62efkuvx4yg0bu',
        order: 1,
        updated: '2025-07-16 21:44:09.963Z',
        user: '621zj74d6h1fjm3',
      },
      {
        additionalData: {},
        collectionId: '9fxx39w3wb6ulbf',
        collectionName: 'widgets',
        created: '2025-06-19 21:03:11.204Z',
        data: {
          link: 'https://t.me/lnhly',
          platform: 'Telegram',
          type: 'social_media',
        },
        files: [],
        id: 'hl6rh9bf6jo5j1q',
        order: 2,
        updated: '2025-06-19 21:03:47.900Z',
        user: '621zj74d6h1fjm3',
      },
    ],
  };
  let { data } = $props();
  console.log('userID/Data:', data.pageData);
  // const profileData = createQuery(() => ({
  //   queryKey: ['profileData'],
  //   queryFn: async () =>
  //     await client.GET('/api/pages/{page_id}', {
  //       params: {
  //         path: {
  //           page_id: userId,
  //         },
  //       },
  //     }),
  //   select: (data) => data.data,
  // }));
  let screenContainer = $state<HTMLElement>();
  let scroll = new ScrollState({
    element: () => screenContainer,
  });
  let isDataLoad = $state(false);
  window.Telegram.WebApp.MainButton.hide();
  onMount(async () => {
    console.log(data.pageData.data.widgets, typeof data.pageData.data);
    if (data.pageData) {
      userData.current.last_viewed_profile = data.pageData.data.tg_user;
      await serializeWidgetData();
      isDataLoad = true;
    }
  });
  // $effect(() => {
  //   if (profileData.isSuccess) {
  //     untrack(
  //       () =>
  //         (userData.current.last_viewed_profile = profileData.data?.tg_user),
  //     );
  //   }
  // });
  async function serializeWidgetData() {
    for (let widgetData of data!.pageData!.data!.widgets) {
      widgetData.data = JSON.parse(widgetData.data);
      if (widgetData.additionalData) {
        widgetData.additionalData = JSON.parse(widgetData.additionalData);
      }
      console.log('Спаршенные данные:', widgetData);
    }
  }
</script>

{#if data.pageData && isDataLoad}
  <div class="overflow-y-auto relative" bind:this={screenContainer}>
    <Questionnaire
      data={data.pageData.data}
      isNamespaceProfile={true}
      {scroll}
    />
    <WriteUserBtn username={data.pageData?.data?.username} />
  </div>
{:else}
  <div class="flex flex-col items-center items-center mt-5">
    <p class="text-slate-400">Вы еще не посмотрели ни одной анкеты</p>
  </div>
{/if}

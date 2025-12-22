<script lang="ts">
  import Questionnaire from '$lib/components/profile/Questionnaire.svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import { ScrollState } from 'runed';
  import { page } from '$app/state';
  import { createQuery } from '@tanstack/svelte-query';
  import client from '$lib/api/client';

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
  const userId = page.url.href.split('/').at(-1);
  console.log(userId);
  const profileData = createQuery(() => ({
    queryKey: ['profileData'],
    queryFn: async () =>
      await client.GET('/api/pages/{page_id}', {
        params: {
          path: {
            page_id: userId,
          },
        },
      }),
    select: (data) => data.data,
  }));
  console.log(profileData.data);
  let screenContainer = $state<HTMLElement>();
  let scroll = new ScrollState({
    element: () => screenContainer,
  });
</script>

{#if profileData.data}
  <div class="overflow-y-auto relative" bind:this={screenContainer}>
    <Questionnaire data={profileData.data} isNamespaceProfile={true} {scroll} />
  </div>
{/if}

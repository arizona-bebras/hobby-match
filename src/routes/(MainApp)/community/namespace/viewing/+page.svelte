<script lang="ts">
  import Questionnaire from '$lib/components/profile/Questionnaire.svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import { ScrollState } from 'runed';

  let testData: PageData = {
    age: 19,
    gender: 'male',
    id: '621zj74d6h1fjm3',
    interests: [
      {
        collectionId: 'pbc_2896159602',
        collectionName: 'interests',
        created: '2025-06-06 14:02:21.268Z',
        id: '47jy5p4x3f7otu2',
        tag: 'Программирование',
        updated: '2025-06-06 14:02:21.268Z',
      },
      {
        collectionId: 'pbc_2896159602',
        collectionName: 'interests',
        created: '2025-06-06 14:02:04.940Z',
        id: 'w5jl33kj0c9k62h',
        tag: 'Музыка',
        updated: '2025-06-06 14:02:04.940Z',
      },
      {
        collectionId: 'pbc_2896159602',
        collectionName: 'interests',
        created: '2025-06-06 14:04:01.794Z',
        id: '1442kq32c40p6b0',
        tag: 'Звукорежиссура',
        updated: '2025-06-06 14:04:01.794Z',
      },
      {
        collectionId: 'pbc_2896159602',
        collectionName: 'interests',
        created: '2025-06-06 14:02:06.676Z',
        id: '3crhy34ry582qv6',
        tag: 'Литература',
        updated: '2025-06-06 14:02:06.676Z',
      },
      {
        collectionId: 'pbc_2896159602',
        collectionName: 'interests',
        created: '2025-06-06 14:01:31.156Z',
        id: '87g4o49y621cd50',
        tag: 'Фотография',
        updated: '2025-06-06 14:01:31.156Z',
      },
    ],
    location: 'Екатеринбург',
    miniapp_name: 'Эмиль',
    user_info: 'Программирую, пишу музыку и занимаюсь копирайтингом.\n',
    user_photo: 'avatar_xskzlm0x9g.jpg',
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
  let screenContainer = $state<HTMLElement>();
  let scroll = new ScrollState({
    element: () => screenContainer,
  });
</script>

<div class="overflow-y-auto" bind:this={screenContainer}>
  <Questionnaire data={testData} isNamespaceProfile={true} {scroll} />
</div>

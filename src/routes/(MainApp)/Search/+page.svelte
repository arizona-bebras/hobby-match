<script lang="ts">
  import { scrollY } from 'svelte/reactivity/window';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  import { goto } from '$app/navigation';
  import Questionnaire from '$lib/components/profile/Questionnaire.svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import { onMount } from 'svelte';

  const userProfiles: PageData[] = [
    {
      age: 0,
      gender: 'male',
      id: 'xb25g6o0oy91395',
      interests: [
        {
          collectionId: 'pbc_2896159602',
          collectionName: 'interests',
          created: '2025-06-06 14:04:37.575Z',
          id: '5d3keflcrh7zf0c',
          tag: 'Ролевые игры',
          updated: '2025-06-06 14:04:37.575Z',
        },
        {
          collectionId: 'pbc_2896159602',
          collectionName: 'interests',
          created: '2025-06-06 14:04:35.981Z',
          id: 'ts62vu2ii86l189',
          tag: 'Оригами',
          updated: '2025-06-06 14:04:35.981Z',
        },
        {
          collectionId: 'pbc_2896159602',
          collectionName: 'interests',
          created: '2025-06-06 14:04:33.754Z',
          id: '654q212ytq694nu',
          tag: 'Коллекционирование',
          updated: '2025-06-06 14:04:33.754Z',
        },
      ],
      location: 'Da',
      miniapp_name: 'Xrystik',
      user_info: 'Yep ',
      user_photo: 'avatar_ke580ghznp.jpg',
      widgets: [
        {
          collectionId: '9fxx39w3wb6ulbf',
          collectionName: 'widgets',
          created: '2025-06-10 12:24:48.908Z',
          data: {
            text: 'Rhehhdgdgdgdvdvdvfvfvfvfv',
            type: 'text',
          },
          files: [],
          id: '2nq3yh2427m23if',
          order: -1,
          updated: '2025-06-10 12:24:48.908Z',
          user: 'xb25g6o0oy91395',
        },
        {
          collectionId: '9fxx39w3wb6ulbf',
          collectionName: 'widgets',
          created: '2025-06-10 12:24:48.908Z',
          data: {
            text: 'RhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfvRhehhdgdgdgdvdvdvfvfvfvfv',
            type: 'text',
          },
          files: [],
          id: '2nq3yh242qm23if',
          order: -1,
          updated: '2025-06-10 12:24:48.908Z',
          user: 'xb25g6o0oy91395',
        },
      ],
    },
    {
      age: 59,
      gender: 'male',
      id: 'be8v14s4o785bz1',
      interests: [
        {
          collectionId: 'pbc_2896159602',
          collectionName: 'interests',
          created: '2025-06-06 14:00:35.403Z',
          id: 'me9j99rc187n1a2',
          tag: 'Футбол',
          updated: '2025-06-06 14:00:35.403Z',
        },
        {
          collectionId: 'pbc_2896159602',
          collectionName: 'interests',
          created: '2025-06-06 14:02:26.340Z',
          id: '2681n9u2d0tr6kj',
          tag: 'История',
          updated: '2025-06-06 14:02:26.340Z',
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
      miniapp_name: '123',
      user_info: 'Я люблю рисовать ',
      user_photo: 'avatar_8nmvyge0ie.jpg',
      widgets: [
        {
          collectionId: '9fxx39w3wb6ulbf',
          collectionName: 'widgets',
          created: '2025-06-10 12:23:03.315Z',
          data: {
            link: 'IIUSUGlobal/939',
            type: 'post',
          },
          files: [],
          id: 'gk2036h17dci770',
          order: -1,
          updated: '2025-06-10 12:23:03.315Z',
          user: 'be8v14s4o785bz1',
        },
      ],
    },
    {
      age: 19,
      gender: 'male',
      id: '3068b3z9347v81y',
      interests: [
        {
          collectionId: 'pbc_2896159602',
          collectionName: 'interests',
          created: '2025-06-06 14:03:38.168Z',
          id: '0ywgm9771phr98r',
          tag: 'Сборка компьютеров',
          updated: '2025-06-06 14:03:38.168Z',
        },
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
          created: '2025-06-06 14:04:05.409Z',
          id: '54nv8g5meu8wdbh',
          tag: 'Машинное обучение',
          updated: '2025-06-06 14:04:05.409Z',
        },
      ],
      location: 'Екатеринбург',
      miniapp_name: 'Ваня',
      user_info: 'Программирую, играю на гитаре',
      user_photo: 'avatar_cd4xulpzk1.jpg',
      widgets: [
        {
          additionalData: {
            hours_played: 12,
            icon: 'https://media.steampowered.com/steamcommunity/public/images/apps/641990/c1d2bb67884d14088b7befc3d442df6eb801cd13.jpg',
            title: 'The Escapists 2',
            type: 'steam_game',
          },
          collectionId: '9fxx39w3wb6ulbf',
          collectionName: 'widgets',
          created: '2025-06-10 13:30:36.394Z',
          data: {
            accountLink: 'https://steamcommunity.com/id/lowderplay/',
            gameId: '641990',
            type: 'steam_game',
          },
          files: [],
          id: 'gq1jkeew96jn6sy',
          order: 0,
          updated: '2025-06-10 13:30:36.394Z',
          user: '3068b3z9347v81y',
        },
        {
          additionalData: {
            type: 'photo',
            urls: [
              '9fxx39w3wb6ulbf/4d7n809kk119t26/g_b8_te3_rmy5_a_sycfrtd9wv.jpg',
            ],
          },
          collectionId: '9fxx39w3wb6ulbf',
          collectionName: 'widgets',
          created: '2025-06-10 13:31:13.190Z',
          data: {
            type: 'photo',
          },
          files: ['g_b8_te3_rmy5_a_sycfrtd9wv.jpg'],
          id: '4d7n809kk119t26',
          order: 1,
          updated: '2025-06-10 13:31:13.190Z',
          user: '3068b3z9347v81y',
        },
        {
          additionalData: {
            hours_played: 17,
            icon: 'https://media.steampowered.com/steamcommunity/public/images/apps/620980/d2922f271fca9579ecdcc07408cee13d87ca4148.jpg',
            title: 'Beat Saber',
            type: 'steam_game',
          },
          collectionId: '9fxx39w3wb6ulbf',
          collectionName: 'widgets',
          created: '2025-06-10 13:51:12.570Z',
          data: {
            accountLink: 'https://steamcommunity.com/id/lowderplay/',
            gameId: '620980',
            type: 'steam_game',
          },
          files: [],
          id: '3ux7a9k1j66uhnb',
          order: 2,
          updated: '2025-06-10 13:51:12.570Z',
          user: '3068b3z9347v81y',
        },
        {
          collectionId: '9fxx39w3wb6ulbf',
          collectionName: 'widgets',
          created: '2025-06-10 13:17:11.856Z',
          data: {
            text: 'asd',
            type: 'text',
          },
          files: [],
          id: 'vg9efev6h5294qx',
          order: 3,
          updated: '2025-06-10 13:30:14.570Z',
          user: '3068b3z9347v81y',
        },
      ],
    },
  ];
  let firstProfile = userProfiles[0];
  let secondProfile = userProfiles[1];
  let thirdProfile = userProfiles[2];
  let currentProfile = $state(0);
  let isMousePress = $state(false);
  let container: HTMLDivElement = $state();
  let profileRef: HTMLDivElement[] = $state([]);
  let profiles: PageData[] = $state([firstProfile, thirdProfile]);
  useTelegramButton(() => goto('/Profile'));

  let testProfiles = ['A', 'B', 'C', 'D'];

  function updateCurrentSection() {
    // const scrollLeft = container.scrollTop;
    // const sectionWidth = container.offsetHeight;
    // console.log(scrollLeft, sectionWidth, container);
    // console.log(document.body.offsetHeight);
    // console.log(
    //   container.scrollTop,
    //   profileRef[0].offsetHeight,
    //   profileRef[1].offsetHeight,
    // );
    if (isMousePress) return;
    console.log(
      `${container.scrollTop + document.body.offsetHeight - 160} / ${
        profileRef[0].offsetHeight + document.body.offsetHeight / 2
      } = ${(container.scrollTop + document.body.offsetHeight - 160) / (profileRef[0].offsetHeight + document.body.offsetHeight / 2)}`,
    );
    currentProfile = Math.trunc(
      (container.scrollTop + document.body.offsetHeight - 160) /
        (profileRef[0].offsetHeight + document.body.offsetHeight / 2),
    );
  }

  onMount(() => {
    window.Telegram.WebApp.MainButton.setText(
      // changeMode ? 'Сохранить' : 'Изменить виджеты',
      'Моя анкета',
    );
  });
  $inspect(isMousePress);
</script>

<!--<div class="bg-green-500 w-full h-screen">123</div>-->
<!--<div class="w-full">-->
<!--  <div class="w-full h-[600px] bg-red-600">1</div>-->
<!--  <p>{scrollY.current}</p>-->
<!--  <div class="w-full h-[600px] bg-green-600">2</div>-->
<!--</div>-->

<div
  class="snap-y snap-mandatory w-full h-full overflow-y-auto space-y-40"
  bind:this={container}
  onscroll={updateCurrentSection}
  onmousedown={() => (isMousePress = true)}
  onmouseup={() => (isMousePress = false)}
>
  <!--  <button-->
  <!--    onclick={() => {-->
  <!--      profiles.splice(0, 1);-->
  <!--      profiles.push(thirdProfile);-->
  <!--      console.log(profiles);-->
  <!--    }}-->
  <!--    class="fixed size-25 bg-purple-500">123</button-->
  <!--  >-->
  <button
    onclick={() => {
      profiles.push(userProfiles[0]);
      profiles.splice(0, 1);
      container.scrollTop = 0;
      console.log(profiles);
      // profileRef.forEach((element: HTMLDivElement) =>
      //   console.log(element.offsetHeight),
      // );
    }}
    class="fixed right-0 size-25 bg-purple-500 z-100 flex">Check state</button
  >
  <div class="fixed size-25 bg-purple-500">
    Текущий профиль: {currentProfile}
  </div>

  {#each profiles as profile, i}
    <div class="snap-start" bind:this={profileRef[i]}>
      <Questionnaire data={profile} />
    </div>
  {/each}

  <!--  <div class="snap-start">-->
  <!--    <Questionnaire data={secondProfile} />-->
  <!--  </div>-->
  <!--  <div class="snap-start">-->
  <!--    <Questionnaire data={firstProfile} />-->
  <!--  </div>-->
  <!--  <div class="snap-start">-->
  <!--    <Questionnaire data={firstProfile} />-->
  <!--  </div>-->
  <!--  <div-->
  <!--    class="snap-start bg-purple-600 w-full h-300 flex items-center justify-center text-8xl"-->
  <!--  >-->
  <!--    1-->
  <!--  </div>-->
  <!--  <div-->
  <!--    class="snap-start bg-purple-600 w-full h-300 flex items-center justify-center text-8xl"-->
  <!--  >-->
  <!--    2-->
  <!--  </div>-->
  <!--  <div-->
  <!--    class="snap-start bg-purple-600 w-full h-300 flex items-center justify-center text-8xl"-->
  <!--  >-->
  <!--    3-->
  <!--  </div>-->
  <!--  <div-->
  <!--    class="snap-start bg-purple-600 w-full h-300 flex items-center justify-center text-8xl"-->
  <!--  >-->
  <!--    4-->
  <!--  </div>-->
</div>

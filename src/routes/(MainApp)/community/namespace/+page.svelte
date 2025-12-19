<script lang="ts">
  import { ChevronRight, PencilLine } from '@lucide/svelte';
  import UserNamespaces from '$lib/components/namespaces/UserNamespaces.svelte';
  import NamespaceMembersLst from '$lib/components/namespaces/NamespaceMembersLst.svelte';
  import type { NamespaceMembers } from '$lib/namespaceTypes';
  import { getCorrectForm } from '$lib/utils';
  import client from '$lib/api/client';
  import { goto } from '$app/navigation';

  let testNameSpaceses = {
    img: 'https://www.soyuz.ru/public/uploads/files/2/7480281/20220315190534af66e2c5d3.jpg',
    title: 'Hello Lowder',
    amount_members: 256,
    is_admin: true,
  };

  let namespaceUsers: NamespaceMembers = [
    {
      img: 'https://www.soyuz.ru/public/uploads/files/2/7480281/20220315190534af66e2c5d3.jpg',
      username: 'Valera',
      is_admin: true,
    },
    {
      img: 'https://www.soyuz.ru/public/uploads/files/2/7480281/20220315190534af66e2c5d3.jpg',
      username: 'Valera1',
      is_admin: false,
    },
    {
      img: 'https://www.soyuz.ru/public/uploads/files/2/7480281/20220315190534af66e2c5d3.jpg',
      username: 'Valera2',
      is_admin: false,
    },
    {
      img: 'https://www.soyuz.ru/public/uploads/files/2/7480281/20220315190534af66e2c5d3.jpg',
      username: 'Valera3',
      is_admin: false,
    },
  ];
</script>

<div class="flex flex-col justify-center items-center relative z-1 mt-4 mb-4">
  <img
    alt="Картинка неймспейса"
    src={testNameSpaceses.img}
    class="size-24 rounded-full object-cover"
  />
  <div class="mt-2.5">
    <p class="font-medium">{testNameSpaceses.title}</p>
    <p class="text-sm text-inactive">
      {testNameSpaceses.amount_members}
      {getCorrectForm(testNameSpaceses.amount_members, [
        'участник',
        'участника',
        'участников',
      ])}
    </p>
  </div>
</div>
<img
  alt="Картинка неймспейса"
  src={testNameSpaceses.img}
  class="size-100 rounded-full object-cover absolute -top-[35%] z-0 opacity-60 blur-2xl"
/>
<div class="px-4 flex flex-col justify-start items-start w-full">
  <button
    class="border-text-color/25 border-x-2 border-t-2 rounded-t-xl px-4 py-3
    w-full flex justify-between"
    onclick={() => goto('/view')}
  >
    <p>Смотреть анкеты</p>
    <ChevronRight class="size-5" />
  </button>
  <div
    class="border-text-color/25 border-x-2 border-b-2 rounded-b-xl border-t-2 rounded-b-2 px-4 py-3 w-full flex justify-between mb-4"
  >
    <p>Редактировать</p>
    <button>
      <PencilLine class="size-5" />
    </button>
  </div>
  <p>Участники</p>
  <NamespaceMembersLst users={namespaceUsers} />
</div>

<script lang="ts">
  import {
    ChevronRight,
    PencilLine,
    Share2,
    Trash2,
    LogOut,
  } from '@lucide/svelte';
  import UserNamespaces from '$lib/components/namespaces/UserNamespaces.svelte';
  import NamespaceMembersLst from '$lib/components/namespaces/NamespaceMembersLst.svelte';
  import { getCorrectForm } from '$lib/utils';
  import client from '$lib/api/client';
  import { goto } from '$app/navigation';
  import { createQuery } from '@tanstack/svelte-query';
  import { page } from '$app/state';
  import EditGroup from '$lib/components/namespaces/EditGroup.svelte';
  import { db } from '$lib';
  import CreateGroupBtn from '$lib/components/namespaces/CreateGroupBtn.svelte';
  import { userData } from '$lib/storage/userData.svelte';
  import { toast } from 'svelte-sonner';

  let testNameSpaceses = {
    img: 'https://www.soyuz.ru/public/uploads/files/2/7480281/20220315190534af66e2c5d3.jpg',
    title: 'Hello Lowder',
    amount_members: 256,
    is_admin: true,
  };

  let namespaceUsers = [
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
  let isEditOpen = $state(false);
  let isShareOpen = $state(false);
  const namespaceId = page.url.pathname.split('/').pop();
  const namespaceInfo = createQuery(() => ({
    queryKey: ['namespaceInfo'],
    queryFn: async () =>
      await client.GET('/api/namespace/{namespace_id}/pages', {
        params: {
          path: {
            namespace_id: namespaceId,
          },
        },
      }),
    select: (data) => data.data,
    gcTime: 0,
    staleTime: 0,
  }));
</script>

{#if namespaceInfo.isSuccess}
  <div class="w-full overflow-auto">
    <div
      class="flex flex-col justify-center items-center relative z-1 mt-4 mb-4"
    >
      <img
        alt="Картинка неймспейса"
        src={`${db}/api/files/namespaces/${namespaceInfo.data?.namespace?.id}/undefined?t=${Date.now()}
`}
        class="size-24 rounded-full object-cover"
      />
      <div class="mt-2.5 text-center">
        <p class="font-medium">{namespaceInfo.data?.namespace?.title}</p>
        <p class="text-sm text-inactive">
          {namespaceInfo.data?.namespace?.members_count}
          {getCorrectForm(namespaceInfo.data?.namespace?.members_count, [
            'участник',
            'участника',
            'участников',
          ])}
        </p>
        <p class="max-w-[300px] line-clamp-3 text-sm break-words">
          {namespaceInfo.data?.namespace?.description}
        </p>
      </div>
    </div>
    <img
      alt="Картинка неймспейса"
      src={`${db}/api/files/namespaces/${namespaceInfo.data?.namespace?.id}/undefined?t=${Date.now()}
`}
      class="size-100 rounded-full object-cover absolute -top-[35%] z-0 opacity-60 blur-2xl"
    />
    <div class="px-4 flex flex-col justify-start items-start w-full">
      <button
        class="border-text-color/25 border-2 px-4 py-3
    w-full flex justify-between rounded-t-xl"
        onclick={() =>
          goto(`/view?namespaceId=${namespaceInfo.data?.namespace?.id}`)}
      >
        <p>Смотреть анкеты</p>
        <ChevronRight class="size-5" />
      </button>

      <button
        class="border-text-color/25 border-x-2 px-4 py-3
    w-full flex justify-between"
        onclick={() => (isShareOpen = !isShareOpen)}
      >
        <p>Поделиться</p>
        <Share2 class="size-5" />
      </button>
      {#if namespaceInfo.data?.namespace?.admin === userData.current.tg_user}
        <button
          class="border-text-color/25 border-x-2 border-t-2 px-4 py-3 w-full flex justify-between"
          onclick={() => (isEditOpen = !isEditOpen)}
        >
          <p>Редактировать</p>
          <PencilLine class="size-5" />
        </button>
      {/if}
      <button
        class="border-text-color/25 border-2 rounded-b-xl px-4 py-3
    w-full flex justify-between mb-4"
        onclick={async () => {
          if (
            namespaceInfo.data?.namespace?.admin === userData.current?.tg_user
          ) {
            console.log('Выполняем удаление неймспейса');
            await client.DELETE('/api/admin/{namespace_id}', {
              params: {
                path: {
                  namespace_id: namespaceInfo.data?.namespace?.id,
                },
              },
            });
          } else {
            console.log('Выходим из неймспейса');
            await client.DELETE('/api/namespace/{namespace_id}', {
              params: {
                path: {
                  namespace_id: namespaceInfo.data?.namespace?.id,
                },
              },
            });
          }
          await goto('/community');
        }}
      >
        {#if namespaceInfo.data?.namespace?.admin === userData.current?.tg_user}
          <p>Удалить</p>
          <Trash2 class="size-5" />
        {:else}
          <p>Покинуть</p>
          <LogOut class="size-5" />
        {/if}
      </button>
      <p>Участники</p>
      <NamespaceMembersLst
        users={namespaceInfo.data?.users}
        adminId={namespaceInfo.data?.namespace?.admin}
        namespaceId={namespaceInfo.data?.namespace?.id}
      />
    </div>
    <EditGroup
      bind:open={isEditOpen}
      namespaceData={namespaceInfo.data?.namespace}
      onSave={() => setTimeout(() => namespaceInfo.refetch(), 200)}
    />
    <div class="hidden">
      <CreateGroupBtn
        bind:isSheetOpen={isShareOpen}
        currentStage={2}
        data={{
          id: namespaceId,
          title: namespaceInfo.data?.namespace?.title,
        }}
      />
    </div>
  </div>
{/if}

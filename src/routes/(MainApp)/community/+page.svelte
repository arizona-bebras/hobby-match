<script lang="ts">
  import { Input } from '$lib/components/ui/input/index.js';
  import UserNamespaces from '$lib/components/namespaces/UserNamespaces.svelte';
  import CreateGroupBtn from '$lib/components/namespaces/CreateGroupBtn.svelte';
  import { db } from '$lib';
  import client from '$lib/api/client';
  import { onMount } from 'svelte';
  import { createQuery } from '@tanstack/svelte-query';
  import { accessToken } from '$lib/storage/accessToken.svelte';

  const namespaces = createQuery(() => ({
    queryKey: ['namespaces'],
    queryFn: async () => await client.GET('/api/me/namespaces'),
    select: (data) => data.data,
  }));

  let testNameSpaceses = [
    {
      img: '',
      title: 'Hello Lowder',
      amount_members: 256,
      is_admin: true,
    },
    {
      img: '',
      title: 'Hello Lowders',
      amount_members: 2,
      is_admin: false,
    },
    {
      img: '',
      title: 'Yep',
      amount_members: 201,
      is_admin: false,
    },
    {
      img: '',
      title: 'Sup',
      amount_members: 201,
      is_admin: false,
    },
    {
      img: '',
      title: 'Hello Lowder1',
      amount_members: 256,
      is_admin: true,
    },
    {
      img: '',
      title: 'Hello Lowders1',
      amount_members: 201,
      is_admin: false,
    },
    {
      img: '',
      title: 'Yep1',
      amount_members: 201,
      is_admin: false,
    },
    {
      img: '',
      title: 'Sup1',
      amount_members: 201,
      is_admin: false,
    },
    {
      img: '',
      title: 'Sup2',
      amount_members: 201,
      is_admin: false,
    },
    {
      img: '',
      title: 'Sup3',
      amount_members: 201,
      is_admin: false,
    },
    {
      img: '',
      title: 'Sup4',
      amount_members: 201,
      is_admin: false,
    },
  ];

  let searchFilter = $state('');
  $inspect(searchFilter);
</script>

<!--<img src={`data:image/png;base64,${query.data.data.files[0]}`} />-->
<div class="p-4 w-full text-text-color">
  <div class="flex justify-between font-medium items-center mb-6">
    <p class="">Мои неймспейсы</p>
    <CreateGroupBtn />
  </div>
  <Input
    placeholder="Введи название неймспейса..."
    class="mb-6"
    bind:value={searchFilter}
  />
  <UserNamespaces namespacesList={namespaces.data?.namespaces} {searchFilter} />
</div>

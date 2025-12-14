<script lang="ts">
  import { Input } from '$lib/components/ui/input/index.js';
  import { type UserNamespaces } from '$lib/namespaceTypes';
  import CreateGroupBtn from '$lib/components/namespaces/CreateGroupBtn.svelte';

  let testNameSpaceses: UserNamespaces = [
    {
      img: '',
      title: 'Hello Lowder',
      amount_members: 256,
      is_admin: true,
    },
    {
      img: '',
      title: 'Hello Lowders',
      amount_members: 201,
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
  <UserNamespaces namespacesList={testNameSpaceses} {searchFilter} />
</div>

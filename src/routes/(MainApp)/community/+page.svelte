<script lang="ts">
  import { Input } from '$lib/components/ui/input/index.js';
  import UserNamespaces from '$lib/components/namespaces/UserNamespaces.svelte';
  import CreateGroupBtn from '$lib/components/namespaces/CreateGroupBtn.svelte';

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
      amount_members: 201,
      is_admin: false,
    },
  ];
</script>

<div class="p-4 w-full text-text-color">
  <div class="flex justify-between font-medium items-center mb-6">
    <p class="">Мои неймспейсы</p>
    <CreateGroupBtn />
  </div>
  <Input placeholder="Введи название неймспейса..." class="mb-6" />
  <UserNamespaces namespacesList={testNameSpaceses} />
</div>

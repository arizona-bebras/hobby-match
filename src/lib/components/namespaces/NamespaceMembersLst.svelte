<script lang="ts">
  import { ShieldUser, Shield } from '@lucide/svelte';
  import { goto } from '$app/navigation';
  import UserInfoBtn from '$lib/components/namespaces/UserInfoBtn.svelte';
  import { getImage } from '$lib/utils';
  import { IsInViewport } from 'runed';
  import MemberRow from '$lib/components/namespaces/MemberRow.svelte';

  type NamespaceMembers = {
    tg_user: number;
    miniapp_name: string;
  };

  const {
    users,
    adminId,
    namespaceId,
  }: { users: NamespaceMembers[]; adminId: number; namespaceId: string } =
    $props();
</script>

<div class="w-full text-text-color">
  {#each users as user, i (user.tg_user)}
    <MemberRow {user} {i} {adminId} {namespaceId} />
  {/each}
</div>

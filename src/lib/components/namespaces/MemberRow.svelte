<script lang="ts">
  import { getImage } from '$lib/utils.ts';
  import UserInfoBtn from '$lib/components/namespaces/UserInfoBtn.svelte';
  import { ShieldUser, ArrowRight } from '@lucide/svelte';
  import { IsInViewport } from 'runed';
  import { db } from '$lib';

  let targetNode = $state<HTMLElement>();
  const inViewport = new IsInViewport(() => targetNode, {
    once: true,
  });
  type NamespaceMembers = {
    tg_user: number;
    miniapp_name: string;
  };

  const {
    user,
    adminId,
    i,
    namespaceId,
  }: {
    user: NamespaceMembers;
    adminId: number;
    i: number;
    namespaceId: number;
  } = $props();

  let isSheetOpen = $state(false);
</script>

<div
  class="flex py-1.5 items-center justify-between border-b-[2px] border-text-color/25"
  bind:this={targetNode}
>
  <button
    class="flex items-center justify-between w-full select-none"
    onclick={() => (isSheetOpen = !isSheetOpen)}
  >
    <div class="flex gap-2.5 items-center">
      <img
        src={`${db}/api/files/users/${user.tg_user}/undefined
`}
        class="size-8 object-cover rounded-[8px]"
        alt="group image"
      />
      <div class="flex flex-col">
        <div class="font-medium items-center gap-2">
          <div class="flex items-center gap-2">
            {user.miniapp_name}
            {#if user.tg_user === adminId}
              <div
                class="size-4 bg-accent rounded-[4px] flex flex-col items-center justify-center"
              >
                <ShieldUser class="size-3.5 " />
              </div>
            {/if}
          </div>
          {#if user.tg_user === adminId}
            <p>АДМИНИСТРАТОР</p>
          {/if}
        </div>
      </div>
    </div>
    <ArrowRight class="size-5" />
    <UserInfoBtn {user} bind:open={isSheetOpen} {namespaceId} {adminId} } />
  </button>
</div>

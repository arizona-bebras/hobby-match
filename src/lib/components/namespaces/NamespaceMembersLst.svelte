<script lang="ts">
  import { ShieldUser, Shield } from '@lucide/svelte';
  import { goto } from '$app/navigation';
  import UserInfoBtn from '$lib/components/namespaces/UserInfoBtn.svelte';

  const { users }: { users: object[] } = $props();
</script>

<div class="overflow-y-auto h-[calc(100vh-390px)] w-full">
  {#each users as user (user.username)}
    <div
      class="flex py-1.5 items-center justify-between border-b-[2px] border-text-color/25"
    >
      <div class="flex gap-2.5 items-center">
        <img
          src="https://www.soyuz.ru/public/uploads/files/2/7480281/20220315190534af66e2c5d3.jpg"
          class="size-8 object-cover"
          alt="group image"
        />
        <div class="flex flex-col">
          <div class="font-medium items-center gap-2">
            <div class="flex items-center gap-2">
              {user.username}
              {#if user.is_admin}
                <div
                  class="size-4 bg-accent rounded-[4px] flex flex-col items-center justify-center"
                >
                  <ShieldUser class="size-3.5 " />
                </div>
              {/if}
            </div>
            {#if user.is_admin}
              <p>АДМИНИСТРАТОР</p>
            {/if}
          </div>
        </div>
      </div>
      <UserInfoBtn />
    </div>
  {/each}
</div>

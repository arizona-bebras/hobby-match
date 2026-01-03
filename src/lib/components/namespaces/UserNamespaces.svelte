<script lang="ts">
  import { ArrowRight, ShieldUser, Shield } from '@lucide/svelte';
  import { goto } from '$app/navigation';
  import { getCorrectForm } from '$lib/utils';
  import { userData } from '$lib/storage/userData.svelte';
  import client from '$lib/api/client';
  import { onMount } from 'svelte';
  import { getImage } from '$lib/utils';
  import { createQuery } from '@tanstack/svelte-query';
  import { db } from '$lib';

  type namespaces = {
    admin: string;
    description: string;
    id: string;
    picture: number[];
    title: string;
    members_count: number;
  }[];
  const {
    namespacesList,
    searchFilter,
  }: { namespacesList: namespaces; searchFilter: string } = $props();
</script>

<div class="overflow-y-auto h-[calc(100vh-200px)]">
  {#each namespacesList as namespace (namespace.id)}
    {#if namespace.title.toLowerCase().includes(searchFilter.toLowerCase())}
      <button
        onclick={() => goto(`/community/${namespace.id}`)}
        class="flex py-1.5 items-center justify-between border-b-[2px] border-text-color/25 w-full"
      >
        <div class="flex gap-2.5 items-center text-justify">
          <img
            src={`${db}/api/files/namespaces/${namespace.id}/undefined
`}
            class="size-8 object-cover rounded-[8px]"
            alt="group image"
          />
          <div class="flex flex-col">
            <div class="font-medium flex items-center gap-2">
              {namespace.title}
              {#if userData.current?.tg_user === namespace.admin.toString()}
                <div
                  class="size-4 bg-accent rounded-[4px] flex items-center justify-center"
                >
                  <ShieldUser class="size-3.5 " />
                </div>
              {/if}
            </div>
            <p>
              {namespace.members_count}
              {getCorrectForm(namespace.members_count, [
                'участник',
                'участника',
                'участников',
              ])}
            </p>
          </div>
        </div>
        <ArrowRight class="size-5" />
      </button>
    {/if}
  {/each}
</div>

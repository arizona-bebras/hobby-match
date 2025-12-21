<script lang="ts">
  import { ArrowRight, ShieldUser, Shield } from '@lucide/svelte';
  import { goto } from '$app/navigation';
  import { getCorrectForm } from '$lib/utils';

  type namespaces = {
    admin: string;
    description: string;
    id: string;
    picture: number[];
    title: string;
  }[];
  const {
    namespacesList,
    searchFilter,
  }: { namespacesList: namespaces; searchFilter: string } = $props();
</script>

<div class="overflow-y-auto h-[calc(100vh-200px)]">
  {#each namespacesList as namespace (namespace.title)}
    {#if namespace.title.toLowerCase().includes(searchFilter.toLowerCase())}
      <div
        class="flex py-1.5 items-center justify-between border-b-[2px] border-text-color/25"
      >
        <div class="flex gap-2.5 items-center">
          <img
            src="https://www.soyuz.ru/public/uploads/files/2/7480281/20220315190534af66e2c5d3.jpg"
            class="size-8 object-cover rounded-[8px]"
            alt="group image"
          />
          <div class="flex flex-col">
            <div class="font-medium flex items-center gap-2">
              {namespace.title}
              {#if namespace.is_admin}
                <div
                  class="size-4 bg-accent rounded-[4px] flex items-center justify-center"
                >
                  <ShieldUser class="size-3.5 " />
                </div>
              {/if}
            </div>
            <p>
              {namespace.amount_members}
              {getCorrectForm(namespace.amount_members, [
                'участник',
                'участника',
                'участников',
              ])}
            </p>
          </div>
        </div>
        <button onclick={() => goto(`/community/${namespace.id}}`)}>
          <ArrowRight class="size-5" />
        </button>
      </div>
    {/if}
  {/each}
</div>

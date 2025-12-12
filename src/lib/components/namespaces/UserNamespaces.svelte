<script lang="ts">
  import { ArrowRight, ShieldUser, Shield } from '@lucide/svelte';
  import { goto } from '$app/navigation';

  const {
    namespacesList,
    searchFilter,
  }: { namespacesList: object[]; searchFilter: string } = $props();

  function getCorrectForm(count: number): string {
    const lastDigit = count % 10;
    const lastTwoDigits = count % 100;
    let form;
    if (lastTwoDigits >= 11 && lastTwoDigits <= 19) {
      form = 'участников';
    } else {
      if (lastDigit === 1) {
        form = 'участник';
      } else if (lastDigit >= 2 && lastDigit <= 4) {
        form = 'участника';
      } else {
        form = 'участников';
      }
    }
    return form;
  }
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
            class="size-8"
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
              {getCorrectForm(namespace.amount_members)}
            </p>
          </div>
        </div>
        <button onclick={() => goto('/community/namespace')}>
          <ArrowRight class="size-5" />
        </button>
      </div>
    {/if}
  {/each}
</div>

<script lang="ts">
  import { LoaderCircle } from '@lucide/svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import { db } from '$lib';
  import { userData } from '$lib/storage/userData.svelte';

  function close(node: HTMLDivElement, { duration }: { duration: number }) {
    const startWidth = node.offsetWidth;
    const startHeight = node.offsetHeight;

    return {
      duration,
      css: (t: number) => {
        return `
          width: ${startWidth * t}px; 
          height: ${startHeight * t}px; 
          opacity: ${t};
          overflow: hidden;
        `;
      },
    };
  }

  const {
    elementSize = $bindable(),
    offeredProfiles,
  }: { elementSize: number; offeredProfiles: PageData[] } = $props();
</script>

<div
  out:close={{ duration: 200 }}
  class="bg-accent/25 max-w-15 max-h-27 rounded-full mx-auto overflow-hidden"
  style:width="{elementSize}px"
  style:height="{elementSize}px"
>
  <p style:font-size="min(30px, {elementSize}px)" class="text-center">
    &#8593;
  </p>
  {#if elementSize >= 35}
    {#if offeredProfiles.length > 1}
      <img
        class="rounded-full p-1 aspect-square object-cover max-w-15 max-h-15 mx-auto"
        style:width="{elementSize - 35}px"
        style:height="{elementSize - 35}px"
        src={`${db}/api/files/users/${userData.current?.tg_user}/undefined
`}
        alt="userImage"
      />
    {:else}
      <LoaderCircle class="animate-spin w-full" />
    {/if}
  {/if}
</div>

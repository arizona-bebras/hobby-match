<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import {
    blur,
    crossfade,
    draw,
    fade,
    fly,
    scale,
    slide,
  } from 'svelte/transition';
  import { onMount } from 'svelte';

  let { open = $bindable(), onClose } = $props();
  $inspect(open);

  onMount(() => console.log(123));
</script>

<div transition:fade={{ duration: 2000 }}>
  <Sheet.Root bind:open onOpenChange={(open) => onClose()}>
    <Sheet.Content>
      <Sheet.Header>
        <Sheet.Title>Are you sure absolutely sure?</Sheet.Title>
        <Sheet.Description>
          This action cannot be undone. This will permanently delete your
          account and remove your data from our servers.
        </Sheet.Description>
      </Sheet.Header>
    </Sheet.Content>
  </Sheet.Root>
</div>

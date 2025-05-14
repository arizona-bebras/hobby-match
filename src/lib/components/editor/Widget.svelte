<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  let { open } = $props();
</script>

<Sheet.Root bind:open>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <Sheet.Title>Are you sure absolutely sure?</Sheet.Title>
      <Sheet.Description>
        This action cannot be undone. This will permanently delete your account
        and remove your data from our servers.
      </Sheet.Description>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>

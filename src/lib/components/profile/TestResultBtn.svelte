<script lang="ts">
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import Similarity from '$lib/components/profile/Similarity.svelte';
  let { userSimilarity }: { userSimilarity: number } = $props();
</script>

<Sheet.Root>
  <Sheet.Trigger class="w-full"><Similarity {userSimilarity} /></Sheet.Trigger>
  <Sheet.Content side="bottom">
    <Sheet.Header>
      <Sheet.Title>Результаты личностного теста</Sheet.Title>
      <Sheet.Description>
        This action cannot be undone. This will permanently delete your account
        and remove your data from our servers.
      </Sheet.Description>
    </Sheet.Header>
  </Sheet.Content>
</Sheet.Root>

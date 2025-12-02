<script lang="ts">
  import { deleteWidget } from '$lib/components/widgetConstructors/widgetsConstructor.js';
  import { Trash2 } from '@lucide/svelte';
  import { invalidate } from '$app/navigation';
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  let { widgetId }: { widgetId: string } = $props();
</script>

<Sheet.Close
  onclick={async () => {
    await deleteWidget(widgetId.toString());
    await invalidate('user:widgets');
  }}
  class="ring-offset-background focus:ring-ring data-[state=open]:bg-secondary absolute right-4 top-2 rounded-sm opacity-70 transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-offset-2 disabled:pointer-events-none p-2"
>
  <Trash2 class="size-5 text-destructive" />
  <span class="sr-only">Close</span>
</Sheet.Close>

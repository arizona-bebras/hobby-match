<script lang="ts">
  import Header from '$lib/components/profile/Header.svelte';
  import UserInfo from '$lib/components/profile/UserInfo.svelte';
  import RenderWidget from '$lib/components/profile/RenderWidget.svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';
  import NamespaceHeader from '$lib/components/profile/NamespaceHeader.svelte';
  import { ScrollState } from 'runed';

  let {
    data,
    isNamespaceProfile = false,
    scroll,
  }: {
    data: PageData;
    isNamespaceProfile?: boolean;
    scroll: ScrollState;
  } = $props();
</script>

<Header {data} />
<div class="font-[Inter] px-4 w-full max-w-full relative">
  <UserInfo {data} {isNamespaceProfile} {scroll} changeMode={false} />
  {#each data.widgets as widget (widget.id)}
    <div class="relative mb-2">
      <RenderWidget {widget} isViewingMode={true} />
    </div>
  {/each}
</div>

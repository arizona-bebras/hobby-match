<script lang="ts">
  import { Plus } from '@lucide/svelte';
  import EditTgPost from '$lib/components/editor/tg-post/EditTgPost.svelte';
  import EditSurvey from '$lib/components/editor/survey/EditSurvey.svelte';
  import EditText from '$lib/components/editor/text/EditText.svelte';
  import EditVideo from '$lib/components/editor/video/EditVideo.svelte';
  import type { WidgetType } from '$lib/widgetTypes/widgetTypes';
  import Edit from '$lib/components/editor/Edit.svelte';
  import WidgetsListMenu from '$lib/components/editor/WidgetsListMenu.svelte';
  import EditImage from '$lib/components/editor/image/EditImage.svelte';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  import EditSocial from '$lib/components/editor/social/EditSocial.svelte';
  import EditToDo from '$lib/components/editor/todo/EditToDo.svelte';
  import EditProgress from '$lib/components/editor/progress/EditProgress.svelte';
  import type { WidgetWithService } from '$lib/components/widgetConstructors/widgetsConstructor';
  import EditAudio from '$lib/components/editor/audio/EditAudio.svelte';
  import EditSteamGame from '$lib/components/editor/steam-game/EditSteamGame.svelte';
  import Header from '$lib/components/profile/Header.svelte';
  import UserInfo from '$lib/components/profile/UserInfo.svelte';
  import RenderWidget from '$lib/components/profile/RenderWidget.svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';

  let { data }: { data: PageData } = $props();
  console.log(data);
</script>

<Header {data} />
<div class="font-[Inter] px-4 w-full max-w-full relative">
  <UserInfo {data} />

  {#each data.widgets as widget (widget.id)}
    <div class="relative mb-2">
      <RenderWidget {widget} isViewingMode={true} />
    </div>
  {/each}
</div>

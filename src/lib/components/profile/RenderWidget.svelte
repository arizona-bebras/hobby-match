<script lang="ts">
  import TextWidget from '$lib/components/widgets/TextWidget.svelte';
  import AudioWidget from '$lib/components/widgets/AudioWidget.svelte';
  import GameWidget from '$lib/components/widgets/GameWidget.svelte';
  import VideoWidget from '$lib/components/widgets/VideoWidget.svelte';
  import SocialWidget from '$lib/components/widgets/SocialWidget.svelte';
  import ProgressWidget from '$lib/components/widgets/ProgressWidget.svelte';
  import ToDoWidget from '$lib/components/widgets/ToDoWidget.svelte';
  import SurveyWidget from '$lib/components/widgets/SurveyWidget.svelte';
  import PhotoWidget from '$lib/components/widgets/PhotoWidget.svelte';
  import TgPostWidget from '$lib/components/widgets/TgPostWidget.svelte';
  import ActivityWidget from '$lib/components/widgets/ActivityWidget.svelte';
  import type { Widget } from '$lib/widgetTypes/widgetTypes';

  let { widget, isViewingMode }: { widget: Widget; isViewingMode: boolean } =
    $props();
</script>

{#if widget.data.type === 'text'}
  {console.log('TRYING RENDER TEXT')}
  <TextWidget data={widget.data} />
{:else if widget.data.type === 'audio'}
  <AudioWidget data={widget.data} />
{:else if widget.data.type === 'steam_game'}
  <GameWidget data={widget.additionalData} />
{:else if widget.data.type === 'video'}
  <VideoWidget data={widget.data} />
{:else if widget.data.type === 'social_media' && widget.additionalData?.type !== 'photo' && widget.additionalData?.type !== 'survey'}
  <SocialWidget data={widget.data} socialMediaData={widget.additionalData} />
{:else if widget.data.type === 'progress_bar'}
  <ProgressWidget data={widget.data} />
{:else if widget.data.type === 'todo'}
  <ToDoWidget data={widget.data} widgetId={widget.id} {isViewingMode} />
{:else if widget.data.type === 'survey' && widget.additionalData?.type === 'survey'}
  <SurveyWidget
    data={widget.data}
    id={widget.id}
    survey={widget.additionalData}
  />
{:else if widget.data.type === 'photo'}
  <PhotoWidget urls={widget.files} isTestImage={false} />
{:else if widget.data.type === 'post'}
  <TgPostWidget data={widget.data} />
{:else if widget.data.type === 'activity'}
  <ActivityWidget data={widget.data} />
{/if}

<script lang="ts">
  import {
    createWidget,
    deleteWidget,
    updateWidget,
    changeWidgetPosition,
    updateWidgetsOrder,
  } from '$lib/components/widgetConstructors/widgetsConstructor';

  import { pb } from '$lib/index';

  import { Star, Plus } from '@lucide/svelte';
  import GameWidget from '$lib/components/widgets/GameWidget.svelte';
  import VideoWidget from '$lib/components/widgets/VideoWidget.svelte';
  import TextWidget from '$lib/components/widgets/TextWidget.svelte';
  import SocialWidget from '$lib/components/widgets/SocialWidget.svelte';
  import StickerWidget from '$lib/components/widgets/StickerWidget.svelte';
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import ProgressWidget from '$lib/components/widgets/ProgressWidget.svelte';
  import InterestsWidget from '$lib/components/widgets/InterestsWidget.svelte';
  import ToDoWidget from '$lib/components/widgets/ToDoWidget.svelte';
  import SurveyWidget from '$lib/components/widgets/SurveyWidget.svelte';
  import PhotoWidget from '$lib/components/widgets/PhotoWidget.svelte';
  import TgPostWidget from '$lib/components/widgets/TgPostWidget.svelte';
  import type { BasicWidget } from '$lib/widgetsTypes/widgetsTypes';

  import EditTgPost from '$lib/components/editor/EditTgPost.svelte';
  import EditSurvey from '$lib/components/editor/EditSurvey.svelte';
  import EditText from '$lib/components/editor/EditText.svelte';
  import EditVideo from '$lib/components/editor/EditVideo.svelte';
  import EditAudio from '$lib/components/editor/EditAudio.svelte'
  import type { Option } from '$lib/widgetTypes/widgetTypes';
  import Edit from '$lib/components/editor/Edit.svelte';
  import WidgetsListMenu from '$lib/components/editor/WidgetsListMenu.svelte';
  import { Input } from '$lib/components/ui/input';
  import type { WidgetWithService } from '$lib/components/widgetConstructors/widgetsConstructor';
  import type { PageProps } from '../../../../.svelte-kit/types/src/routes/registration/$types';
  import type { PhotoData } from '$lib/widgetTypes/widgetTypes';

  import ProfileTopLayer from '$lib/components/profile/ProfileTopLayer.svelte';
  import ProfileInfo from '$lib/components/profile/ProfileInfo.svelte';

  let changeMode = $state(false);

  onMount(() => {
    window.Telegram.WebApp.MainButton.show();
    return () => {
      window.Telegram.WebApp.MainButton.hide();
    };
  });

  useTelegramButton(() => {
    changeMode = !changeMode;
  });

  $effect(() => {
    window.Telegram.WebApp.MainButton.setText(
      changeMode ? 'Сохранить' : 'Изменить анкету',
    );
  });

  let showWidgetMenu = $state(false);
  let addedWidget = $state('');
  let nextStage = $state(false); //dakdoawdkpowadkwapodkwapodkwadaopd)))))

  import { scrollY } from 'svelte/reactivity/window';
  import EditImage from '$lib/components/editor/EditImage.svelte';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  import { onMount } from 'svelte';
  import EditSocial from '$lib/components/editor/EditSocial.svelte';
  import EditToDo from '$lib/components/editor/EditToDo.svelte';
  import EditProgress from '$lib/components/editor/EditProgress.svelte';
    import AudioWidget from '$lib/components/widgets/AudioWidget.svelte';

  let height = $derived(Math.max(300, 380 - (scrollY.current ?? 0) * 0.5));

  let { data } = $props();
  let widgets: WidgetWithService[] = $state(data.widgets);

  $effect(() => {
    widgets = data.widgets;
  });
  console.log(widgets);
</script>

{#if height - 90 <= 210}
  <ProfileTopLayer {data} />
{/if}
<img
  src={pb.files.getURL(pb.authStore.record!, data.user_photo)}
  alt="person"
  class="w-full h-95 rounded-b-[24px] object-cover"
  style="height: {height}px"
/>
<div class="container font-[Inter] p-[16px] w-full max-w-full relative">
  <ProfileInfo {data} />
  <!--  <TgPostWidget post="durov/68" />-->

  {#if changeMode}
    <button
      class="my-2 w-full h-12 bg-[#34C759] rounded-xl flex justify-center items-center gap-3"
      onclick={() => (showWidgetMenu = !showWidgetMenu)}
    >
      <Plus class="size-5" />
      <p class="font-medium">Добавить виджет</p>
    </button>
  {/if}

  {#if showWidgetMenu}
    <WidgetsListMenu bind:showWidgetMenu bind:addedWidget bind:nextStage />
  {/if}

  {#if addedWidget === 'Видео'}
    <EditVideo bind:nextStage numberOfWidgets={widgets.length} />
  {:else if addedWidget === 'Текст'}
    <EditText bind:nextStage numberOfWidgets={widgets.length} />
  {:else if addedWidget === 'Опрос'}
    <EditSurvey bind:nextStage numberOfWidgets={widgets.length} />
  {:else if addedWidget === 'Социальная сеть'}
    <EditTgAccount bind:nextStage numberOfWidgets={widgets.length} />
  {:else if addedWidget === 'Телеграм Пост'}
    <EditSocial bind:nextStage numberOfWidgets={widgets.length} />
  {:else if addedWidget === 'Изображение'}
    <EditImage bind:nextStage numberOfWidgets={widgets.length} />
  {:else if addedWidget === 'Список задач'}
    <EditToDo bind:nextStage numberOfWidgets={widgets.length} />
  {:else if addedWidget === 'Прогресс'}
    <EditProgress bind:nextStage numberOfWidgets={widgets.length} />
  {:else if addedWidget === 'Аудио'}
    <EditAudio bind:nextStage numberOfWidgets={widgets.length} />
  {/if}

  {#each widgets as { widget }, i}
    <!--{console.log(widgets[2].additionalData)}-->
    <div class="relative">
      {#if changeMode}
        {console.log(widget.data.type)}
        <Edit
          widgetType={widget.data.type}
          widgetId={widget.id}
          {widgets}
          onMove={(delta) => {
            if ((delta < 0 && i <= 0) || (delta > 0 && i >= widgets.length - 1))
              return;

            [widgets[i + delta], widgets[i]] = [widgets[i], widgets[i + delta]];
          }}
        />
      {/if}
      {#if widget.data.type === 'text'}
        <TextWidget data={widget.data} />
      {:else if widget.data.type === 'audio'}
        <AudioWidget data={widget.data} />
      {:else if widget.data.type === 'steam_game'}
        <GameWidget data={widget.data} />
      {:else if widget.data.type === 'video'}
        <VideoWidget data={widget.data} />
      {:else if widget.data.type === 'social_media'}
        <SocialWidget
          data={widget.data}
          socialMediaData={widget.additionalData}
        />
      {:else if widget.data.type === 'progress_bar'}
        <ProgressWidget data={widget.data} />
        <!--{:else if widget.data.type === 'photo'}-->
        <!--  <ProgressWidget data={widget.data} />-->
      {:else if widget.data.type === 'todo'}
        <ToDoWidget data={widget.data} />
      {:else if widget.data.type === 'survey' && widget.additionalData?.type === 'survey'}
        <SurveyWidget
          data={widget.data}
          id={widget.id}
          survey={widget.additionalData}
        />
      {:else if widget.data.type === 'photo' && widget.additionalData?.type === 'photo'}
        <PhotoWidget data={widget} additionalData={widget.additionalData} />
      {:else if widget.data.type === 'post'}
        <TgPostWidget data={widget.data} />
      {/if}
    </div>
  {/each}
  <!--    <button onclick={() => (open = true)}>fsdfsd</button>-->
  <!--    <Video bind:open bind:addedWidget />-->
  <!--    <StickerWidget image="Stickers/Fire.svg" left={285} top={185} rotate={0} />-->
  <!--    <StickerWidget image="Stickers/Eyes.svg" left={15} top={315} rotate={-30} />-->
  <!--    &lt;!&ndash; Сделать чтобы при True комнонент помечал как выполненное задание &ndash;&gt;-->
  <!--    <ToDoWidget-->
  <!--    title="Что хочу сделать:)"-->
  <!--    tasks={[-->
  <!--      ['hello', true],-->
  <!--      ['world', true],-->
  <!--      [25, true],-->
  <!--      [50, true],-->
  <!--      [75, true],-->
  <!--      [100, false],-->
  <!--    ]}-->
  <!--  />-->
  <!--  <SurveyWidget-->
  <!--    tasks={[-->
  <!--      ['hello', 100],-->
  <!--      ['world', 25],-->
  <!--    ]}-->
  <!--    votes={125}-->
  <!--  />-->
</div>

<style>
  .container {
    margin-bottom: 8px;
  }
</style>

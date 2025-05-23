<script lang="ts">
  import {
    createWidget,
    deleteWidget,
    updateWidget,
    changeWidgetPosition,
    updateWidgetsOrder,
  } from '$lib/components/widgetConstructors/widgetsConstructor';

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
  import type { BasicWidget } from '$lib/widgetsTypes/widgetsTypes';

  import EditTgPost from '$lib/components/editor/EditTgPost.svelte';
  import EditTgAccount from '$lib/components/editor/EditTgAccount.svelte';
  import EditSurvey from '$lib/components/editor/EditSurvey.svelte';
  import EditText from '$lib/components/editor/EditText.svelte';
  import EditVideo from '$lib/components/editor/EditVideo.svelte';
  import type { Option } from '$lib/widgetTypes/widgetTypes';
  import Edit from '$lib/components/editor/Edit.svelte';
  import WidgetsListMenu from '$lib/components/editor/WidgetsListMenu.svelte';
  import { Input } from '$lib/components/ui/input';
  import type { WidgetWithService } from '$lib/components/widgetConstructors/widgetsConstructor';
  import type { PageProps } from '../../../../.svelte-kit/types/src/routes/registration/$types';

  let changeMode = $state(false);
  $effect(() => {
    window.Telegram.WebApp.MainButton.setParams({
      text: changeMode ? 'Сохранить' : 'Изменить анкету',
      is_visible: true,
    });
  });
  window.Telegram.WebApp.MainButton.onClick(() => (changeMode = !changeMode));

  let showWidgetMenu = $state(false);
  let addedWidget = $state('');
  let nextStage = $state(false); //dakdoawdkpowadkwapodkwapodkwadaopd)))))

  import { scrollY } from 'svelte/reactivity/window';
  import EditImage from '$lib/components/editor/EditImage.svelte';

  let height = $derived(Math.max(300, 380 - (scrollY.current ?? 0) * 0.5));

  let { data } = $props();
  let widgets: WidgetWithService[] = $derived(data.widgets);

  let widgetsLenght = $derived(widgets.length);
</script>

{#if height - 90 <= 210}
  <div class="w-full h-16 bg-background fixed flex z-2">
    <img src="Girl.png" class="w-16 h-16 rounded-full p-2" alt="UserPhoto" />
    <div class="container font-[Inter] p-[4px]">
      <p class="font-extrabold text-[16px] flex items-center">
        Илона Абудаби, 18<Star />
      </p>
      <p class="font-semibold text-[16px]"><span>Махачкала, Россия</span></p>
    </div>
  </div>
{/if}
<img
  src="Girl.png"
  alt="person"
  class="w-full h-95 rounded-b-[24px] object-cover"
  style="height: {height}px"
/>
<div class="container font-[Inter] p-[16px] w-full max-w-full relative">
  <p class="font-extrabold text-[32px] flex items-center">
    Илона Абудаби, 18<Star />
  </p>
  <p class="font-semibold text-[20px]"><span>Екатеринбург, Россия</span></p>
  <InterestsWidget
    interests={[
      ['🎮', 'Игры'],
      ['🚗', 'Развлечения'],
      ['💅', 'Красота'],
      ['🧮', 'Математика'],
      ['🎮', 'Кодинг'],
      ['🎮', 'игры'],
      ['🎮', 'игры'],
      ['🎮', 'игры'],
    ]}
  />

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
  {:else if addedWidget === 'Телеграм Аккаунт'}
    <EditTgAccount bind:nextStage numberOfWidgets={widgets.length} />
  {:else if addedWidget === 'Телеграм Пост'}
    <EditTgPost bind:nextStage numberOfWidgets={widgets.length} />
  {:else if addedWidget === 'Изображение'}
    <EditImage bind:nextStage numberOfWidgets={widgets.length} />
  {/if}

  {#each widgets as { widget }}
    <div class="relative">
      {#if changeMode}
        <Edit widgetType={widget.data.type} widgetId={widget.id} {widgets} />
      {/if}
      {#if widget.data.type === 'text'}
        <TextWidget data={widget.data} />
      {:else if widget.data.type === 'steam_game'}
        <GameWidget data={widget.data} />
      {:else if widget.data.type === 'video'}
        <VideoWidget data={widget.data} />
      {:else if widget.data.type === 'social_media'}
        <SocialWidget
          data={widget.data}
          socialMediaData={widget.additionalData.socialMediaData}
        />
      {:else if widget.data.type === 'progress_bar'}
        <ProgressWidget data={widget.data} />
        <!--{:else if widget.data.type === 'photo'}-->
        <!--  <ProgressWidget data={widget.data} />-->
      {:else if widget.data.type === 'todo'}
        <ToDoWidget data={widget.data} />
      {:else if widget.data.type === 'survey'}
        <SurveyWidget data={widget.data} />
      {:else if widget.data.type === 'photo'}
        <PhotoWidget data={widget.data} />
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

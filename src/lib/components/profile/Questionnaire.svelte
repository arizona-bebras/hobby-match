<script lang="ts">
  import { Plus } from '@lucide/svelte';
  import EditTgPost from '$lib/components/editor/EditTgPost.svelte';
  import EditSurvey from '$lib/components/editor/EditSurvey.svelte';
  import EditText from '$lib/components/editor/EditText.svelte';
  import EditVideo from '$lib/components/editor/EditVideo.svelte';
  import type { WidgetType } from '$lib/widgetTypes/widgetTypes';
  import Edit from '$lib/components/editor/Edit.svelte';
  import WidgetsListMenu from '$lib/components/editor/WidgetsListMenu.svelte';
  import EditImage from '$lib/components/editor/EditImage.svelte';
  import { useTelegramButton } from '$lib/components/registration/useTelegramButton.svelte';
  import EditSocial from '$lib/components/editor/EditSocial.svelte';
  import EditToDo from '$lib/components/editor/EditToDo.svelte';
  import EditProgress from '$lib/components/editor/EditProgress.svelte';
  import type { WidgetWithService } from '$lib/components/widgetConstructors/widgetsConstructor';
  import EditAudio from '$lib/components/editor/EditAudio.svelte';
  import EditSteamGame from '$lib/components/editor/EditSteamGame.svelte';
  import Header from '$lib/components/profile/Header.svelte';
  import UserInfo from '$lib/components/profile/UserInfo.svelte';
  import RenderWidget from '$lib/components/profile/RenderWidget.svelte';
  import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';

  let changeMode = $state(false);
  let showWidgetMenu = $state(false);
  let addedWidget: WidgetType | undefined = $state(undefined);
  let editingWidget: string | undefined = $state(undefined);

  let { data }: { data: PageData } = $props();
  console.log(data);
  let widgets = $state(data.widgets);
  console.log('Data', data);
  console.log('Widgets:', widgets);
  useTelegramButton(() => {
    changeMode = !changeMode;
  });

  $effect(() => {
    window.Telegram.WebApp.MainButton.setText(
      changeMode ? 'Сохранить' : 'Моя анкета',
    );
  });

  $effect(() => {
    if (!addedWidget && !showWidgetMenu)
      window.Telegram.WebApp.MainButton.show();
    else window.Telegram.WebApp.MainButton.hide();
    return () => {
      window.Telegram.WebApp.MainButton.hide();
    };
  });

  $effect(() => {
    widgets = data.widgets;
  });

  function onEditClose() {
    addedWidget = undefined;
    editingWidget = undefined;
  }
</script>

<Header {data} {changeMode} />
<div class="font-[Inter] px-4 w-full max-w-full relative">
  <UserInfo {data} {changeMode} />

  {#if changeMode}
    <button
      class="my-2 w-full h-12 bg-[#34C759] rounded-xl flex justify-center items-center gap-3 font-medium"
      onclick={() => (showWidgetMenu = true)}
    >
      <Plus class="size-5" />
      Добавить виджет
    </button>
  {/if}

  <WidgetsListMenu
    open={showWidgetMenu}
    onClick={(name) => {
      showWidgetMenu = false;
      addedWidget = name;
    }}
  />

  <EditVideo
    open={addedWidget === 'video'}
    onClose={onEditClose}
    widgetId={addedWidget === 'video' ? editingWidget : undefined}
  />
  <EditProgress
    open={addedWidget === 'progress_bar'}
    onClose={onEditClose}
    widgetId={addedWidget === 'progress_bar' ? editingWidget : undefined}
  />
  <EditText
    open={addedWidget === 'text'}
    onClose={onEditClose}
    widgetId={addedWidget === 'text' ? editingWidget : undefined}
  />
  <EditSurvey
    open={addedWidget === 'survey'}
    onClose={onEditClose}
    widgetId={addedWidget === 'survey' ? editingWidget : undefined}
  />
  <EditSocial
    open={addedWidget === 'social_media'}
    onClose={onEditClose}
    widgetId={addedWidget === 'social_media' ? editingWidget : undefined}
  />
  <EditTgPost
    open={addedWidget === 'post'}
    onClose={onEditClose}
    widgetId={addedWidget === 'post' ? editingWidget : undefined}
  />
  <EditImage
    open={addedWidget === 'photo'}
    onClose={onEditClose}
    widgetId={addedWidget === 'photo' ? editingWidget : undefined}
  />
  <EditToDo
    open={addedWidget === 'todo'}
    onClose={onEditClose}
    widgetId={addedWidget === 'todo' ? editingWidget : undefined}
  />
  <EditAudio
    open={addedWidget === 'audio'}
    onClose={onEditClose}
    widgetId={addedWidget === 'audio' ? editingWidget : undefined}
  />
  <EditSteamGame
    open={addedWidget === 'steam_game'}
    onClose={onEditClose}
    widgetId={addedWidget === 'steam_game' ? editingWidget : undefined}
  />
  {#each widgets as widget, i (widget.id)}
    <div class="relative mb-2">
      {#if changeMode}
        <Edit
          {widget}
          onEdit={() => {
            addedWidget = widget.data.type;
            editingWidget = widget.id;
          }}
          onMove={(delta) => {
            if ((delta < 0 && i <= 0) || (delta > 0 && i >= widgets.length - 1))
              return;

            [widgets[i + delta], widgets[i]] = [widgets[i], widgets[i + delta]];
          }}
        />
      {/if}
      <RenderWidget {widget} />
    </div>
  {/each}
</div>

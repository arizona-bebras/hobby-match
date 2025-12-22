<script lang="ts">
  import {
    EditVideo,
    EditSurvey,
    EditText,
    EditTgPost,
    EditProgress,
    EditToDo,
    EditImage,
    EditSocial,
    EditAudio,
    EditSteamGame,
  } from '$lib/components/editor/index.js';
  import {
    type Widget,
    type WidgetType,
  } from '$lib/widgetTypes/widgetTypes.js';

  const {
    addedWidget,
    onEditClose,
    widgetId,
    widgetData,
    files,
  }: {
    addedWidget: WidgetType;
    onEditClose: () => void;
    widgetId: string;
    widgetData: Widget['data'];
    files: string[];
  } = $props();
  console.log(widgetData, 'Данные виджета!');
</script>

<EditVideo
  open={addedWidget === 'video'}
  onClose={onEditClose}
  widgetId={addedWidget === 'video' ? widgetId : undefined}
  {widgetData}
/>
<EditProgress
  open={addedWidget === 'progress_bar'}
  onClose={onEditClose}
  widgetId={addedWidget === 'progress_bar' ? widgetId : undefined}
  {widgetData}
/>
<EditText
  open={addedWidget === 'text'}
  onClose={onEditClose}
  widgetId={addedWidget === 'text' ? widgetId : undefined}
  {widgetData}
/>
<EditSurvey
  open={addedWidget === 'survey'}
  onClose={onEditClose}
  widgetId={addedWidget === 'survey' ? widgetId : undefined}
  {widgetData}
/>
<EditSocial
  open={addedWidget === 'social_media'}
  onClose={onEditClose}
  widgetId={addedWidget === 'social_media' ? widgetId : undefined}
  {widgetData}
/>
<EditTgPost
  open={addedWidget === 'post'}
  onClose={onEditClose}
  widgetId={addedWidget === 'post' ? widgetId : undefined}
  {widgetData}
/>

<EditImage
  open={addedWidget === 'photo'}
  onClose={onEditClose}
  widgetId={addedWidget === 'photo' ? widgetId : undefined}
  {files}
  {widgetData}
/>

<EditToDo
  open={addedWidget === 'todo'}
  onClose={onEditClose}
  widgetId={addedWidget === 'todo' ? widgetId : undefined}
  {widgetData}
/>
<EditAudio
  open={addedWidget === 'audio'}
  onClose={onEditClose}
  widgetId={addedWidget === 'audio' ? widgetId : undefined}
  {widgetData}
/>
<EditSteamGame
  open={addedWidget === 'steam_game'}
  onClose={onEditClose}
  widgetId={addedWidget === 'steam_game' ? widgetId : undefined}
  {widgetData}
/>

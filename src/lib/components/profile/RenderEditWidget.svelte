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
  import { type WidgetType } from '$lib/widgetTypes/widgetTypes.js';

  const {
    addedWidget,
    onEditClose,
    widgetId,
  }: { addedWidget: WidgetType; onEditClose: () => void; widgetId: string } =
    $props();
</script>

<EditVideo
  open={addedWidget === 'video'}
  onClose={onEditClose}
  widgetId={addedWidget === 'video' ? widgetId : undefined}
/>
<EditProgress
  open={addedWidget === 'progress_bar'}
  onClose={onEditClose}
  widgetId={addedWidget === 'progress_bar' ? widgetId : undefined}
/>
<EditText
  open={addedWidget === 'text'}
  onClose={onEditClose}
  widgetId={addedWidget === 'text' ? widgetId : undefined}
/>
<EditSurvey
  open={addedWidget === 'survey'}
  onClose={onEditClose}
  widgetId={addedWidget === 'survey' ? widgetId : undefined}
/>
<EditSocial
  open={addedWidget === 'social_media'}
  onClose={onEditClose}
  widgetId={addedWidget === 'social_media' ? widgetId : undefined}
/>
<EditTgPost
  open={addedWidget === 'post'}
  onClose={onEditClose}
  widgetId={addedWidget === 'post' ? widgetId : undefined}
/>
<EditImage
  open={addedWidget === 'photo'}
  onClose={onEditClose}
  widgetId={addedWidget === 'photo' ? widgetId : undefined}
/>
<EditToDo
  open={addedWidget === 'todo'}
  onClose={onEditClose}
  widgetId={addedWidget === 'todo' ? widgetId : undefined}
/>
<EditAudio
  open={addedWidget === 'audio'}
  onClose={onEditClose}
  widgetId={addedWidget === 'audio' ? widgetId : undefined}
/>
<EditSteamGame
  open={addedWidget === 'steam_game'}
  onClose={onEditClose}
  widgetId={addedWidget === 'steam_game' ? widgetId : undefined}
/>

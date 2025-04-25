<script lang="ts">
  import { invalidate } from '$app/navigation';
  import { onMount } from 'svelte';
  import { pb } from '$lib/index';
  import TextWidget from '$lib/components/widgets/TextWidget.svelte';
  import SurveyOption from '$lib/components/widgetConstructors/SurveyOption.svelte';
  import Task from '$lib/components/widgetConstructors/Task.svelte';
  import type { Widget } from '$lib/widgetTypes/widgetTypes';
  import type { RecordModel } from 'pocketbase';
  import AudioWidget from '$lib/components/widgets/AudioWidget.svelte';
  //import { processText } from '../../../lib/processWidgets'

  let tasksNumbers: Array<{ id: number }> = $state([{ id: 1 }]);
  let surveyOptionsNumbers: Array<{ id: number }> = $state([{ id: 1 }]);

  let widgetTypes = [
    { widgetType: 'audio', widgetDisplay: 'Аудио' },
    { widgetType: 'video', widgetDisplay: 'Видео' },
    { widgetType: 'photo', widgetDisplay: 'Фото' },
    { widgetType: 'todo', widgetDisplay: 'Списко заданий' },
    { widgetType: 'progress_bar', widgetDisplay: 'Отслеживание прогресса' },
    { widgetType: 'geo', widgetDisplay: 'Геолокаия' },
    { widgetType: 'social_media', widgetDisplay: 'Социальная сеть' },
    { widgetType: 'steam_game', widgetDisplay: 'Статистика игры в Steam' },
    { widgetType: 'sticker', widgetDisplay: 'Стикер' },
    { widgetType: 'survey', widgetDisplay: 'Опрос' },
    { widgetType: 'text', widgetDisplay: 'Текст' },
  ];

  let selectedWidget = $state<(typeof widgetTypes)[number]>();

  let widget: Widget = {
    telegram_id: '',
    order: 1,
    data: {},
  };

  function removeTask(id: number) {
    tasksNumbers = tasksNumbers.filter((component) => component.id !== id);
    for (let i = 0; i < tasksNumbers.length; i++) {
      tasksNumbers[i].id = i + 1;
    }
  }

  function removeOption(id: number) {
    surveyOptionsNumbers = surveyOptionsNumbers.filter(
      (component) => component.id !== id,
    );
    for (let i = 0; i < surveyOptionsNumbers.length; i++) {
      surveyOptionsNumbers[i].id = i + 1;
    }
  }

  function processWidgetData(): object {
    if (selectedWidget?.widgetType == 'text') {
      return processText();
    }
    if (selectedWidget?.widgetType == 'audio') {
      return processAudio();
    }
    return {};
  }

  function processText(): object {
    const textArea = document.getElementById('input-widget-text');
    console.log(textArea);
    const text = textArea?.value;
    return {
      text: text,
    };
  }

  function processAudio(): object {
    const link = document.getElementById('audio-link')?.value;
    return {
      link: link,
    };
  }

  let widgetsStatus = $state<Array<boolean>>([]);
  let addedWidgets = $state(0);

  async function createWidget() {
    console.log(selectedWidget?.widgetType);
    const record = await pb.collection('widgets').create({
      telegram_id: pb.authStore.record?.telegram_id,
      type: selectedWidget?.widgetType,
      order: widgetsStatus.length + 1,
      data: processWidgetData(),
    });
    addedWidgets++;
    widgetsStatus.push(false);
  }

  async function getUserWidgets() {
    const widgetList = await pb.collection('widgets').getFullList({
      telegram_id: `${pb.authStore.record?.telegram_id}`,
      sort: `+order`,
    });
    for (let i = 0; i < widgetList.length; i++) {
      widgetsStatus.push(false);
    }
    widget = {
      telegram_id: pb.authStore.record?.telegram_id,
      order: widgetsStatus.length + 1,
      data: {},
    };
    return widgetList;
  }

  async function updateWidgetsOrder(widgets: RecordModel[]) {
    for (let i = 1; i <= widgets.length; i++) {
      await pb.collection('widgets').update(widgets[i - 1].id, {
        order: i,
      });
      console.log(widgets[i - 1].id);
    }
  }

  async function deleteWidget(widgets: RecordModel[], widget: RecordModel) {
    await pb.collection('widgets').delete(widget.id);
    widgetsStatus[widget.order - 1] = true;
    console.log(widgetsStatus);
    const filtred = [];
    for (let i = 0; i < widgetsStatus.length; i++) {
      if (!widgetsStatus[i]) {
        filtred.push(widgets[i]);
      }
    }
    widgets = filtred;
    console.log(widgets);
    updateWidgetsOrder(filtred);
  }

  function saveChanges() {
    window.location.reload();
  }

  const widgetsPromise = getUserWidgets();
</script>

<main>
  <form class="widget-form">
    <select bind:value={selectedWidget}>
      {#each widgetTypes as widget}
        <option value={widget}>
          {widget.widgetDisplay}
        </option>
      {/each}
    </select>
    <br />
    {#if selectedWidget?.widgetType == 'audio'}
      <p>Ссылка на sc</p>
      <input type="text" id="audio-link" />
    {:else if selectedWidget?.widgetType == 'video'}
      <p>Ссылка на видео</p>
      <input type="text" />
    {:else if selectedWidget?.widgetType == 'photo'}
      <p>Фото</p>
      <input type="file" />
    {:else if selectedWidget?.widgetType == 'todo'}
      {#each tasksNumbers as task (task.id)}
        <p>{task.id}</p>
        <Task onRemove={() => removeTask(task.id)} />
      {/each}
      <button
        onclick={() =>
          (tasksNumbers = [...tasksNumbers, { id: tasksNumbers.length + 1 }])}
        id="taskCreateButton">Добавить Цель</button
      >
    {:else if selectedWidget?.widgetType == 'progress_bar'}
      <p>Название цели</p>
      <input type="text" />
      Текущий прогресс
      <input type="text" />
      Максимальный прогресс
      <input type="text" />
    {:else if selectedWidget?.widgetType == 'social_media'}
      <p>Ссылка на соц. сеть</p>
      <input type="text" />
    {:else if selectedWidget?.widgetType == 'steam_game'}
      <p>Ссылка на профиль Steam</p>
      <input type="text" />
    {:else if selectedWidget?.widgetType == 'sticker'}
      <p>Координаты</p>
      <input type="text" />
    {:else if selectedWidget?.widgetType == 'survey'}
      <p>Название опроса</p>
      <input type="text" />
      {#each surveyOptionsNumbers as task (task.id)}
        <SurveyOption onRemove={() => removeOption(task.id)} />
      {/each}
      <button
        onclick={() =>
          (surveyOptionsNumbers = [
            ...surveyOptionsNumbers,
            { id: surveyOptionsNumbers.length + 1 },
          ])}
        id="taskCreateButton">Добавить вариант ответа</button
      >
    {:else if selectedWidget?.widgetType == 'text'}
      <p>Текст</p>
      <textarea class="input-widget-text" id="input-widget-text"></textarea>
    {/if}
  </form>
  <button class="submit-btn" onclick={createWidget}>Добавить виджет</button>
  <div>Добавлено виджетов: {addedWidgets}</div>
  <button onclick={saveChanges}>Сохранить изменения</button>
  {#await widgetsPromise}
    <p>Загрузка виджетов...</p>
  {:then widgets}
    {console.log(widgets)}
    {#each widgets as widget}
      {#if widget.type == 'text'}
        <div
          class={widgetsStatus[widget.order - 1]
            ? 'widget-container deleted'
            : 'widget-container'}
        >
          <TextWidget text={widget.data.text} />
          <button onclick={deleteWidget(widgets, widget)}>X</button>
        </div>
      {:else if widget.type == 'audio'}
        <div
          class={widgetsStatus[widget.order - 1]
            ? 'widget-container deleted'
            : 'widget-container'}
        >
          <AudioWidget link={widget.data.link} />
          <button onclick={deleteWidget(widgets, widget)}>X</button>
        </div>
      {/if}
    {/each}
  {/await}
</main>

<style>
  input {
    outline: solid;
  }
  .submit-btn {
    padding-top: 30px;
  }
  .widget-form {
    display: flex;
    flex-direction: column;
    max-width: 300px;
  }
  .input-widget-text {
    min-width: 300px;
    min-height: 200px;
  }
  .widget-container {
    display: flex;
    flex-direction: row;
    margin-bottom: 30px;
  }
  .deleted {
    background-color: #ffebee;
    border-color: #f44336;
    opacity: 0.7;
    transform: scale(0.98);
  }
</style>

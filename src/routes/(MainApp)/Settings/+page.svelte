<script lang="ts">
  import { invalidate } from '$app/navigation';
  import { onMount } from 'svelte';
  import { diefinePlatofrm, getImageDimensions, getUsernameFromUrl, getYoutubeChannelStats, pb } from '$lib/index';
  import { createWidget, deleteWidget, updateWidget, changeWidgetPostion, updateWidgetsOrder } from '$lib/components/widgetConstructors/widgetsConstructor';
  import type { WidgetWithService } from '$lib/components/widgetConstructors/widgetsConstructor';
  import { getSocialMediaData } from '$lib/index';
  import TextWidget from '$lib/components/widgets/TextWidget.svelte';
  import SurveyOption from '$lib/components/widgetConstructors/SurveyOption.svelte';
  import Task from '$lib/components/widgetConstructors/Task.svelte';
  import * as wid from '$lib/widgetTypes/widgetTypes';
  import type { RecordModel } from 'pocketbase';
  import type { RequestHandler } from '@sveltejs/kit';
  import AudioWidget from '$lib/components/widgets/AudioWidget.svelte';
  import { enhance } from '$app/forms';
  import VideoWidget from '$lib/components/widgets/VideoWidget.svelte';
  import PhotoWidget from '$lib/components/widgets/PhotoWidget.svelte';
  import ToDoWidget from '$lib/components/widgets/ToDoWidget.svelte';
  import ProgressWidget from '$lib/components/widgets/ProgressWidget.svelte';
  import SocialWidget from '$lib/components/widgets/SocialWidget.svelte';
  //import { processText } from '../../../lib/processWidgets'

  let { data } = $props();

  let  { widgets } = $state(data);
  console.log($state.snapshot(widgets))
  updateWidgetsOrder(widgets);

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

  let selectedWidget = $state();

    //widget constructor logic

    function removeTask(id: number) {
        tasksNumbers = tasksNumbers.filter(component => component.id !== id);
        for (let i = 0; i< tasksNumbers.length; i++) {
            tasksNumbers[i].id = i + 1
        }
    };

    function removeOption(id: number) {
    surveyOptionsNumbers = surveyOptionsNumbers.filter(
        (component) => component.id !== id,
    );
    for (let i = 0; i < surveyOptionsNumbers.length; i++) {
        surveyOptionsNumbers[i].id = i + 1;
    }
    }

    /*let widgets = $state<Array<wid.Widget>>([]);
    let widgetsDeleteStatus = $state<Array<boolean>>([]);
    let widgetsChangeStatus = $state<Array<boolean>>([]);
    let addedWidgets = $state(0)

    async function getUserWidgets() {
        let userWidgets: wid.Widget[] = []; 
        const widgetList = await pb.collection('widgets').getFullList({
            sort: `+order`
        })
        for (let i = 0; i < widgetList.length; i++) {
            widgetsDeleteStatus.push(false)
            widgetsChangeStatus.push(false)
            userWidgets.push(converRecordToWidget(widgetList[i]))
        }
        return userWidgets;
    }

    async function getFilesURL(widgetId: string) {
      const record = await pb.collection('widgets').getOne(widgetId);
      const urls: string[] = [];
      const photos = record.files;
      console.log(photos)
      for (let photo of photos) {
        urls.push(pb.files.getURL(record, photo));
      }
      console.log(urls)
      return urls;
    }

    async function loadWidgets() {
        widgets = await getUserWidgets();
    }
    */
    function showChangeWidgetField(widget: WidgetWithService) {
        widget.changeStatus = !widget.changeStatus;
    };

    /*function saveChanges() {
        window.location.reload();
    }

    const widgetsPromise = getUserWidgets();*/
</script>

<main>
  <form 
    method="POST"
    use:enhance={() => createWidget}
    enctype="multipart/form-data"
    class="widget-form"
    >
    <select name="type" bind:value={selectedWidget}>
      {#each widgetTypes as widget}
        <option value={widget.widgetType}>
          {widget.widgetDisplay}
        </option>
      {/each}
    </select>
    <br />
    {#if selectedWidget == 'audio'}
      <p>Ссылка</p>
      <input type="text" id="audio-link" name="link" />
    {:else if selectedWidget == 'video'}
      <p>Ссылка на видео</p>
      <input type="text" name="link"/>
    {:else if selectedWidget == 'photo'}
      <p>Выберите до 3 фото</p>
      <input type="file" name="files" multiple/>
    {:else if selectedWidget == 'todo'}
      <input type="text" name="title" />
      {#each tasksNumbers as task (task.id)}
        <Task 
          onRemove={() => removeTask(task.id)} 
          taskOrder = {task.id}
          />
      {/each}
      <button
        onclick={() =>
          (tasksNumbers = [...tasksNumbers, { id: tasksNumbers.length + 1 }])}
        id="taskCreateButton" type="button">Добавить Цель</button
      >
    {:else if selectedWidget == 'progress_bar'}
      <p>Название цели</p>
      <input type="text" name="description"/>
      Текущий прогресс
      <input type="text" name="currentProgress"/>
      Максимальный прогресс
      <input type="text" name="maxProgress"/>
    {:else if selectedWidget == 'social_media'}
      <p>Ссылка на соц. сеть</p>
      <input type="text" name="link"/>
    {:else if selectedWidget == 'steam_game'}
      <p>Ссылка на профиль Steam</p>
      <input type="text" />
    {:else if selectedWidget == 'sticker'}
      <p>Координаты</p>
      <input type="text" />
    {:else if selectedWidget == 'survey'}
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
    {:else if selectedWidget == 'text'}
      <p>Текст</p>
      <textarea class="input-widget-text" id="input-widget-text" name="text"></textarea>
    {/if}
    <div class="submit-btn">
      <button type="submit">Добавить виджет</button>
    </div>
  </form>
    {#each widgets as widget}
        {#if (widget.widget.data.type == "text")}
            <div 
                class={widget.deleteStatus ? "widget-container deleted" : "widget-container"}
            >
                <TextWidget 
                    data = {widget.widget.data}
                />
                
                <button style="margin-right:15px" onclick={() => showChangeWidgetField(widget)}>
                    {widget.changeStatus ? 'Отмена' : 'Изменить'}
                </button>
                <button class="ch-btn" onclick={() => deleteWidget(widgets, widget)}>X</button>
                <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, -1)}>^</button>
                <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, 1)}>v</button>
            </div>
            {#if (widget.changeStatus == true)} 
                <form method = "POST" use:enhance={() => updateWidget(widget)} style="margin-bottom:15px">
                    <input type="text" name ="text">
                    <button type="submit">Подтверить</button>
                </form>
            {/if}
        {:else if (widget.widget.data.type == "audio")}
            <div 
                class={widget.deleteStatus ? "widget-container deleted" : "widget-container"}
            >
                <AudioWidget 
                    link = {widget.widget.data.link}
                />
                <button class="ch-btn" onclick={() => deleteWidget(widgets, widget)}>X</button>
                <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, -1)}>^</button>
                <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, 1)}>v</button>
            </div>
          {:else if (widget.widget.data.type == "video")}
            <div 
                class={widget.deleteStatus ? "widget-container deleted" : "widget-container"}
            >
                <VideoWidget 
                    data = {widget.widget.data}
                />
                <button class="ch-btn" onclick={() => deleteWidget(widgets, widget)}>X</button>
                <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, -1)}>^</button>
                <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, 1)}>v</button>
            </div>
          {:else if (widget.widget.data.type == "todo")}
            <div 
                class={widget.deleteStatus ? "widget-container deleted" : "widget-container"}
            >
                <ToDoWidget 
                    title = {widget.widget.data.title}
                    tasks = {widget.widget.data.tasks}
                    widgetId = {widget.widget.id}
                />
                <button class="ch-btn" onclick={() => deleteWidget(widgets, widget)}>X</button>
                <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, -1)}>^</button>
                <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, 1)}>v</button>
            </div>
          {:else if (widget.widget.data.type == "progress_bar")}
            <div 
                class={widget.deleteStatus ? "widget-container deleted" : "widget-container"}
            >
                <ProgressWidget 
                    data = { widget.widget.data }
                />
                <button class="ch-btn" onclick={() => deleteWidget(widgets, widget)}>X</button>
                <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, -1)}>^</button>
                <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, 1)}>v</button>
            </div>
          {:else if (widget.widget.data.type == "social_media")}
          <div 
              class={widget.deleteStatus ? "widget-container deleted" : "widget-container"}
          >
              <SocialWidget 
                  data = {widget.widget.data}
                  socialMeidaData = {widget.additionalData.socialMeidaData}
              />
              <button class="ch-btn" onclick={() => deleteWidget(widgets, widget)}>X</button>
              <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, -1)}>^</button>
              <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, 1)}>v</button>
          </div>
        {/if}
    {/each}
</main>

<style>
  body{
    color: black;
  }
  input {
    outline: solid;
  }
  .submit-btn {
    padding-top: 30px;
  }
  .ch-btn {
    margin-right: 5px;
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

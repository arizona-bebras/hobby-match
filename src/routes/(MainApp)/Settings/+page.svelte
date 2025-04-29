<script lang="ts">
  import { invalidate } from '$app/navigation';
  import { onMount } from 'svelte';
  import { diefinePlatofrm, getImageDimensions, pb } from '$lib/index';
  import { swapElements } from '$lib/index';
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

    function converRecordToWidget(record: RecordModel) : wid.Widget {
      let widget : wid.Widget = {
          id: record.id,
          telegram_id: record.telegram_id,
          order: record.order,
          type: record.type,
          data: record.data
      }
      return widget;
    }

    let widgets = $state<Array<wid.Widget>>([]);
    let widgetsStatus = $state<Array<boolean>>([]);
    let widgetsChangeStatus = $state<Array<boolean>>([]);
    let addedWidgets = $state(0)

    async function getUserWidgets() {
        let userWidgets: wid.Widget[] = []; 
        const widgetList = await pb.collection('widgets').getFullList({
            telegram_id: `${pb.authStore.record?.telegram_id}`,
            sort: `+order`
        })
        for (let i = 0; i < widgetList.length; i++) {
            widgetsStatus.push(false)
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

    //widgets logic

    async function createWidget( { formData }: any ) {
        
        const uploadedFiles = formData.getAll('files') as File[];
        let formValues = Object.fromEntries(formData);

        console.log(formValues)

        delete formValues.files;

        if (selectedWidget?.widgetType == "video") { 
          formValues.platform = diefinePlatofrm(formValues.link);
        }

        if (selectedWidget?.widgetType == "photo") { 
          formValues.photos = [];
          for (const file of uploadedFiles) {
            const { width, height } : any = await getImageDimensions(file);
            formValues.photos.push({
              name: file.name,
              width : width,
              height : height,
              size : file.size
            })
          }
        }

        if (selectedWidget?.widgetType == "todo") {
          const formValuesList = Object.entries(formValues);
          formValues.tasks = [];
          for (const [key, value] of formValuesList) {
            if (key.includes("task")) {
              let task : wid.Task = {
                order: parseInt(key.slice(4)),
                description: value,
                isCompleted: false
              }
              formValues.tasks.push(task)
              delete formValues[key]
            }
          }
        }

        await pb.collection("widgets").create({
            "telegram_id": pb.authStore.record?.telegram_id,
            "type": selectedWidget?.widgetType,
            "order" : widgetsStatus.length+1,
            "files" : uploadedFiles,
            "data": formValues
        })
        addedWidgets++;
        widgetsStatus.push(false)

        return { status: 200 };
    }

    async function updateWidgetsOrder(widgets:wid.Widget[]) {
        for (let i = 1; i <= widgets.length; i++) {
            await pb.collection('widgets').update(widgets[i-1].id, {
                    "order": i
                }
            )
            console.log(widgets[i-1].id)
        }
    }

    async function deleteWidget(widgets: wid.Widget[], widget: wid.Widget) {
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

    async function updateWidget(widget: wid.Widget) {
        return async ({ formData } : any) => {
            const formValues = Object.fromEntries(formData);
            await pb.collection('widgets').update(widget.id, {
            data: formValues
            });
        widgetsChangeStatus[widget.order-1] = false;
        return { status: 200 };
        };
    }

    async function changeWidgetPostion(widgets: wid.Widget[], widget: wid.Widget, posChange: number) {
        const record = await pb.collection('widgets').getFullList({
            filter: `telegram_id = "${pb.authStore.record?.telegram_id}" && order = "${widget.order + posChange}"`
        }).then(record => record)

        let widgetOrder = widget.order
        let recordOrder = record[0].order
        await pb.collection('widgets').update(record[0].id, {
            "order": recordOrder + (posChange * -1)
        })
        await pb.collection('widgets').update(widget.id, {
            "order": widgetOrder + posChange
        })
        const temp = widgets[recordOrder-1];
        widgets[recordOrder-1].order += (posChange * -1);
        widgets[recordOrder-1] = widgets[widgetOrder-1];
        widgets[widgetOrder-1].order += posChange;
        widgets[widgetOrder-1] = temp;
    }

    function showChangeWidgetField(widget: wid.Widget) {
        widgetsChangeStatus[widget.order-1] = !widgetsChangeStatus[widget.order-1];
    };

    function saveChanges() {
        window.location.reload();
    }

    const widgetsPromise = getUserWidgets();
</script>

<main>
  <form 
    method="POST"
    use:enhance={createWidget}
    enctype="multipart/form-data"
    class="widget-form"
    >
    <select bind:value={selectedWidget}>
      {#each widgetTypes as widget}
        <option value={widget}>
          {widget.widgetDisplay}
        </option>
      {/each}
    </select>
    <br />
    {#if selectedWidget?.widgetType == 'audio'}
      <p>Ссылка</p>
      <input type="text" id="audio-link" name="link" />
    {:else if selectedWidget?.widgetType == 'video'}
      <p>Ссылка на видео</p>
      <input type="text" name="link"/>
    {:else if selectedWidget?.widgetType == 'photo'}
      <p>Выберите до 3 фото</p>
      <input type="file" name="files" multiple/>
    {:else if selectedWidget?.widgetType == 'todo'}
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
    {:else if selectedWidget?.widgetType == 'progress_bar'}
      <p>Название цели</p>
      <input type="text" name="description"/>
      Текущий прогресс
      <input type="text" name="currentProgress"/>
      Максимальный прогресс
      <input type="text" name="maxProgress"/>
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
      <textarea class="input-widget-text" id="input-widget-text" name="text"></textarea>
    {/if}
    <div class="submit-btn">
      <button type="submit">Добавить виджет</button>
    </div>
  </form>
  <div>Добавлено виджетов: {addedWidgets}</div>
  <button onclick={saveChanges}>Сохранить изменения</button>
  {#await loadWidgets()}
    <p>Загрузка виджетов...</p>
  {:then _}
    {console.log(widgets)}
        {#each widgets as widget}
            {#if (widget.type == "text")}
                <div 
                    class={widgetsStatus[widget.order-1] ? "widget-container deleted" : "widget-container"}
                >
                    <TextWidget 
                        text = {widget.data.text}
                    />
                    
                    <button style="margin-right:15px" onclick={() => showChangeWidgetField(widget)}>
                        {widgetsChangeStatus[widget.order-1] ? 'Отмена' : 'Изменить'}
                    </button>
                    <button class="ch-btn" onclick={() => deleteWidget(widgets, widget)}>X</button>
                    <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, -1)}>^</button>
                    <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, 1)}>v</button>
                </div>
                {#if (widgetsChangeStatus[widget.order-1] == true)} 
                    <form method = "POST" use:enhance={() => updateWidget(widget)} style="margin-bottom:15px">
                        <input type="text" name ="text">
                        <button type="submit">Подтверить</button>
                    </form>
                {/if}
            {:else if (widget.type == "audio")}
                <div 
                    class={widgetsStatus[widget.order-1] ? "widget-container deleted" : "widget-container"}
                >
                    <AudioWidget 
                        link = {widget.data.link}
                    />
                    <button class="ch-btn" onclick={() => deleteWidget(widgets, widget)}>X</button>
                    <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, -1)}>^</button>
                    <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, 1)}>v</button>
                </div>
              {:else if (widget.type == "video")}
                <div 
                    class={widgetsStatus[widget.order-1] ? "widget-container deleted" : "widget-container"}
                >
                    <VideoWidget 
                        link = {widget.data.link}
                        platform = {widget.data.platform}
                    />
                    <button class="ch-btn" onclick={() => deleteWidget(widgets, widget)}>X</button>
                    <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, -1)}>^</button>
                    <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, 1)}>v</button>
                </div>
              {:else if (widget.type == "photo")}
                {#await getFilesURL(widget.id)}
                  <p></p>
                {:then photoUrls}
                  <div 
                      class={widgetsStatus[widget.order-1] ? "widget-container deleted" : "widget-container"}
                  >
                      <PhotoWidget
                          data = {photoUrls}
                      />
                      <button class="ch-btn" onclick={() => deleteWidget(widgets, widget)}>X</button>
                      <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, -1)}>^</button>
                      <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, 1)}>v</button>
                  </div>
                {/await}
              {:else if (widget.type == "todo")}
                <div 
                    class={widgetsStatus[widget.order-1] ? "widget-container deleted" : "widget-container"}
                >
                    <ToDoWidget 
                        title = {widget.data.title}
                        tasks = {widget.data.tasks}
                    />
                    <button class="ch-btn" onclick={() => deleteWidget(widgets, widget)}>X</button>
                    <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, -1)}>^</button>
                    <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, 1)}>v</button>
                </div>
              {:else if (widget.type == "progress_bar")}
                <div 
                    class={widgetsStatus[widget.order-1] ? "widget-container deleted" : "widget-container"}
                >
                    <ProgressWidget 
                        label = {widget.data.description}
                        progress = {widget.data.currentProgress}
                        max = {widget.data.maxProgress}
                    />
                    <button class="ch-btn" onclick={() => deleteWidget(widgets, widget)}>X</button>
                    <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, -1)}>^</button>
                    <button class="ch-btn" onclick={() => changeWidgetPostion(widgets, widget, 1)}>v</button>
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

<script lang="ts">
  import { invalidate } from '$app/navigation';
  import { onMount } from 'svelte';
  import { pb } from '$lib/index';
  import TextWidget from '$lib/components/widgets/TextWidget.svelte';
  import SurveyOption from '$lib/components/widgetConstructors/SurveyOption.svelte';
  import Task from '$lib/components/widgetConstructors/Task.svelte';
  import type { Widget } from '$lib/widgetTypes/widgetTypes';
  import type { RecordModel } from 'pocketbase';
  import type { RequestHandler } from '@sveltejs/kit';
  import AudioWidget from '$lib/components/widgets/AudioWidget.svelte';
  import { enhance } from '$app/forms';
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

    let widgets = $state<Array<RecordModel>>([]);
    let widgetsStatus = $state<Array<boolean>>([]);
    let widgetsChangeStatus = $state<Array<boolean>>([]);
    let addedWidgets = $state(0)

    async function getUserWidgets() {
        const widgetList = await pb.collection('widgets').getFullList({
            telegram_id: `${pb.authStore.record?.telegram_id}`,
            sort: `+order`
        })
        for (let i = 0; i < widgetList.length; i++) {
            widgetsStatus.push(false)
            widgetsChangeStatus.push(false)
        }
        return widgetList;
    }

    async function loadWidgets() {
        widgets = await getUserWidgets();
   }

    //widgets logic

    async function createWidget( { formData }: any ) {
        const formValues = Object.fromEntries(formData);
        await pb.collection("widgets").create({
            "telegram_id": pb.authStore.record?.telegram_id,
            "type": selectedWidget?.widgetType,
            "order" : widgetsStatus.length+1,
            "data": formValues
        })
        addedWidgets++;
        widgetsStatus.push(false)

        return { status: 200 };
    }

    async function updateWidgetsOrder(widgets:RecordModel[]) {
        for (let i = 1; i <= widgets.length; i++) {
            await pb.collection('widgets').update(widgets[i-1].id, {
                    "order": i
                }
            )
            console.log(widgets[i-1].id)
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

    async function updateWidget(widget: RecordModel) {
        return async ({ formData } : any) => {
            const formValues = Object.fromEntries(formData);
            await pb.collection('widgets').update(widget.id, {
            data: formValues
            });
        widgetsChangeStatus[widget.order-1] = false;
        return { status: 200 };
        };
    }

    async function changeWidgetPostion(widget: RecordModel, posChange: number) {
        const record = await pb.collection('widgets').getFirstListItem(`telegram_id ~ "${pb.authStore.record?.telegram_id}"`, {
            order: `${widget.order + posChange}`
        }).then(record => record)
        console.log(posChange)
        await pb.collection('widgets').update(record.id, {
            "order": record.order + (posChange * -1)
        })
        await pb.collection('widgets').update(widget.id, {
            "order": widget.order + posChange
        })
    }

    function showChangeWidgetField(widget: RecordModel) {
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
      <textarea class="input-widget-text" id="input-widget-text" name="text"></textarea>
    {/if}
    <button class="submit-btn" type="submit">Добавить виджет</button>
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
                    <button onclick={() => deleteWidget(widgets, widget)}>X</button>
                    <button onclick={() => changeWidgetPostion(widget, -1)}>^</button>
                    <button onclick={() => changeWidgetPostion(widget, 1)}>v</button>
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
                    <button onclick={() => deleteWidget(widgets, widget)}>X</button>
                    <button onclick={() => changeWidgetPostion(widget, -1)}>^</button>
                    <button onclick={() => changeWidgetPostion(widget, 1)}>v</button>
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

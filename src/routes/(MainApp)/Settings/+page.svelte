<script lang="ts">
    import SurveyOption from '../../../lib/components/widgetConstructors/SurveyOption.svelte';
    import Task from '../../../lib/components/widgetConstructors/Task.svelte'

    let { widgets } = $props();

    let tasksNumbers: Array<{id: number}> = $state([{id: 1}]);
    let surveyOptionsNumbers: Array<{id: number}> = $state([{id: 1}]);

    let widgetTypes = [
        {widgetType: "audio", widgetDisplay: "Аудио"}, 
        {widgetType: "video", widgetDisplay: "Видео"},
        {widgetType: "photo", widgetDisplay: "Фото"},
        {widgetType: "todo", widgetDisplay: "Списко заданий"},
        {widgetType: "progress_bar", widgetDisplay: "Отслеживание прогресса"},
        {widgetType: "geo", widgetDisplay: "Геолокаия"},
        {widgetType: "social_media", widgetDisplay: "Социальная сеть"},
        {widgetType: "steam_game", widgetDisplay: "Статистика игры в Steam"},
        {widgetType: "sticker", widgetDisplay: "Стикер"},
        {widgetType: "survey", widgetDisplay: "Опрос"},
        {widgetType: "text", widgetDisplay: "Текст"}
    ];

    let selectedWidget = $state<(typeof widgetTypes)[number]>();

    function removeTask(id: number) {
        tasksNumbers = tasksNumbers.filter(component => component.id !== id);
        for (let i = 0; i< tasksNumbers.length; i++) {
            tasksNumbers[i].id = i + 1
        }
    };
    function removeOption(id: number) {
        surveyOptionsNumbers = surveyOptionsNumbers.filter(component => component.id !== id);
        for (let i = 0; i< surveyOptionsNumbers.length; i++) {
            surveyOptionsNumbers[i].id = i + 1
        }
    };
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
        <br>
        {#if (selectedWidget?.widgetType == "audio")}
            <p>Ссылка на sc</p>
            <input type="text">
        {/if}
        {#if (selectedWidget?.widgetType == "video")}
            <p>Ссылка на видео</p>
            <input type="text">
        {/if}
        {#if (selectedWidget?.widgetType == "photo")}
            <p>Фото</p>
            <input type="file">
        {/if}
        {#if (selectedWidget?.widgetType == "todo")}
            {#each tasksNumbers as task (task.id)}
                <p>{task.id}</p>
                <Task onRemove={() => removeTask(task.id)}/>
            {/each}
            <button onclick={() => tasksNumbers = [...tasksNumbers, {id: tasksNumbers.length + 1}]} id="taskCreateButton">Добавить Цель</button>
        {/if}
        {#if (selectedWidget?.widgetType == "progress_bar")}
            <p>Название цели</p>
            <input type="text">
            Текущий прогресс
            <input type="text">
            Максимальный прогресс
            <input type="text">
        {/if}
        {#if (selectedWidget?.widgetType == "social_media")}
            <p>Ссылка на соц. сеть</p>
            <input type="text">
        {/if}
        {#if (selectedWidget?.widgetType == "steam_game")}
            <p>Ссылка на профиль Steam</p>
            <input type="text">
        {/if}
        {#if (selectedWidget?.widgetType == "sticker")}
            <p>Координаты</p>
            <input type="text">
        {/if}
        {#if (selectedWidget?.widgetType == "survey")}
            <p>Название опроса</p>
            <input type="text">
            {#each surveyOptionsNumbers as task (task.id)}
                <SurveyOption onRemove={() => removeOption(task.id)}/>
            {/each}
            <button onclick={() => surveyOptionsNumbers = [...surveyOptionsNumbers, {id: surveyOptionsNumbers.length + 1}]} id="taskCreateButton">Добавить вариант ответа</button>
        {/if}
        {#if (selectedWidget?.widgetType == "text")}
            <p>Текст</p>
            <input type="text">
        {/if}
    </form>
</main>

<style>
    .widget-form {
        display: flex;
        flex-direction: column;
        max-width: 300px;
    }
</style>
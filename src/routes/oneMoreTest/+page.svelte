<script>
    let show = false;
</script>

<button on:click={() => show = true}>Открыть модальное окно</button>

{#if show}
  <div class="modal">
    <div class="modal-content">
      <h2>Модальное окно</h2>
      <p>Тут какой-то текст...</p>
      <button on:click={() => show = false}>Закрыть</button>
    </div>
  </div>
{/if}

<style>
    .modal {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background: rgba(0, 0, 0, 0.5);
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .modal-content {
        background: white;
        padding: 20px;
        border-radius: 10px;
        text-align: center;
    }
</style>

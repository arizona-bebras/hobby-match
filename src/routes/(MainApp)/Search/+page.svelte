<script lang="ts">
  let { data } = $props();
  console.log(data);
</script>

<main>
  <form method="post">
    <p>Ваш пароль</p>
    <input name="password" type="password" />
    <button type="submit">Отправить</button>
  </form>
</main>

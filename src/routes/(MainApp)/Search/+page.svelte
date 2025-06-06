<script lang="ts">
  let { data } = $props();
  console.log(data);
  async function getUserGames() {
    const games = await fetch('/api/steam?link=1231');
    console.log(games);
  }
  getUserGames();
</script>

<main>
  <form method="post">
    <p>Ваш пароль</p>
    <input name="password" type="password" />
    <button type="submit">Отправить</button>
  </form>
</main>

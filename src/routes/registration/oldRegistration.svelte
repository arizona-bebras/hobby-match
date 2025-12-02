<script lang="ts">
  import { pb } from '$lib';
  import { goto } from '$app/navigation';

  let data = {
    location: '',
    birth_date: '',
  };
  async function handleSubmit() {
    //@ts-ignore
    const record = await pb
      .collection('users')
      .update(pb.authStore.record.id, data);
    data = {
      location: '',
      birth_date: '',
    };
    console.log('Создана запись:', record);
    goto('/profile');
  }
</script>

<main>
  <form on:submit|preventDefault={handleSubmit}>
    Ваш город
    <input
      type="text"
      id="location"
      name="location"
      bind:value={data.location}
      required
    />
    <br />
    Дата рождения
    <input
      type="date"
      id="birth_date"
      name="birth_date"
      bind:value={data.birth_date}
      required
    />
    <br />
    <button type="submit">Отправить</button>
  </form>
</main>

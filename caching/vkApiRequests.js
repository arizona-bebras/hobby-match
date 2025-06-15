export async function getPageInfo(key, handle) {
    const response = await fetch(`https://api.vk.ru/method/users.get?user_ids=${handle}&fields=followers_count&access_token=${key}&v=5.199`)
        .then((result) => result.json());
    console.log(response.response[0])
    return {
      name: response.response[0].first_name + ' ' + response.response[0].last_name,
      followers: response.response[0].followers_count,
    };
}
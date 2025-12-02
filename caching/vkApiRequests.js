export async function getPageInfo(key, handle) {
  const user = await fetch(
    `https://api.vk.com/method/users.get?user_ids=${handle}&fields=followers_count&access_token=${key}&v=5.199`,
  ).then((result) => result.json());
  if (user?.response?.[0]) {
    return {
      name: user.response[0].first_name + ' ' + user.response[0].last_name,
      followers: user.response[0].followers_count,
    };
  }

  const groups = await fetch(
    `https://api.vk.com/method/groups.getById?group_id=${handle}&fields=members_count&access_token=${key}&v=5.199`,
  ).then((result) => result.json());
  console.log('groups', groups);

  if (groups?.response?.groups?.[0]) {
    return {
      name: groups.response.groups[0].name,
      followers: groups.response.groups[0].members_count,
    };
  }
}

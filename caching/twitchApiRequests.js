export async function getTwitchChannelInfo(key, clientId, handle) {
    const headers = {
      'Client-Id': clientId,
      'Authorization': `Bearer ${key}`,
    };
    let info = await fetch(`https://api.twitch.tv/helix/users?login=${handle}`, {
      method: 'GET',
      headers: headers,
    }).then((result) => result.json());
    console.log(info)
    if (!info.data || info.data.length !== 1) return null;
    const id = info.data[0].id;
    const title = info.data[0]['display_name'];
    let followers = await fetch(`https://api.twitch.tv/helix/channels/followers?broadcaster_id=${id}`, {
      method: 'GET',
      headers: headers,
    }).then((result) => result.json());
    console.log(followers)
    return followers.total;
}
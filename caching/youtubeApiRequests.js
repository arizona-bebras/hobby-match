export async function getYoutubeChannelInfo(key, handle) {
  let response = await fetch(
    `https://www.googleapis.com/youtube/v3/channels?part=snippet,statistics&forHandle=${handle}&key=${key}`,
  ).then((result) => result.json());
  console.log(response.items[0].statistics)
  return response.items[0].statistics.subscriberCount;
}
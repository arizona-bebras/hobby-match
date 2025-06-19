export async function getYoutubeChannelInfo(key, handle, isId = false) {
  let param;
  if (isId) {
    param = `id=${handle}`;
  } else {
    param = `forHandle=${handle}`;
  }

  let response = await fetch(
    `https://www.googleapis.com/youtube/v3/channels?part=snippet,statistics&key=${key}&${param}`,
  ).then((result) => result.json());
  return {
    subscribers: response.items[0].statistics.subscriberCount,
    title: response.items[0].snippet.title,
  };
}

//import { YOUTUBE_API_KEY } from '$env/static/private';
import type { RequestHandler } from './$types';

const YOUTUBE_API_KEY = 'AIzaSyBKIqw2as6ud5sfg2TOYl__NooTh7jJeuQ';

export const POST: RequestHandler = async ({ request }) => {
  const requestData = await request.json();
  const username = requestData.username;

  const searchResponse = await fetch(
    `https://www.googleapis.com/youtube/v3/search?part=snippet&q=@${username}&type=channel&key=${YOUTUBE_API_KEY}`,
  );
  const searchData = await searchResponse.json();
  const channelId = searchData.items[0].id.channelId;

  const data = await fetch(
    `https://www.googleapis.com/youtube/v3/channels?part=snippet,statistics&id=${channelId}&key=${YOUTUBE_API_KEY}`,
  ).then((result) => result.json());
  const subs = data.items[0].statistics.subscriberCount;
  return new Response(
    JSON.stringify({
      success: true,
      subscribers: parseInt(subs),
    }),
    {
      headers: { 'Content-Type': 'application/json' },
    },
  );
};

import type { PageLoad } from './$types';
import client from '$lib/api/client';

export const load: PageLoad = async ({ params, url }) => {
  const userId = url.pathname.split('/').at(-1);
  if (parseInt(userId)) {
    const pageData = await client.GET('/api/pages/{page_id}', {
      params: {
        path: {
          page_id: userId,
        },
      },
    });
    return {
      pageData,
    };
  }
  return {
    pageData: undefined,
  };
};

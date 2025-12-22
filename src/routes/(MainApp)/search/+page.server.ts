import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';

export const actions = {
  default: async ({ request }) => {
    const formData = await request.formData();
    const password = formData.get('password');
    //const data = Object.fromEntries(formData);
    console.log('Получены данные:', password);

    return { success: true };
  },
};

export async function load({ url }): Promise<{ page?: PageData }> {
  if (!url.searchParams.has('opened')) return {};
  return {};
  // try {
  //   const {
  //     // @ts-expect-error it exists
  //     expand: { interests },
  //     ...page
  //   } = await pb.collection('users').getOne(url.searchParams.get('opened')!, {
  //     fields: 'expand,id,miniapp_name,age,gender,location,user_photo,user_info',
  //     expand: 'interests',
  //   });
  //
  //   return {
  //     // @ts-expect-error the types are fine
  //     page: {
  //       ...page,
  //       interests,
  //       widgets: await pb
  //         .collection('widgets')
  //         .getFullList({ filter: `user = '${page.id}'` }), // TODO: impersonate user
  //     },
  //   };
  // } catch (_) {
  //   return {};
  // }
}

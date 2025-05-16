import type { PageServerLoad } from '../../../../.svelte-kit/types/src/routes/registration/$types';

export const actions = {
  default: async ({ request }) => {
    const formData = await request.formData();
    const password = formData.get('password');
    //const data = Object.fromEntries(formData);
    console.log('Получены данные:', password);

    return { success: true };
  },
};
export const load: PageServerLoad = async () => {
  return {
    b: 2,
  };
};

export const actions = {
  default: async ({ request }) => {
    const formData = await request.formData();

    const data = Object.fromEntries(formData);
    console.log('Получены данные:', data);

    return { success: true };
  },
};

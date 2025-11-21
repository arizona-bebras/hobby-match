import type { PageServerLoad } from './$types.js';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { textSchema } from '$lib/components/editor/schemes/textSheme';
import { informationSchema } from '$lib/components/registration/information/InformationFormShema';

export const load: PageServerLoad = async () => {
  return {
    textForm: await superValidate(zod(textSchema)),
  };
};
export const actions = {
  default: async (event) => {
    const someData = await superValidate(event, zod(informationSchema));
    console.log(someData);
    return {
      someData,
    };
  },
  // const formData = await request.formData();
  //
  // const data = Object.fromEntries(formData);
  // console.log('Получены данные:', data, textSchema.parse(data));
  //
  // return { success: true };
};

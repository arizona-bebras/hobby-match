import { type Actions, fail, redirect, type RequestEvent } from '@sveltejs/kit';
import { defaults, superValidate } from 'sveltekit-superforms';
import { photoSchema } from '$lib/components/registration/photo/PhotoFormShema';
import { informationSchema } from '$lib/components/registration/information/InformationFormShema';
import { interestsScheme } from '$lib/components/registration/interests/InterestsFormShema';
import { zod4 } from 'sveltekit-superforms/adapters';
import type { PageServerLoad } from './$types';
import { z } from 'zod';
import { updateUserData } from '$lib/components/registration/updateUserData';
//import { goto } from '$app/navigation';

export const load: PageServerLoad = async ({ locals }) => {
  return {
    information: await superValidate(zod4(informationSchema)), //{defaults=locals.}
    photo: await superValidate(zod4(photoSchema)),
    interests: await superValidate(zod4(interestsScheme)),
  };
};
export const actions: Actions = {
  information: async (event: RequestEvent) => {
    //const photo = await superValidate(event, zod(photoSchema));
    const information = await superValidate(event, zod4(informationSchema));
    console.log(information);
    if (!information.valid) {
      return fail(400, {
        information,
      });
    }
    return {
      information,
    };
  },
  photo: async (event: RequestEvent) => {
    //const photo = await superValidate(event, zod(photoSchema));
    const photo = await superValidate(event, zod4(photoSchema));
    console.log(photo.data.user_photo);
    if (!photo.valid) {
      return fail(400, {
        photo,
      });
    }
    return {
      //photo,
    };
  },
  interests: async (event: RequestEvent) => {
    //const photo = await superValidate(event, zod(photoSchema));
    const interests = await superValidate(event, zod4(interestsScheme));
    console.log(interests);
    if (!interests.valid) {
      return fail(400, {
        interests,
      });
    }
    await updateUserData();
    redirect(303, '/profile');
    return {
      interests,
    };
  },
};

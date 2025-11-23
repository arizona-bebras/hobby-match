import { pb } from '$lib';
import type { PageLoad } from './$types';
import type { WidgetWithService } from '$lib/components/widgetConstructors/widgetsConstructor';
import type { PageData } from '$lib/questionnaireTypes/questionnaireTypes';

export const load: PageLoad = async ({ data, depends }) => {
  depends('user:widgets');
  console.log('RELOADED!!!');
  if (pb.authStore.isValid) {
    const { textForm } = data;
    //
    // const collection = await pb.collections.getOne('widgets');
    // console.log(collection.schema);
    const widgetsRecords = await pb.collection('widgets').getFullList({
      sort: `+order`,
      requestKey: null,
    });

    // @ts-expect-error because
    const widgets: WidgetWithService[] = widgetsRecords.map((record) => ({
      widget: record,
      deleteStatus: false,
      changeStatus: false,
      additionalData: {},
    }));
    const pageData: PageData = {
      id: pb.authStore.record!.id,
      user_photo: pb.authStore.record!.user_photo,
      user_info: pb.authStore.record!.user_info,
      miniapp_name: pb.authStore.record!.miniapp_name,
      birth_date: pb.authStore.record!.birth_date,
      age: pb.authStore.record!.age,
      location: pb.authStore.record!.location,
      interests: await pb
        .collection('users')
        .getOne(pb.authStore.record!.id, {
          fields: 'expand',
          expand: 'interests',
          requestKey: null,
        })
        .then((result) => {
          return result.expand?.interests;
        }),
      widgets: structuredClone(widgets),
      textForm,
    };
    console.log(pageData);
    return pageData;
  }
  throw new Error('Data validation error');
};

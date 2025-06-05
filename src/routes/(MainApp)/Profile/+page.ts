import { pb } from '$lib';
import type { PageLoad } from './$types';
import type { WidgetWithService } from '$lib/components/widgetConstructors/widgetsConstructor';

export const load: PageLoad = async ({ data, depends }) => {
  depends('user:widgets');
  if (pb.authStore.isValid) {
    const { textForm } = data;
    //
    // const collection = await pb.collections.getOne('widgets');
    // console.log(collection.schema);
    const widgetsRecords = await pb.collection('widgets').getFullList({
      sort: `+order`,
    });

    // @ts-expect-error because
    const widgets: WidgetWithService[] = widgetsRecords.map((record) => ({
      widget: record,
      deleteStatus: false,
      changeStatus: false,
      additionalData: {},
    }));
    return {
      user_photo: pb.authStore.record!.user_photo,
      miniapp_name: pb.authStore.record!.miniapp_name,
      age: pb.authStore.record!.age,
      location: pb.authStore.record!.location,
      interests: pb.authStore.record!.interests,
      widgets: structuredClone(widgets),
      textForm,
    };
  }
  throw new Error('Data validation error');
};

import * as wid from '$lib/widgetTypes/widgetTypes';
import {
  diefinePlatofrm,
  getVideoPlatform,
  getImageDimensions,
  getUsernameFromUrl,
  pb,
} from '$lib/index';

export interface AdditionalData {
  socialMeidaData?: number;
}

export interface WidgetWithService {
  widget: wid.Widget;
  deleteStatus: boolean;
  changeStatus: boolean;
  additionalData: AdditionalData;
}

export async function createWidget({
  formData,
}: {
  formData: FormData;
}): Promise<void> {
  const uploadedFiles = formData.getAll('files') as File[];

  const formValues = Object.fromEntries(formData);

  const widgetType = formValues.type;

  console.log(widgetType);

  delete formValues.files;

  let widgetData:
    | wid.Audio
    | wid.Video
    | wid.Photos
    | wid.Todo
    | wid.ProgressBar
    | wid.Geo
    | wid.SocialMediaLink
    | wid.SteamGame
    | wid.Sticker
    | wid.Survey
    | wid.Text
    | wid.Empty = {
    type: 'empty',
  };

  switch (widgetType) {
    case 'text':
      widgetData = {
        type: 'text',
        text: formValues.text as string,
      };
      break;
    case 'audio':
      widgetData = {
        type: 'audio',
        link: '',
      };
      if (diefinePlatofrm(formValues.link as string) == 'SoundCloud') {
        widgetData.link = formValues.link as string;
      } else {
        throw 'wrong link';
      }
      break;

    case 'video':
      widgetData = {
        type: 'video',
        link: '',
        platform: 'Undefined',
      };
      widgetData.link = formValues.link as string;
      widgetData.platform = getVideoPlatform(
        diefinePlatofrm(formValues.link as string),
      );
      break;

    case 'photo':
      widgetData = {
        type: 'photo',
        photos: [],
      };
      for (const file of uploadedFiles) {
        const { width, height } = await getImageDimensions(file);
        widgetData.photos.push({
          name: file.name,
          width: width,
          height: height,
          size: file.size,
        });
      }
      break;

    case 'todo':
      widgetData = {
        type: 'todo',
        title: '',
        tasks: [],
      };
      for (const [key, value] of Object.entries(formValues)) {
        if (key.includes('task')) {
          const task: wid.Task = {
            order: parseInt(key.slice(4)),
            description: value as string,
            isCompleted: false,
          };
          widgetData.tasks.push(task);
          delete formValues[key];
        }
      }
      widgetData.tasks.sort((a: wid.Task, b: wid.Task) => a.order - b.order);
      break;
    case 'progress_bar':
      widgetData = {
        type: 'progress_bar',
        description: formValues.description as string,
        currentProgress: parseInt(formValues.currentProgress as string),
        maxProgress: parseInt(formValues.maxProgress as string),
      };
      break;
    case 'social_media':
      widgetData = {
        type: 'social_media',
        platform: 'Undefined',
        username: '',
        link: '',
        subscribers: 0,
      };
      if (diefinePlatofrm(formValues.link as string) == 'Youtube') {
        widgetData.link = formValues.link as string;
        widgetData.username = getUsernameFromUrl(formValues.link as string);
        widgetData.platform = diefinePlatofrm(formValues.link as string);
      }
      if (diefinePlatofrm(formValues.link as string) == 'Steam') {
        widgetData.username = '###';
        widgetData.platform = diefinePlatofrm(formValues.link as string);
        console.log(formValues);
      }
  }

  console.log(widgetData);
  await pb.collection('widgets').create({
    telegram_id: pb.authStore.model?.telegram_id,
    order: 1000,
    files: uploadedFiles,
    data: widgetData,
  });
}

export async function updateWidgetsOrder(widgets: WidgetWithService[]) {
  for (let i = 1; i <= widgets.length; i++) {
    await pb.collection('widgets').update(widgets[i - 1].widget.id, {
      order: i,
    });
    console.log(widgets[i - 1].widget.id);
  }
}

export async function deleteWidget(
  widgets: WidgetWithService[],
  widget: WidgetWithService,
): Promise<void> {
  await pb.collection('widgets').delete(widget.widget.id);
  widget.deleteStatus = true;
  console.log(widget.deleteStatus);
  const filtred = [];
  for (let i = 0; i < widgets.length; i++) {
    if (!widgets[i].deleteStatus) {
      filtred.push(widgets[i]);
    }
  }
  widgets = filtred;
  console.log(widgets);
  updateWidgetsOrder(filtred);
}

export async function updateWidget(widget: WidgetWithService) {
  return async ({ formData }: { formData: FormData }) => {
    const formValues = Object.fromEntries(formData);
    formValues.type = widget.widget.data.type;
    await pb.collection('widgets').update(widget.widget.id, {
      data: formValues,
    });
    widget.changeStatus = false;
  };
}

export async function changeWidgetPostion(
  widgets: WidgetWithService[],
  widget: WidgetWithService,
  posChange: number,
): Promise<void> {
  const record = await pb
    .collection('widgets')
    .getFullList({
      filter: `telegram_id = "${pb.authStore.model?.telegram_id}" && order = "${widget.widget.order + posChange}"`,
    })
    .then((record) => record);

  const widgetOrder = widget.widget.order;
  const recordOrder = record[0].order;
  await pb.collection('widgets').update(record[0].id, {
    order: recordOrder + posChange * -1,
  });
  await pb.collection('widgets').update(widget.widget.id, {
    order: widgetOrder + posChange,
  });
  const temp = widgets[recordOrder - 1];
  widgets[recordOrder - 1].widget.order += posChange * -1;
  widgets[recordOrder - 1] = widgets[widgetOrder - 1];
  widgets[widgetOrder - 1].widget.order += posChange;
  widgets[widgetOrder - 1] = temp;
}

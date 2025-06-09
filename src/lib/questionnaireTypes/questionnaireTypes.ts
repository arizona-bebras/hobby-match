import type { WidgetWithService } from '$lib/components/widgetConstructors/widgetsConstructor';

export type PageData = {
  user_photo: string;
  miniapp_name: string;
  age: number;
  location: string;
  interests: { expand: { interests: object[] } };
  widgets: WidgetWithService[];
  textForm: any;
};

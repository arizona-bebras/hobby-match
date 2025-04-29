import PocketBase from 'pocketbase';
import type { Widget } from '$lib/widgetTypes/widgetTypes';
export const pb = new PocketBase('http://127.0.0.1:8090');

export const swapElements = (
  widgets: Widget[],
  recordOrder: number,
  widgetOrder: number,
  posChange: number,
) => {
  const temp = widgets[recordOrder - 1];
  widgets[recordOrder - 1].order += posChange * -1;
  widgets[recordOrder - 1] = widgets[widgetOrder - 1];
  widgets[widgetOrder - 1].order += posChange;
  widgets[widgetOrder - 1] = temp;
};

export function diefinePlatofrm(url: string): string {
  const patterns = {
    youtube: /(youtube\.com|youtu\.be)/i,
    vimeo: /vimeo\.com/i,
    tiktok: /tiktok\.com/i,
    instagram: /instagram\.com/i,
    twitter: /(twitter\.com|x\.com)/i,
    rutube: /rutube\.ru/i,
    vk: /(vk\.com|vkontakte\.ru)/i,
    dzen: /dzen\.ru/i,
    ok: /ok\.ru/i,
  };

  for (const [platform, regex] of Object.entries(patterns)) {
    if (regex.test(url)) {
      return platform;
    }
  }

  return 'unknown';
}

export async function getImageDimensions(file: File) {
  return new Promise((resolve) => {
    const img = new Image();
    const url = URL.createObjectURL(file);

    img.onload = () => {
      resolve({ width: img.width, height: img.height });
      URL.revokeObjectURL(url);
    };

    img.src = url;
  });
}

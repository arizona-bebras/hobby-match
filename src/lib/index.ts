import PocketBase, { type RecordModel } from 'pocketbase';
import type { Widget } from '$lib/widgetTypes/widgetTypes';
import type { AdditionalData } from './components/widgetConstructors/widgetsConstructor';
import { PUBLIC_PB_ENDPOINT } from '$env/static/public';

export const pb = new PocketBase(PUBLIC_PB_ENDPOINT ?? 'http://127.0.0.1:8090');

export function convertRecordToWidget(record: RecordModel): Widget {
  const widget: Widget = {
    id: record.id,
    order: record.order,
    data: record.data,
    additionalData: record.additionalData,
  };
  return widget;
}

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

export function definePlatofrm(url: string): string {
  const patterns = {
    Youtube: /(youtube\.com|youtu\.be)/i,
    Tiktok: /tiktok\.com/i,
    Twitter: /(twitter\.com|x\.com)/i,
    Rutube: /rutube\.ru/i,
    VK: /(vk\.com|vkontakte\.ru)/i,
    Steam: /(steamcommunity\.com|store\.steampowered\.com)/i,
    SoundCloud: /soundcloud\.com/i,
  };

  for (const [platform, regex] of Object.entries(patterns)) {
    if (regex.test(url)) {
      return platform;
    }
  }

  return 'unknown';
}

export function getVideoPlatform(
  platform: string,
): 'YouTube' | 'Rutube' | 'Tiktok' | 'Undefined' {
  if (platform == 'YouTube') return 'YouTube';
  if (platform == 'Rutube') return 'Rutube';
  if (platform == 'Tiktok') return 'Tiktok';
  return 'Undefined';
}

export async function getImageDimensions(
  file: File,
): Promise<{ width: number; height: number }> {
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

export function getSocialMediaData(
  widget: Widget,
): Promise<AdditionalData> | undefined {
  const platform = widget.data.platform;
  const link = widget.data.link;
  if (platform == 'Youtube') {
    return getYoutubeChannelStats(getUsernameFromUrl(link));
  }
  if (platform == 'Steam') {
    return getSteamData(link);
  }
}

export async function getYoutubeChannelStats(
  username: string,
): Promise<AdditionalData> {
  const res = await fetch('/api/youtube', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username: username }),
  });
  const data = await res.json();
  const youtubeStats: AdditionalData = {
    socialMediaData: data.subscribers,
  };
  console.log(youtubeStats);
  return youtubeStats;
}

export async function getSteamData(link: string) {
  const res = await fetch('/api/steam', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ link: link }),
  });
  const data = await res.json();
  const steamStats: AdditionalData = {
    socialMediaData: data.player_level,
    steamUsername: data.player_name,
  };
  console.log(steamStats);
  return steamStats;
}

export function getUsernameFromUrl(url: string): string {
  const platform = definePlatofrm(url);
  if (platform == 'Youtube') return url.slice(25);
  return '';
}

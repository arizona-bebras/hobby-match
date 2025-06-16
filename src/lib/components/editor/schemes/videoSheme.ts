import { z } from 'zod';

export const videoSchema = z.object({
  link: z.string().refine((value) => getYouTubeVideoId(value) !== null, {
    message: 'Неверный формат ссылки.',
  }),
  // .regex(
  //   /^(https?:\/\/(?:www\.)?(?:youtube\.com\/embed\/[a-zA-Z0-9_-]{11}(?:\?si=[a-zA-Z0-9_-]+)?|rutube\.ru\/play\/embed\/[a-zA-Z0-9]+(?:\/)?|www\.tiktok\.com\/player\/v1\/\d+))$/,
  //   'Неверный формат ссылки. Поддерживается YouTube, Rutube и TikTok',
  // ),
});

export type FormSchema = typeof videoSchema;

export function getYouTubeVideoId(url: string): string | null {
  const patterns = [
    // YouTube
    /(?:youtube\.com\/(?:watch\?v=|shorts\/|live\/)|youtu\.be\/)([a-zA-Z0-9_-]{11})/,
    // RuTube
    /rutube\.ru\/video\/([a-f0-9]{32})/,
    // TikTok
    /tiktok\.com\/(?:@[^/]+\/video|v)\/(\d{19})/,
  ];
  for (const pattern of patterns) {
    const match = url.match(pattern);
    if (match) return match[1];
  }
  return null;
}

import { z } from 'zod';

export const audioScheme = z.object({
  link: z
    .union(
      [z.string()
      .regex(
        /soundcloud\.com/i,
        'Неверный формат ссылки',
      ),
      z.string()
      .regex(
        /music.yandex\.ru\/album\/[0-9]+\/track\/[0-9]+$/i,
        'Неверный формат ссылки',
      ),
      z.string()
      .regex(
        /open.spotify\.com\/track\/[\s\S]*$/i,
        'Неверный формат ссылки',
      )]
    )
});

export type FormSchema = typeof audioScheme;

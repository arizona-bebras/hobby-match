import { z } from 'zod';

export const audioScheme = z.object({
  link: z
    .string()
    .regex(
      /soundcloud\.com/i,
      'Неверный формат ссылки',
    ),
});

export type FormSchema = typeof audioScheme;

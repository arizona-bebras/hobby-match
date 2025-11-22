import { z } from 'zod';

export const postScheme = z.object({
  link: z
    .string()
    .regex(
      /^https:\/\/t\.me\/[a-zA-Z0-9_]{1,32}\/\d{1,10}$/,
      'Неверный формат ссылки',
    ),
});

export type FormSchema = typeof postScheme;

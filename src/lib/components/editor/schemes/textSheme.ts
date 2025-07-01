import { z } from 'zod';

export const textSchema = z.object({
  text: z
    .string()
    .min(1, 'Поле не должно быть пустым')
    .max(1000, 'Текст не должен превышать 1000 символов'),
});

export type FormSchema = typeof textSchema;

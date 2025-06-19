import { z } from 'zod';

export const textSchema = z.object({
  text: z.string().min(1, 'Поле не должно быть пустым'),
});

export type FormSchema = typeof textSchema;

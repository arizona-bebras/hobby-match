import { z } from 'zod';

export const createSchema = z.object({
  title: z.string().min(2).max(50),
  photo: z
    .instanceof(File, { message: 'Please upload a file.' })
    .refine(
      (f) => f.size < 5_242_880,
      'Размер аватарки должен быть не больше 5 МБ',
    ),
  description: z.string().min(2).max(50),
  secret_word: z.string().min(5).max(50),
});

export type FormSchema = typeof createSchema;

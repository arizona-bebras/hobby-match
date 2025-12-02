import { z } from 'zod';

export const photoSchema = z.object({
  user_photo: z
    .instanceof(File, { message: 'Please upload a file.' })
    .refine(
      (f) => f.size < 5_242_880,
      'Размер аватарки должен быть не больше 5 МБ',
    )
    .or(z.string().url()),
});

export type FormSchema = typeof photoSchema;

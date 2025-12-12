import { z } from 'zod';

export const createSchema = z.object({
  title: z
    .string()
    .min(2, { error: 'Длинна поля должна быть более 2 символов' })
    .max(50, { error: 'Длинна поля должна быть не более 50 символов' }),
  photo: z
    .instanceof(File, { message: 'Please upload a file.' })
    .refine(
      (f) => f.size < 5_242_880,
      'Размер аватарки должен быть не больше 5 МБ',
    )
    .optional(),
  description: z
    .string()
    .min(2, { error: 'Длинна поля должна быть более 2 символов' })
    .max(50, { error: 'Длинна поля должна быть не более 50 символов' }),
});

export type FormSchema = typeof createSchema;

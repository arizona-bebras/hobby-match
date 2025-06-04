import { z } from 'zod';

export const progressScheme = z
  .object({
    description: z.string().min(1),
    currentProgress: z.number(),
    maxProgress: z.number(),
  })
  .refine((data) => data.currentProgress <= data.maxProgress, {
    message: 'Текущий прогресс не может быть больше максимального',
    // path: ['currentProgress'],
  });

export type FormSchema = typeof progressScheme;

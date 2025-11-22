import { z } from 'zod';

export const progressScheme = z
  .object({
    description: z.string().min(1, 'Поле не может быть пустым'),
    currentProgress: z.number({ message: 'Неверный формат ввода' }),
    maxProgress: z.number({ message: 'Неверный формат ввода' }),
  })
  .refine((data) => data.currentProgress <= data.maxProgress, {
    message: 'Текущий прогресс не может быть больше максимального',
    // path: ['currentProgress'],
  });

export type FormSchema = typeof progressScheme;

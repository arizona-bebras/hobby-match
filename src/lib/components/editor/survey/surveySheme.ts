import { z } from 'zod';

export const surveyScheme = z.object({
  question: z.string().min(1, 'Поле не должно быть пустым'),
  options: z.array(z.string().min(1, 'Поле не должно быть пустым')).min(2),
});

export type FormSchema = typeof surveyScheme;

import { z } from 'zod';

export const surveyScheme = z.object({
  question: z.string().min(1),
  options: z.array(z.string().min(1)).min(2),
});

export type FormSchema = typeof surveyScheme;

import { z } from 'zod';

export const surveyScheme = z.object({
  question: z.string(),
  options: z.array(z.string()),
});

export type FormSchema = typeof surveyScheme;

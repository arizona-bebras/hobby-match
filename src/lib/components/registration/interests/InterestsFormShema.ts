import { z } from 'zod';

export const interestsScheme = z.object({
  interests: z.string().array().max(10).min(3),
});

export type FormSchema = typeof interestsScheme;

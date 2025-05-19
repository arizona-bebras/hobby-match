import { z } from 'zod';

export const videoSchema = z.object({
  link: z.string(),
});

export type FormSchema = typeof videoSchema;

import { z } from 'zod';

export const textSchema = z.object({
  text: z.string(),
});

export type FormSchema = typeof textSchema;

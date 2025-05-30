import { z } from 'zod';

export const imageScheme = z.object({
  files: z
    .instanceof(File, { message: 'Please upload a file.' })
    .refine((f) => f.size < 100_000, 'Max 100 kB upload size.')
    .array(),
});
export type FormSchema = typeof imageScheme;

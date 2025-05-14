import { z } from 'zod';

export const photoSchema = z.object({
  user_photo: z
    .instanceof(File, { message: 'Please upload a file.' })
    .or(z.string().url()),
});

export type FormSchema = typeof photoSchema;

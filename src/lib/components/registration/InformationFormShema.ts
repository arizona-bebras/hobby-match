import { z } from 'zod';

export const informationSchema = z.object({
  username: z.string().min(2).max(50),
  // @ts-expect-error sex
  gender: z.enum(['male', 'female']).default(undefined),
  dateOfBirth: z.string().datetime(),
  city: z.string().min(2).max(50),
  information: z.string().min(2).max(250),
});

export type FormSchema = typeof informationSchema;

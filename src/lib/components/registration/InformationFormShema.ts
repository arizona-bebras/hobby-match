import { z } from 'zod';

export const informationSchema = z.object({
  miniapp_name: z.string().min(2).max(50),
  // @ts-expect-error sex
  gender: z.enum(['male', 'female']),
  birth_date: z.string().datetime(),
  location: z.string().min(2).max(50),
  user_info: z.string().min(2).max(250),
});

export type FormSchema = typeof informationSchema;

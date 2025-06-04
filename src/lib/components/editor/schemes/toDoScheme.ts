import { z } from 'zod';

export const toDoScheme = z.object({
  title: z.string().min(1),
  tasks: z.array(z.string().min(1)).min(1),
});

export type FormSchema = typeof toDoScheme;

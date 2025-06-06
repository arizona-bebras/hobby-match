import { z } from 'zod';

export const toDoScheme = z.object({
  title: z.string().min(1),
  tasks: z
    .array(
      z.object({
        description: z.string().min(1),
        isCompleted: z.boolean(),
      }),
    )
    .min(1)
    .max(8),
});

export type FormSchema = typeof toDoScheme;

import { z } from 'zod';

export const toDoScheme = z.object({
  title: z.string().min(1, 'Поле не должно быть пустым'),
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

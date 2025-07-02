import { z } from 'zod';

export const reasons = ['Неприемлемый контент', 'Спам', 'Прочее'] as const;

export const reportSchema = z.object({
  reason: z.enum(reasons).default(reasons[0]),
  info: z.string(),
});

export type FormSchema = typeof reportSchema;

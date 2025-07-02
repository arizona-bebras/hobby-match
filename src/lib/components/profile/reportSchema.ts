import { z } from 'zod';

export const reasons = [
  'Неподобающий контент',
  'Оскорбления',
  'Ложная информация',
  'Фейковый профиль',
  'Мошенничество',
  'Спам',
  'Другое',
] as const;

export const reportSchema = z.object({
  reason: z.enum(reasons).default(reasons[0]),
  info: z.string(),
});

export type FormSchema = typeof reportSchema;

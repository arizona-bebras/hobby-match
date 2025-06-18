import { z } from 'zod';

export const informationSchema = z.object({
  miniapp_name: z
    .string()
    .min(2, 'Длинна должна быть больше 2-ух символов')
    .max(50, 'Длина не должна быть больше 50-ти символов'),
  // @ts-expect-error sex
  gender: z.enum(['male', 'female']),
  birth_date: z.string().datetime(),
  location: z
    .string()
    .min(1, 'Поле не должно быть пустым')
    .max(50, 'Длина не должна быть больше 50-ти символов'),
  user_info: z
    .string()
    .min(10, 'Длинна должна быть больше 10-и символов')
    .max(250, 'Длина не должна быть больше 250-ти символов'),
});

export type FormSchema = typeof informationSchema;

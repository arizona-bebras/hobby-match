import { z } from 'zod';

export const scheme = z
  .object({
    isStudent: z.boolean(),
    isWorker: z.boolean(),
    education: z
      .union([
        z.object({
          type: z.literal('school'),
          class: z
            .number()
            .positive({ error: 'Класс не может быть отрицательным' })
            .lte(12, { error: 'Класс не может быть больше 11' }),
          place: z.string({ error: 'Поле не должно быть пустым' }),
        }),
        z.object({
          type: z.literal('university'),
          educationStage: z.literal(['bachelor', 'master', 'postgraduate']),
          course: z
            .number()
            .positive({ error: 'Курс не может быть отрицательным' })
            .lte(7, { error: 'Курс не может быть больше 6' }),
          place: z.string({ error: 'Поле не должно быть пустым' }),
        }),
      ])

      .optional(),
    work: z
      .object({
        company: z.string(),
        job_title: z.string(),
        experience: z
          .number()
          .min(0, { error: 'Стаж не может быть отрицательным' })
          .optional()
          .nullable(),
      })
      .optional(),
    // text: z
    //   .string()
    //   .min(1, 'Поле не должно быть пустым')
    //   .max(1000, 'Текст не должен превышать 1000 символов'),
  })
  .superRefine((data, ctx) => {
    if (!data.isStudent && !data.isWorker) {
      ctx.addIssue({
        code: 'custom',
        message: 'Выберите хотя бы одну занятость',
      });
    }
  });

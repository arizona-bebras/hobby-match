import { z } from 'zod';

export const scheme = z
  .object({
    isStudent: z.boolean(),
    isWorker: z.boolean(),
    education: z
      .union([
        z.object({
          type: z.literal('school'),
          class: z.number(),
        }),
        z.object({
          type: z.literal('university'),
          educationStage: z.literal(['bachelor', 'master', 'postgraduate']),
          course: z.number(),
        }),
      ])

      .optional(),
    work: z
      .object({
        company: z.string(),
        job_title: z.string(),
        experience: z.number(),
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

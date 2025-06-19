import { z } from 'zod';

export const gameScheme = z.object({
  accountLink: z
    .string()
    .regex(
      /^(?:https:\/\/)?steamcommunity\.com\/((?:id)|(?:profiles))\/(\w+)/gm,
      'Неверный формат ссылки',
    ),
  gameId: z.string(),
});

export type FormSchema = typeof gameScheme;

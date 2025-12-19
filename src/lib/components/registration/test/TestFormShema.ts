import { z } from 'zod';

export const testSchema = z.object({
  socialScale: z.number(),
  actionScale: z.number(),
  organizationScale: z.number(),
  interactionRoleScale: z.number(),
  focusDepthScale: z.number(),
});

export type FormSchema = typeof testSchema;

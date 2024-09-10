import { z } from 'zod'

export const EffectConfigSchema = z
  .object({
    name: z.string(),
    type: z.enum(['slider', 'screenSelector']),
    min: z.number().optional(),
    max: z.number().optional(),
    label_max: z.number().optional(),
    step: z.number().optional(),
    unit: z.enum(['s', 'ms', '%']).optional()
  })
  .transform((config) => ({
    ...config,
    value: config.min
  }))

export const EffectSettingsSchema = z.object({
  effect: z.enum(['Twinkle', 'Color Cycle', 'Strobe', 'Flame', 'Visualizer', 'Theater']),
  description: z.string(),
  configs: z.record(EffectConfigSchema)
})

export const EffectTypeSchema = z.object({
  id: z.number(),
  settings: EffectSettingsSchema
})

export const EffectsArraySchema = z.array(EffectTypeSchema)

export type EffectType = z.infer<typeof EffectTypeSchema>
export type EffectSettings = z.infer<typeof EffectSettingsSchema>
export type EffectConfig = z.infer<typeof EffectConfigSchema>

import { type EffectType, EffectsArraySchema } from '@renderer/schema/effectSchema'

export const parseEffectData = (data: unknown): EffectType[] => {
  return EffectsArraySchema.parse(data)
}

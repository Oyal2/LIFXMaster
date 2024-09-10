import { Label } from '@radix-ui/react-label'
import { Card, CardContent } from '@renderer/components/ui/card'
import { Dialog, DialogTrigger } from '@renderer/components/ui/dialog'
import { IconAdjustmentsAlt, IconPlayerStop, IconPlayerPlay } from '@tabler/icons-react'
import EffectSettingsModal from './EffectSettings'
import type { EffectSettings, EffectType } from '@renderer/schema/effectSchema'

interface CardEffectProps {
  isActive: boolean
  startEffect: (cardId: number, settings: EffectSettings) => void
  stopEffect: (cardId: number) => void
  updateCardSettings: (cardId: number, newSettings: EffectSettings) => void
  effect: EffectType
}

function CardEffect(props: CardEffectProps): JSX.Element {
  const { isActive, startEffect, stopEffect, updateCardSettings, effect } = props
  const handleStart = (): void => {
    if (isActive) {
      stopEffect(effect.id)
    } else {
      startEffect(effect.id, effect.settings)
    }
  }
  return (
    <Dialog>
      <Card className="bg-[#1A1C48] text-[#9499C3] border-[#9499C3] rounded-lg w-[24vw] h-[18vh] ">
        <CardContent className="h-full pb-0">
          <div className="w-full h-full flex py-3">
            <DialogTrigger asChild disabled={isActive}>
              <IconAdjustmentsAlt stroke={2} className="cursor-pointer" size={30} />
            </DialogTrigger>
            <div className="w-full flex justify-between items-center">
              <Label className="font-bold text-xl flex-grow text-center">
                {effect.settings.effect}
              </Label>
              {isActive ? (
                <IconPlayerStop
                  stroke={2}
                  size={40}
                  className="cursor-pointer"
                  onClick={handleStart}
                />
              ) : (
                <IconPlayerPlay
                  stroke={2}
                  size={40}
                  className="cursor-pointer"
                  onClick={handleStart}
                />
              )}
            </div>
          </div>
          <EffectSettingsModal effect={effect} updateCardSettings={updateCardSettings} />
        </CardContent>
      </Card>
    </Dialog>
  )
}

export default CardEffect

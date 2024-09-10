import { useEffect, useState } from 'react'
import CardEffect from './CardEffect'
import effectsData from '../../../data/effectsData.json'
import { chunkArray } from '@renderer/utils/chunkArray'
import { parseEffectData } from '@renderer/utils/parseEffectData'
import type { EffectSettings, EffectType } from '@renderer/schema/effectSchema'
import { effectCommands } from '@renderer/utils/effectCommands'
import useLightsStore, { type Light } from '@renderer/hooks/useLightsStore'
import { checkIfAllLightsExists } from '@renderer/utils/lightUtils'

interface EffectsMenuProps {
  lights: Light[]
  locationId: string | undefined
  groupId: string | undefined
}

function EffectsMenu(props: EffectsMenuProps): JSX.Element {
  const { lights, groupId, locationId } = props
  const [activeEffectId, setActiveEffectId] = useState<number | null>(null)
  const [effectCards, setEffectCards] = useState<EffectType[]>([])
  const { updateGroupLights, updateLight } = useLightsStore()

  useEffect(() => {
    const parsedData = parseEffectData(effectsData)
    setEffectCards(parsedData)
  }, [])

  useEffect(() => {
    // Check if any light has an active effect
    const findLight = lights.find((light) => light.effect !== undefined)
    if (findLight?.effect) {
      const activeEffect = findLight?.effect
      // Find the effect card with the matching effect name
      const activeCard = effectCards.find((card) => card.settings.effect === activeEffect)
      if (activeCard) {
        setActiveEffectId(activeCard.id)
      }
    } else {
      setActiveEffectId(null)
    }
  }, [lights, effectCards])

  const startEffect = (cardId: number, settings: EffectSettings): void => {
    if (locationId === undefined || groupId === undefined) {
      return
    }
    if (activeEffectId !== null) {
      stopEffect(activeEffectId)
    }
    setActiveEffectId(cardId)

    const command = effectCommands[settings.effect]
    if (command) {
      const isAllLights = checkIfAllLightsExists(lights)
      const updatedLights: Light[] = [...lights]
      command(settings, isAllLights ? updatedLights.slice(1) : updatedLights, true)
      for (let i = 0; i < updatedLights.length; i++) {
        if (!isAllLights) {
          updateLight(updatedLights[i].target.toString(), { effect: settings.effect })
        } else {
          updatedLights[i].effect = settings.effect
        }
      }
      if (isAllLights) {
        updateGroupLights(locationId, groupId, updatedLights)
      }
    }
  }

  const stopEffect = (cardId: number): void => {
    const effectCard = effectCards.find((x) => x.id === cardId)
    if (effectCard === undefined || locationId === undefined || groupId === undefined) {
      return
    }

    const command = effectCommands[effectCard.settings.effect]
    if (command) {
      const isAllLights = checkIfAllLightsExists(lights)
      const updatedLights: Light[] = [...lights]

      command(effectCard.settings, isAllLights ? updatedLights.slice(1) : updatedLights, false)
      for (let i = 0; i < updatedLights.length; i++) {
        if (!isAllLights) {
          updateLight(updatedLights[i].target.toString(), {
            effect: undefined
          })
        } else {
          updatedLights[i].effect = undefined
        }
      }
      if (isAllLights) {
        updateGroupLights(locationId, groupId, updatedLights)
      }
    }
    setActiveEffectId(null)
  }

  const updateCardSettings = (cardId: number, newSettings: EffectSettings): void => {
    setEffectCards((prevEffectCards) =>
      prevEffectCards.map((effectCard) =>
        effectCard.id === cardId ? { ...effectCard, settings: newSettings } : effectCard
      )
    )
  }

  return (
    <div className="px-12 py-8 h-[65vh] overflow-auto scrollbar scrollbar-thumb-rounded-full scrollbar-thumb-slate-700 scrollbar-thin">
      {chunkArray(effectCards, 2).map((effectCardPair, index) => (
        <div key={`${index}_${effectCardPair}`} className="flex mb-8 space-x-8">
          {effectCardPair.map((effectCard) => (
            <CardEffect
              key={effectCard.id}
              effect={effectCard}
              isActive={effectCard.id === activeEffectId}
              startEffect={startEffect}
              updateCardSettings={updateCardSettings}
              stopEffect={stopEffect}
            />
          ))}
          {effectCardPair.length === 1 && <div className="w-[24vw]" />}
        </div>
      ))}
    </div>
  )
}

export default EffectsMenu

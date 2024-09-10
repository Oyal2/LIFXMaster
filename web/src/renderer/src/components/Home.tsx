import { dotPulse } from 'ldrs'
import RoomComponent from './RoomComponent'
import useLightsStore, { type LocationStore } from '@renderer/hooks/useLightsStore'
import { checkAnyLightOn } from '@renderer/utils/lightUtils'
import { useMemo } from 'react'
import type { SetLocationLabelRequest } from 'src/proto/proto/message_service'
import { HsvaToHex } from '@renderer/utils/hsv'
import { toast } from 'sonner'

interface HomeProp {
  isLoading: boolean
}

function Home(props: HomeProp): JSX.Element {
  const { isLoading } = props
  dotPulse.register()
  const { locations, updateLocations, updateLocation } = useLightsStore()
  const locationKeys = Object.entries(locations)
  const handleAllLightsClick = async (): Promise<void> => {
    const powerRequest: { [key: string]: boolean } = {}
    const newLocations: {
      [x: string]: LocationStore
    } = { ...locations }

    for (const locationKey in newLocations) {
      const location = newLocations[locationKey]
      const newGroups = { ...location.groups }
      newLocations[locationKey].isOn = false

      for (const groupKey in newGroups) {
        const group = newGroups[groupKey]
        const newLights = [...group.lights]
        newGroups[groupKey].isOn = false

        for (let i = 0; i < newLights.length; i++) {
          const light = newLights[i]
          if (light.target) {
            const oldLevel = light.power?.level || 0
            const newLevel = oldLevel > 0 ? 0 : 100
            const turnOn = oldLevel === 0

            const lightTarget = light.target.toString()
            if (lightTarget !== '-1') {
              powerRequest[lightTarget] = turnOn
            }
            if (turnOn) {
              group.isOn = turnOn
              location.isOn = turnOn
            }
            newLights[i] = {
              ...light,
              power: {
                level: newLevel
              }
            }
          }
        }
        newGroups[groupKey].lights = newLights
      }
      location.groups = newGroups
      newLocations[locationKey] = location
    }

    try {
      await window.electron.ipcRenderer.invoke('set-power', powerRequest)
    } catch (error) {
      console.error("Failed to set location's power:", error)
      toast.error(`Failed to set location's power: ${error}`)
    }
    updateLocations(newLocations)
  }

  const updateLocationLabel = async (locationId: string, newLabel: string): Promise<void> => {
    if (locationId === 'all_lights') return
    const setLocationLabelRequest: SetLocationLabelRequest = {
      locationID: locationId,
      newLabel: newLabel
    }

    try {
      await window.electron.ipcRenderer.invoke('set-location-label', setLocationLabelRequest)
    } catch (error) {
      console.error("Failed to set location's label:", error)
      toast.error(`Failed to set location's label: ${error}`)
    }
    updateLocation(locationId, { label: newLabel })
  }

  const allLightsStatus = useMemo(() => checkAnyLightOn(locations), [locations])

  return (
    <div>
      <h1 className="text-white text-3xl mb-8 font-semibold text-center">Current Locations</h1>
      {!isLoading ? (
        locationKeys.length > 0 && (
          <div className="grid grid-cols-5 gap-4">
            <RoomComponent
              key={'all_lights'}
              lightbulb={{
                name: 'All Lights',
                color: 'white',
                percentage: 100,
                isOn: allLightsStatus
              }}
              uuid={'all_lights'}
              roomType="all_lights"
              onClick={handleAllLightsClick}
            />
            {locationKeys.map(([key, location]) => (
              <RoomComponent
                key={key}
                lightbulb={{
                  name: location.label,
                  color: HsvaToHex(location.color),
                  percentage: location.color.v,
                  isOn: location.isOn
                }}
                uuid={key}
                onUpdateLabel={updateLocationLabel}
                roomType="locations"
              />
            ))}
          </div>
        )
      ) : (
        <div className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
          <l-dot-pulse size="75" speed="1.3" color="white" />
        </div>
      )}
    </div>
  )
}

export default Home

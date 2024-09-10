import { dotPulse } from 'ldrs'
import RoomComponent from './RoomComponent'
import { useNavigate, useParams } from 'react-router-dom'
import { IconArrowLeft } from '@tabler/icons-react'
import useLightsStore, { type LocationStore } from '@renderer/hooks/useLightsStore'
import { backClick } from '@renderer/utils/backClick'
import { checkAnyGroupLightOn } from '@renderer/utils/lightUtils'
import { useMemo } from 'react'
import type { SetGroupLabelRequest } from 'src/proto/proto/message_service'
import { HsvaToHex } from '@renderer/utils/hsv'
import { toast } from 'sonner'

interface GroupProp {
  isLoading: boolean
}

function GroupDisplay(props: GroupProp): JSX.Element {
  const { isLoading } = props
  const { locationId } = useParams<{ locationId: string }>()
  const { locations, updateLocations, updateGroup } = useLightsStore()
  const navigate = useNavigate()

  dotPulse.register()

  const handleAllLightsClick = async (): Promise<void> => {
    const powerRequest: { [key: string]: boolean } = {}
    const newLocations: {
      [x: string]: LocationStore
    } = { ...locations }

    if (locationId === undefined) {
      return
    }

    const location = newLocations[locationId]
    const newGroups = { ...location.groups }
    newLocations[locationId].isOn = false

    for (const groupKey in newGroups) {
      const group = newGroups[groupKey]
      newGroups[groupKey].isOn = false

      const newLights = [...group.lights]
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
    newLocations[locationId] = location

    try {
      await window.electron.ipcRenderer.invoke('set-power', powerRequest)
    } catch (error) {
      console.error("Failed to set group's power:", error);
      toast.error(`Failed to set group's power: ${error}`);
    }
    updateLocations(newLocations)
  }

  const updateGroupLabel = async (groupId: string, newLabel: string): Promise<void> => {
    if (locationId === undefined || groupId === 'all_lights') {
      return
    }
    const setGroupLabelRequest: SetGroupLabelRequest = {
      groupID: groupId,
      newLabel: newLabel
    }
    try {
      await window.electron.ipcRenderer.invoke('set-group-label', setGroupLabelRequest)
    } catch (error) {
      console.error("Failed to set group's label:", error)
      toast.error(`Failed to set group's label: ${error}`);
    }
    updateGroup(locationId, groupId, { label: newLabel })
  }

  const allLightsStatus = useMemo(() => {
    if (locationId === undefined) return false
    return checkAnyGroupLightOn(locations[locationId].groups)
  }, [locations, locationId])

  return (
    <div>
      <div className="flex justify-between items-center mb-8">
        <div className="flex justify-start items-center flex-1 h-2">
          <IconArrowLeft
            stroke={2}
            size={42}
            color="white"
            className="cursor-pointer"
            onClick={() => backClick(navigate, '/')}
          />
        </div>
        {locationId && (
          <h1 className="text-white text-3xl font-semibold text-center flex-1">{`${locations[locationId].label}'s Groups`}</h1>
        )}
        <div className="cursor-pointer flex-1" />
      </div>
      {!isLoading ? (
        locationId && (
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
            {Object.entries(locations[locationId].groups).map(([key, group]) => (
              <RoomComponent
                key={key}
                lightbulb={{
                  name: group.label,
                  color: HsvaToHex(group.color),
                  percentage: group.color.v,
                  isOn: group.isOn
                }}
                uuid={key}
                roomType="groups"
                onUpdateLabel={updateGroupLabel}
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

export default GroupDisplay

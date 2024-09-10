import { dotPulse } from 'ldrs'
import { useNavigate, useParams } from 'react-router-dom'

import { DataTable } from '../table/LightsTable'
import type { HSBK } from '../../../../proto/proto/message_service'
import LightMenu from '../LightMenu'
import ColorMenu from './ColorMenu'
import React from 'react'
import type { HsvaColor } from '@uiw/react-color'
import InfoMenu from './InfoMenu'
import EffectsMenu from './Effects/EffectsMenu'
import { HsvaToHsbk } from '../../utils/hsv'
import { Button } from '../ui/button'
import { ArrowLeft } from 'lucide-react'
import useLightsStore, {
  type Light,
  type GroupStore,
  type LocationStore
} from '@renderer/hooks/useLightsStore'
import { backClick } from '@renderer/utils/backClick'
import { getLights } from '@renderer/utils/lightUtils'
import { toast } from 'sonner'

function LightsHubDisplay(): JSX.Element {
  const { locationId, groupId } = useParams<{
    locationId: string
    groupId: string
  }>()
  dotPulse.register()
  const [rowSelection, setRowSelection] = React.useState({})
  const [menuSelection, setMenuSelection] = React.useState<number>(0)
  const locations = useLightsStore((state) => state.locations)
  const updateGroupLights = useLightsStore((state) => state.updateGroupLights)
  const navigate = useNavigate()
  const menu = ['Color', 'Effects', 'Info']
  const isRowSelected = Object.keys(rowSelection).length > 0
  const lights = getLights(locations, locationId, groupId)

  const updateLightColor = async (
    locationId: string | undefined,
    groupId: string | undefined,
    newColor: HsvaColor,
    targetLight?: bigint
  ): Promise<void> => {
    if (locationId === undefined || groupId === undefined) return
    const colorRequest: { [key: string]: HSBK } = {}
    const location: LocationStore = locations[locationId]
    if (!location) return

    const group: GroupStore = location.groups[groupId]
    if (!group) return

    const updatedLights = group.lights.map((light: Light) => {
      if (targetLight === undefined || light.target === targetLight) {
        const updatedLight = { ...light, color: newColor, effect: undefined }
        if (updatedLight.target.toString() !== '-1') {
          const hsbk = HsvaToHsbk(updatedLight.color)
          colorRequest[updatedLight.target.toString()] = hsbk
        }
        return updatedLight
      }
      return light
    })
    try {
      await window.electron.ipcRenderer.invoke('set-color', colorRequest)
    } catch (error) {
      console.error('Failed to set color:', error)
      toast.error(`Failed to set color: ${error}`)
    }
    updateGroupLights(locationId, groupId, updatedLights)
  }

  function getSelectableLights(index: number): Light[] {
    if (index === 0) {
      return lights
    }

    return [lights[index]]
  }

  function getMenu(index: number): JSX.Element {
    switch (index) {
      case 0: {
        const index = Number(Object.keys(rowSelection)[0])
        return (
          <ColorMenu
            hsva={lights[index].color}
            updateColor={(newColor) => {
              if (Number(Object.keys(rowSelection)[0]) === 0) {
                updateLightColor(locationId, groupId, newColor)
              } else {
                updateLightColor(locationId, groupId, newColor, lights[index].target)
              }
            }}
          />
        )
      }
      case 1: {
        const index = Number(Object.keys(rowSelection)[0])
        return (
          <EffectsMenu
            lights={getSelectableLights(index)}
            locationId={locationId}
            groupId={groupId}
          />
        )
      }
      case 2: {
        const index = Number(Object.keys(rowSelection)[0])
        return <InfoMenu light={lights[index]} />
      }
      default:
        return <></>
    }
  }

  return (
    <div className="flex flex-col bg-[#0B0E1F] rounded">
      <div className="w-full justify-center">
        <div className={`grid  ${isRowSelected ? 'grid-cols-2' : 'grid-cols-1'} gap-6 px-6 py-6`}>
          <div className="bg-[#191933] rounded-xl py-4 px-4">
            <Button
              className="cursor-pointer text-[#9499C3] mb-2 px-1 font-semibold rounded-md  "
              onClick={() => backClick(navigate, `/locations/${locationId}`)}
            >
              <ArrowLeft className="mr-2 h-5 w-5" />
              Back to Groups
            </Button>
            <DataTable
              data={lights}
              locationId={locationId}
              groupId={groupId}
              rowSelection={rowSelection}
              setRowSelection={setRowSelection}
            />
          </div>
          {isRowSelected && (
            <div className="bg-[#191933] rounded-xl px-2">
              <LightMenu
                items={
                  Number(Object.keys(rowSelection)[0]) === 0
                    ? menu.filter((x) => x !== 'Info')
                    : menu
                }
                setMenuSelection={setMenuSelection}
                menuSelection={menuSelection}
              />
              {getMenu(menuSelection)}
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default LightsHubDisplay

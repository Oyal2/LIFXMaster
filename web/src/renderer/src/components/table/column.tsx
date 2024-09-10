import { hsvaToHex } from '@uiw/react-color'
import type { Label, SetDeviceLabelRequest } from 'src/proto/proto/message_service'
import WifiIcon from '../WifiIcon'
import { Checkbox } from '@radix-ui/react-checkbox'
import type { ColumnDef, FilterFn } from '@tanstack/table-core'
import { MdOutlineEdit } from 'react-icons/md'
import GlowingLightBulb from '../LightBulb'
import useEditableLabel from '@renderer/hooks/useEditableLabel'
import { FaCheck } from 'react-icons/fa'
import { TiCancel } from 'react-icons/ti'
import { Button } from '../ui/button'
import { Input } from '../ui/input'
import useLightsStore, { type Light } from '@renderer/hooks/useLightsStore'
import { useState } from 'react'
import { toast } from 'sonner'

interface LightColumns {
  lights: Light[]
  handleIndividualLightClick: (light: Light) => Promise<void>
  handleAllLightsClick: (lights: Light[]) => Promise<void>
}

function GetLightColumns(lightColumns: LightColumns): ColumnDef<Light>[] {
  const { lights, handleIndividualLightClick, handleAllLightsClick } = lightColumns
  const { updateLight } = useLightsStore()
  const updateDeviceLabel = async (lightId: string, newLabel: string): Promise<void> => {
    const setDeviceLabelRequest: SetDeviceLabelRequest = {
      deviceID: BigInt(lightId),
      newLabel: newLabel
    }
    try {
      await window.electron.ipcRenderer.invoke('set-device-label', setDeviceLabelRequest)
    } catch (error) {
      console.error("Failed to set device's label:", error)
      toast.error(`Failed to set device's label: ${error}`)
    }
    updateLight(lightId, {
      label: {
        label: newLabel
      }
    })
  }

  const filterLabel: FilterFn<Light> = (row, columnId: string, filterValue: string) => {
    if (!filterValue.length) return true

    const rowValue = row.getValue<Label>(columnId).label.toLowerCase() || ''
    return rowValue.includes(filterValue.toLowerCase())
  }

  return [
    {
      id: 'select',
      accessorKey: 'select',
      cell: ({ row }) => (
        <Checkbox
          checked={row.getIsSelected()}
          onCheckedChange={(value) => row.toggleSelected(!!value)}
          aria-label="Select row"
        />
      ),
      enableHiding: true
    },
    {
      accessorKey: 'color',
      header: () => <div className="flex justify-center">Turn On/Off</div>,
      cell: ({ row }): JSX.Element => {
        const handleClick =
          row.original.target.toString() === '-1'
            ? (): Promise<void> => handleAllLightsClick(lights)
            : (): Promise<void> => handleIndividualLightClick(row.original)
        const light = lights.find((x) => x.target === row.original.target)
        return (
          <div className="flex justify-center">
            <GlowingLightBulb
              glowColor={
                light && light.power !== undefined && light.power.level > 0
                  ? hsvaToHex(row.getValue('color'))
                  : '#000000'
              }
              imageType={'Lightbulb'}
              size={40}
              className="cursor-pointer"
              onClick={(e) => {
                e.stopPropagation()
                handleClick()
              }}
            />
          </div>
        )
      }
    },
    {
      accessorKey: 'label',
      header: () => <div className="flex justify-center">Label</div>,
      cell: ({ row }): JSX.Element => {
        const [isHovered, setIsHovered] = useState<boolean>(false)
        const { view, value, handleKeyUp, handleChange, handleCancel, handleSave, switchToInput } =
          useEditableLabel({
            initialName: row.original.label?.label || '',
            onUpdateLabel: updateDeviceLabel,
            uuid: row.original.target.toString()
          })

        const renderView = (): JSX.Element => {
          return view === 'label' ? (
            <div
              className="flex items-center"
              onMouseEnter={() => setIsHovered(true && row.original.target.toString() !== '-1')}
              onMouseLeave={() => setIsHovered(false && row.original.target.toString() !== '-1')}
            >
              <div className="flex-grow" />
              <div className="flex-shrink-0">{row.original.label?.label}</div>
              <div
                className="flex-grow flex justify-end"
                onClick={(e) => {
                  if (row.original.target.toString() === '-1') {
                    return
                  }
                  e.stopPropagation()
                  e.preventDefault()
                  switchToInput()
                }}
                onKeyUp={(e) => {
                  if (row.original.target.toString() === '-1') {
                    return
                  }
                  e.stopPropagation()
                  e.preventDefault()
                  switchToInput()
                }}
              >
                <MdOutlineEdit
                  size={20}
                  className={`${isHovered && view === 'label' ? 'show' : 'opacity-0'}`}
                />
              </div>
            </div>
          ) : (
            <div className="mt-2 w-full">
              <div className="flex items-center space-x-2">
                <Input
                  className="flex-grow text-white text-lg py-2 h-10 bg-gray-800 border border-gray-700 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  value={value}
                  onClick={(e) => {
                    e.stopPropagation()
                    e.preventDefault()
                  }}
                  onKeyUp={handleKeyUp}
                  onChange={handleChange}
                />
                <Button
                  className="p-2 h-10 bg-red-600 hover:bg-red-700 transition-colors duration-200 rounded-md focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-opacity-50"
                  onClick={handleCancel}
                >
                  <TiCancel size={24} className="text-white" />
                </Button>
                <Button
                  className="p-2 h-10 bg-green-600 hover:bg-green-700 transition-colors duration-200 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-opacity-50"
                  onClick={handleSave}
                >
                  <FaCheck size={24} className="text-white" />
                </Button>
              </div>
            </div>
          )
        }
        return renderView()
      },
      filterFn: filterLabel
    },
    {
      accessorKey: 'address',
      header: () => <div className="flex justify-center">Address</div>,
      cell: ({ row }) => <div className="flex justify-center">{row.original.address}</div>
    },
    {
      accessorKey: 'port',
      header: () => <div className="flex justify-center">Port</div>,
      cell: ({ row }) => <div className="flex justify-center">{row.original.port}</div>
    },
    {
      accessorKey: 'firmware',
      header: () => <div className="flex justify-center">Firmware</div>,
      cell: ({ row }) => (
        <div className="flex justify-center">
          {row.original.firmware !== undefined &&
            `v${row.original.firmware?.versionMajor}.${row.original.firmware?.versionMinor}`}
        </div>
      )
    },
    {
      accessorKey: 'signal',
      header: () => <div className="flex justify-center">Signal</div>,
      cell: ({ row }): JSX.Element => {
        return (
          <div className="flex justify-center">
            {row.original.target.toString() !== '-1' && (
              <WifiIcon signalStrength={Number(row.original.wifi?.info?.signal)} />
            )}
          </div>
        )
      }
    }
  ]
}

export default GetLightColumns

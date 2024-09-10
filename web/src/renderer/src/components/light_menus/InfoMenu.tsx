import type { Light } from '@renderer/hooks/useLightsStore'
import WifiIcon from '../WifiIcon'

interface InfoMenuProps {
  light: Light
}

function InfoMenu(props: InfoMenuProps): JSX.Element {
  const { light } = props

  const booleanDisplay = (input: boolean | undefined): JSX.Element => {
    return (
      <span className={`${input ? 'text-green-500' : 'text-red-500'}`}>
        {input ? 'true' : 'false'}
      </span>
    )
  }
  return (
    <div className="px-12 py-4 text-[#9499C3]">
      <div className="grid grid-rows-3 gap-8">
        <div className="flex flex-col">
          <h1 className="text-2xl font-semibold text-center mb-6">General</h1>
          <div className="grid grid-cols-2 gap-16">
            <div>
              <div className="flex justify-between">
                <span className="flex font-medium items-center">Product:</span>
                <span>{light.product?.name}</span>
              </div>
              <div className="flex justify-between">
                <span className="flex font-medium items-center">Power:</span>
                <span>{`${((light.power?.level || 0) / 65535) * 100}%`}</span>
              </div>
            </div>
            <div>
              <div className="flex justify-between">
                <span className="flex font-medium items-center">Color:</span>
                <span>Green</span>
              </div>
              <div className="flex justify-between">
                <span className="flex font-medium items-center">Firmware:</span>
                <span>{`v${light.firmware?.versionMajor}.${light.firmware?.versionMinor}`}</span>
              </div>
            </div>
          </div>
        </div>
        <div className="flex flex-col">
          <div className="grid grid-cols-2 gap-16">
            <div>
              <h1 className="text-2xl font-semibold text-center mb-6">Network</h1>
              <div className="flex justify-between">
                <span className="flex font-medium items-center">Address:</span>
                <span>{`${light.address}:${light.port}`}</span>
              </div>
              <div className="flex justify-between">
                <span className="flex font-medium items-center items-center">Signal:</span>
                <WifiIcon signalStrength={light.wifi?.info?.signal || 0} />
              </div>
            </div>
            <div>
              <h1 className="text-2xl font-semibold text-center mb-6">Geo location</h1>
              <div className="flex justify-between">
                <span className="flex font-medium items-center">Location:</span>
                <span>{light.location?.label}</span>
              </div>
              <div className="flex justify-between">
                <span className="flex font-medium items-center">Group:</span>
                <span>{`${light.group?.label}`}</span>
              </div>
              <div className="flex justify-between">
                <span className="flex font-medium items-center">Label</span>
                <span>{light.label?.label}</span>
              </div>
            </div>
          </div>
        </div>
        <div className="flex flex-col">
          <div>
            <h1 className="text-2xl font-semibold text-center mb-6">Features</h1>
            <div className="grid grid-cols-2 gap-16">
              <div>
                <div className="flex justify-between">
                  <span className="flex font-medium items-center">Hev:</span>
                  <span>{booleanDisplay(light.product?.features?.hev)}</span>
                </div>
                <div className="flex justify-between">
                  <span className="flex font-medium items-center">Color:</span>
                  <span>{booleanDisplay(light.product?.features?.color)}</span>
                </div>
                <div className="flex justify-between">
                  <span className="flex font-medium items-center">Multizone:</span>
                  <span>{booleanDisplay(light.product?.features?.multizone)}</span>
                </div>
              </div>
              <div>
                <div className="flex justify-between">
                  <span className="flex font-medium items-center">Matrix:</span>
                  <span>{booleanDisplay(light.product?.features?.matrix)}</span>
                </div>
                <div className="flex justify-between">
                  <span className="flex font-medium items-center">Relays:</span>
                  <span>{booleanDisplay(light.product?.features?.relays)}</span>
                </div>
                <div className="flex justify-between">
                  <span className="flex font-medium items-center">Infrared:</span>
                  <span>{booleanDisplay(light.product?.features?.infrared)}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default InfoMenu

import Saturation from '@uiw/react-color-saturation'
import { type HsvaColor, hsvaToHex } from '@uiw/react-color'
import ShadeSlider from '@uiw/react-color-shade-slider'
import Swatch from '@uiw/react-color-swatch'
import Hue from '@uiw/react-color-hue'

interface ColorMenuProps {
  hsva: HsvaColor
  updateColor: (newColor: HsvaColor) => void
}

function ColorMenu(props: ColorMenuProps): JSX.Element {
  const { hsva, updateColor } = props
  return (
    <div className="px-32 py-8">
      <div className="flex justify-center mb-8">
        <Saturation
          style={{ width: '30vw', height: '30vh' }}
          hsva={hsva}
          onChange={(newColor: HsvaColor) => {
            updateColor(newColor)
          }}
        />
      </div>
      <div className="mb-6">
        <h1 className="text-white text-lg mb-2">{`Brightness: ${hsva.v.toFixed(0)}%`}</h1>
        <ShadeSlider
          hsva={hsva}
          onChange={(newShade) => {
            updateColor({
              ...hsva,
              v: newShade.v
            })
          }}
        />
      </div>
      <div className="mb-6">
        <Hue
          hue={hsva.h}
          onChange={(newHue) => {
            updateColor({
              ...hsva,
              h: newHue.h
            })
          }}
        />
      </div>
      <div>
        <Swatch
          color={hsvaToHex(hsva)}
          colors={[
            '#D0021B',
            '#F5A623',
            '#f8e61b',
            '#8B572A',
            '#7ED321',
            '#417505',
            '#BD10E0',
            '#9013FE',
            '#4A90E2',
            '#50E3C2',
            '#B8E986',
            '#FF6347',
            '#FFD700',
            '#FF4500',
            '#ADFF2F',
            '#32CD32',
            '#00FA9A',
            '#00CED1',
            '#1E90FF',
            '#7B68EE',
            '#9932CC',
            '#8A2BE2',
            '#FF1493',
            '#DC143C',
            '#FF69B4',
            '#FFB6C1',
            '#20B2AA',
            '#48D1CC',
            '#5F9EA0',
            '#6495ED',
            '#6A5ACD',
            '#708090',
            '#778899',
            '#B0C4DE',
            '#4682B4',
            '#D2691E',
            '#A0522D',
            '#CD5C5C',
            '#BC8F8F',
            '#F4A460',
            '#DAA520',
            '#B8860B',
            '#D3D3D3',
            '#A9A9A9',
            '#C0C0C0',
            '#B22222',
            '#DC143C',
            '#FF7F50',
            '#FF8C00',
            '#FFA07A',
            '#FFDAB9',
            '#FFDEAD',
            '#EEE8AA',
            '#F0E68C',
            '#BDB76B',
            '#556B2F',
            '#808000',
            '#6B8E23',
            '#9ACD32',
            '#ADFF2F',
            '#7FFF00',
            '#7CFC00',
            '#000000',
            '#4A4A4A',
            '#9B9B9B',
            '#FFFFFF'
          ]}
          onChange={(newColor: HsvaColor) => {
            updateColor(newColor)
          }}
        />
      </div>
    </div>
  )
}

export default ColorMenu

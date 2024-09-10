import { IconPower } from '@tabler/icons-react'
import { IconBulb } from '@tabler/icons-react'

type LightbulbImageType = 'Lightbulb' | 'Power'

interface GlowingLightBulbProps {
  imageType: LightbulbImageType
  size: number
  glowColor: string
  className?: string
  onClick?: React.MouseEventHandler<SVGSVGElement>
}

function GlowingLightBulb(props: GlowingLightBulbProps): JSX.Element {
  const { glowColor, imageType, size, className, onClick } = props
  const dropShadow = `
    drop-shadow(0 0 1.6px ${glowColor}) 
    drop-shadow(0 0 10px ${glowColor}) 
    drop-shadow(0 0 40px ${glowColor}) 
  `
  return (
    <>
      {imageType === 'Lightbulb' ? (
        <IconBulb
          stroke={2}
          size={size}
          color={glowColor}
          style={{ filter: glowColor === '#000000' ? '' : dropShadow }}
          className={className}
          onClick={onClick}
        />
      ) : (
        <IconPower
          stroke={2}
          size={size}
          color={glowColor}
          style={{ filter: glowColor === '#000000' ? '' : dropShadow }}
          className={className}
          onClick={onClick}
        />
      )}
    </>
  )
}

export default GlowingLightBulb

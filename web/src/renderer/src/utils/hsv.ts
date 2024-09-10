import type { HsvaColor } from '@uiw/react-color'
import { HSBK } from 'src/proto/proto/message_service'

export function Uint16HsvToHex(hueUint16 = 0, saturationUint16 = 0, valueUint16 = 0): string {
  // Scale the incoming values
  const hueDegrees = (hueUint16 / 65535) * 360 // Scale hue to 0-360 degrees
  const saturation = saturationUint16 / 65535 // Scale saturation to 0-1
  const brightness = valueUint16 / 65535 // Scale brightness to 0-1

  return HexConvert(hueDegrees, saturation, brightness)
}

export function HsvaToHex(hsva: HsvaColor): string {
  const { h, s, v } = hsva
  return HexConvert(h, s / 100, v / 100)
}

export function HsvaEqual(hsvaA: HsvaColor, hsvaB: HsvaColor): boolean {
  return hsvaA.a === hsvaB.a && hsvaA.h === hsvaB.h && hsvaA.s === hsvaB.s && hsvaA.v === hsvaB.v
}

export function HsvaToHsbk(hsva: HsvaColor, kelvin = 3500): HSBK {
  const { h, s, v } = hsva

  const hue = Math.round((0x10000 * h) / 360) % 0x10000
  const saturation = Math.round((s / 100) * 65535)
  const brightness = Math.round((v / 100) * 65535)

  const kelvinClamped = Math.min(Math.max(kelvin, 1500), 9000)

  return {
    hue,
    saturation,
    brightness,
    kelvin: kelvinClamped
  }
}

function HexConvert(hue: number, saturation: number, brightness: number): string {
  let red = 0
  let green = 0
  let blue = 0

  const sector = Math.floor(hue / 60)
  const fractional = hue / 60 - sector
  const p = brightness * (1 - saturation)
  const q = brightness * (1 - fractional * saturation)
  const t = brightness * (1 - (1 - fractional) * saturation)

  switch (sector % 6) {
    case 0:
      red = brightness
      green = t
      blue = p
      break
    case 1:
      red = q
      green = brightness
      blue = p
      break
    case 2:
      red = p
      green = brightness
      blue = t
      break
    case 3:
      red = p
      green = q
      blue = brightness
      break
    case 4:
      red = t
      green = p
      blue = brightness
      break
    case 5:
      red = brightness
      green = p
      blue = q
      break
    default:
      throw new Error('Unexpected case in HSV to RGB conversion')
  }

  return `#${toHex(red)}${toHex(green)}${toHex(blue)}`
}

const toHex = (x: number): string => {
  const hex = Math.round(x * 255).toString(16)
  return hex.length === 1 ? `0${hex}` : hex
}

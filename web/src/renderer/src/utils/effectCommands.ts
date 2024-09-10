import type { EffectSettings } from '@renderer/schema/effectSchema'
import type {
  ColorCycleRequest,
  HSBK,
  StrobeRequest,
  TheaterRequest,
  TwinkleRequest,
  VisualizerRequest
} from 'src/proto/proto/message_service'
import { HsvaToHsbk } from './hsv'
import type { Light } from '@renderer/hooks/useLightsStore'
import { toast } from 'sonner'

export const effectCommands: Record<
  string,
  (settings: EffectSettings, lights: Light[], turnOn: boolean) => Promise<void>
> = {
  Twinkle: async (settings: EffectSettings, lights: Light[], turnOn: boolean) => {
    const intensity = settings.configs.intensity.value || 0 / 100
    const speed = (settings.configs.speed.value || 0) / (settings.configs.speed.step || 1)
    const deviceColors: { [key: string]: HSBK } = {}
    for (const light of lights) {
      deviceColors[light.target.toString()] = HsvaToHsbk(light.color)
    }

    const request: TwinkleRequest = {
      deviceColors,
      speed,
      intensity,
      turnOn
    }
    try {
      await window.electron.ipcRenderer.invoke('twinkle', request)
    } catch (error) {
      console.error('Failed to set twinkle effect:', error)
      toast.error(`Failed to set twinkle effect: ${error}`)
    }
  },
  'Color Cycle': async (settings: EffectSettings, lights: Light[], turnOn: boolean) => {
    const request: ColorCycleRequest = {
      deviceIDs: lights.map((light) => light.target),
      speed: settings.configs.speed.value || 0,
      turnOn
    }
    try {
      await window.electron.ipcRenderer.invoke('color-cycle', request)
    } catch (error) {
      console.error('Failed to set color cycle effect:', error)
      toast.error(`Failed to set color cycle effect: ${error}`)
    }
  },
  Strobe: async (settings: EffectSettings, lights: Light[], turnOn: boolean) => {
    if (lights.length === 0) {
      return
    }

    const speed = settings.configs.speed.value || 0
    const request: StrobeRequest = {
      deviceIDs: lights.map((light) => light.target),
      speed,
      turnOn
    }
    try {
      await window.electron.ipcRenderer.invoke('strobe', request)
    } catch (error) {
      console.error('Failed to set strobe effect:', error)
      toast.error(`Failed to set strobe effect: ${error}`)
    }
  },
  Visualizer: async (settings: EffectSettings, lights: Light[], turnOn: boolean) => {
    const variation = (settings.configs.variation.value || 0) / 100
    const request: VisualizerRequest = {
      deviceIDs: lights.map((light) => light.target),
      variation,
      turnOn
    }
    try {
      await window.electron.ipcRenderer.invoke('visualizer', request)
    } catch (error) {
      console.error('Failed to set visualizer effect:', error)
      toast.error(`Failed to set visualizer effect: ${error}`)
    }
  },
  Theater: async (settings: EffectSettings, lights: Light[], turnOn: boolean) => {
    const request: TheaterRequest = {
      deviceIDs: lights.map((light) => light.target),
      turnOn,
      screen: settings.configs.screen.value || 0
    }
    try {
      await window.electron.ipcRenderer.invoke('theater', request)
    } catch (error) {
      console.error('Failed to set theater effect:', error)
      toast.error(`Failed to set theater effect: ${error}`)
    }
  }
}

import { Label } from '@radix-ui/react-label'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle
} from '@renderer/components/ui/dialog'
import { Slider } from '@renderer/components/ui/slider'
import type { EffectConfig, EffectSettings, EffectType } from '@renderer/schema/effectSchema'
import { calculateTime } from '@renderer/utils/time'
import type { Screen } from '@main/screenCapture'
import { useState } from 'react'
import { X } from 'lucide-react'
import { toast } from 'sonner'

interface EffectSettingsModalProps {
  effect: EffectType
  updateCardSettings: (cardId: number, newSettings: EffectSettings) => void
}

function EffectSettingsModal(props: EffectSettingsModalProps): JSX.Element {
  const { effect, updateCardSettings } = props
  const [screens, setScreens] = useState<Screen[]>([])
  const [isScreenSelectionOpen, setIsScreenSelectionOpen] = useState(false)
  const [activeScreenSelector, setActiveScreenSelector] = useState<string | null>(null)

  const handleInputChange = (configName: string, value: number): void => {
    const newConfigs = {
      ...effect.settings.configs,
      [configName]: { ...effect.settings.configs[configName], value }
    }
    const newSettings = { ...effect.settings, configs: newConfigs }
    updateCardSettings(effect.id, newSettings)
  }

  const fetchScreens = async (): Promise<void> => {
    try {
      const fetchedScreens = await window.electron.ipcRenderer.invoke('get-screens')
      setScreens(fetchedScreens)
    } catch (error) {
      console.error('Failed to fetch screens:', error)
      toast.error(`Failed to fetch screens: ${error}`)
    }
  }

  const handleScreenSelectorClick = (configName: string): void => {
    fetchScreens()
    setActiveScreenSelector(configName)
    setIsScreenSelectionOpen(true)
  }

  const handleScreenSelection = (screenId: number): void => {
    if (activeScreenSelector) {
      handleInputChange(activeScreenSelector, screenId)
    }
    setIsScreenSelectionOpen(false)
    setActiveScreenSelector(null)
  }

  const renderControl = (config: EffectConfig): JSX.Element | null => {
    switch (config.type) {
      case 'slider': {
        return (
          <>
            <Slider
              defaultValue={[config.value || 0]}
              max={config.max}
              step={config.step}
              min={config.min}
              className="w-[60%] mr-4"
              onValueChange={(e) =>
                handleInputChange(config.name.toLowerCase(), e[0] || config.min || 0)
              }
              value={[config.value || 0]}
            />
            <Label>{`${calculateTime(config.value || 0, config.label_max || 60).toFixed(0)}${config.unit}`}</Label>
          </>
        )
      }
      case 'screenSelector': {
        const selectedScreen = screens.find((screen) => screen.value === config.value)
        return (
          <button
            type="button"
            onClick={() => handleScreenSelectorClick(config.name.toLowerCase())}
          >
            {selectedScreen ? selectedScreen.name : 'Select Screen'}
          </button>
        )
      }
      default: {
        return null
      }
    }
  }

  return (
    <>
      <DialogContent className="bg-[#0B0E1F] text-white">
        <DialogHeader>
          <DialogTitle>{effect.settings.effect}</DialogTitle>
          <DialogDescription>{effect.settings.description}</DialogDescription>
        </DialogHeader>
        <div className="flex flex-col gap-4 py-4">
          {Object.values(effect.settings.configs).map((config) => (
            <div key={config.name} className="flex items-center">
              <Label className="text-left w-[5vw]">{config.name}</Label>
              {renderControl(config)}
            </div>
          ))}
        </div>
      </DialogContent>
      <Dialog open={isScreenSelectionOpen} onOpenChange={setIsScreenSelectionOpen}>
        <DialogContent className="bg-[#0B0E1F] text-white">
          <DialogHeader>
            <DialogTitle>Select a Screen</DialogTitle>
          </DialogHeader>
          <button
            type="button"
            className="absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground"
            onClick={() => setIsScreenSelectionOpen(false)}
          >
            <X className="h-4 w-4" />
            <span className="sr-only">Close</span>
          </button>
          <div className="grid grid-cols-2 gap-4 py-4">
            {screens.map((screen) => (
              <div
                key={screen.id}
                className="cursor-pointer hover:bg-gray-700 p-2 rounded"
                onClick={() => handleScreenSelection(screen.value)}
                onKeyUp={() => handleScreenSelection(screen.value)}
              >
                <img
                  src={screen.thumbnail.toDataURL()}
                  alt={screen.name}
                  className="w-full h-auto mb-2 rounded"
                />
                <p className="text-center">{screen.name}</p>
                <p className="text-center text-xs text-gray-400">ID: {screen.value}</p>
              </div>
            ))}
          </div>
        </DialogContent>
      </Dialog>
    </>
  )
}

export default EffectSettingsModal

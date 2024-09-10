import { Route, Routes, HashRouter as Router } from 'react-router-dom'
import close from './assets/close.svg'
import minimize from './assets/minimize.svg'
import Home from './components/Home'
import GroupDisplay from './components/GroupDisplay'
import { useState, useEffect } from 'react'
import type { GetDevicesResponse } from '../../proto/proto/message_service'
import { IconRefresh } from '@tabler/icons-react'
import LightsHubDisplay from './components/light_menus/LightsHubDisplay'
import useLightsStore from './hooks/useLightsStore'
import { toast, Toaster } from 'sonner'
import { dotPulse } from 'ldrs'

function App(): JSX.Element {
  const [isLoading, setLoading] = useState<boolean>(false)
  const { setLocations, emptyLocations } = useLightsStore()
  dotPulse.register()

  const fetchDevices = async (): Promise<void> => {
    try {
      const devices = (await window.electron.ipcRenderer.invoke(
        'fetch-devices'
      )) as GetDevicesResponse
      setLocations(devices)
      setLoading(false)
    } catch (error) {
      console.error('Failed to fetch devices:', error)
      toast.error(`Failed to fetch devices: ${error}`)
      setLoading(false)
      emptyLocations()
    }
  }

  const handleMinimize = async (): Promise<void> => {
    try {
      await window.electron.ipcRenderer.invoke('minimize-window')
    } catch (error) {
      console.error('Failed to minimize the window:', error)
      toast.error(`Failed to minimize the window: ${error}`)
    }
  }

  const handleClose = async (): Promise<void> => {
    try {
      await window.electron.ipcRenderer.invoke('close-window')
    } catch (error) {
      console.error('Failed to close the window:', error)
      toast.error(`Failed to close the window: ${error}`)
    }
  }

  useEffect(() => {
    const fetchDevicesInterval = setInterval(async () => {
      await fetchDevices()
    }, 7000)

    setLoading(true)
    fetchDevices().catch(console.error)

    return () => clearInterval(fetchDevicesInterval)
  }, [])

  return (
    <Router>
      <div className="h-screen w-screen relative">
        <Toaster position="bottom-right" theme="dark" richColors closeButton={true} />
        <div className="flex flex-col items-center py-4 px-8">
          {/* Header */}
          <div className="w-full flex justify-between items-center mb-4 ">
            <div className="draggable-header flex-grow">
              <h1 className="text-white text-4xl mb-2 font-bold text-left">LIFX Master</h1>
            </div>
            <div className="flex gap-x-6">
              <IconRefresh
                className="cursor-pointer"
                stroke={1.8}
                color="white"
                size={45}
                onClick={() => {
                  setLoading(true)
                  fetchDevices()
                }}
              />
              <img
                className="text-white cursor-pointer"
                src={minimize}
                onClick={handleMinimize}
                onKeyUp={handleMinimize}
                aria-label={'minimize'}
              />
              <img
                className="text-white cursor-pointer"
                src={close}
                onClick={handleClose}
                onKeyUp={handleClose}
                aria-label={'close'}
              />
            </div>
          </div>
          {/* Content */}
          <div className="w-full">
            <Routes>
              <Route path="/" element={<Home isLoading={isLoading} />} />
              <Route
                path="/locations/:locationId"
                element={<GroupDisplay isLoading={isLoading} />}
              />
              <Route path="/locations/:locationId/groups/:groupId" element={<LightsHubDisplay />} />
            </Routes>
          </div>
          <div className="grid grid-cols-5 gap-20">{}</div>
        </div>
      </div>
    </Router>
  )
}

export default App

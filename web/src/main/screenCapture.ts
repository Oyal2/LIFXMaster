import { desktopCapturer, type SourcesOptions, ipcMain } from 'electron'

export interface Screen {
  id: string
  name: string
  thumbnail: Electron.NativeImage
  value: number
}

async function getScreens(): Promise<Screen[]> {
  const options: SourcesOptions = {
    types: ['screen'],
    thumbnailSize: { width: 800, height: 800 }
  }

  try {
    const sources = await desktopCapturer.getSources(options)
    return sources.map((source, index) => {
      return {
        id: source.id,
        name: source.name,
        thumbnail: source.thumbnail,
        display_id: source.display_id,
        value: index
      }
    })
  } catch (error) {
    console.error('Error fetching screens:', error)
    return []
  }
}

export function setupScreenCapture(): void {
  ipcMain.handle('get-screens', async () => {
    return getScreens()
  })
}

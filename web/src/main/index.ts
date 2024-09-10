import { app, shell, BrowserWindow, ipcMain } from 'electron'
import { join } from 'node:path'
import { electronApp, optimizer, is } from '@electron-toolkit/utils'
import icon from '../../resources/icon.png'
import { spawn } from 'node:child_process'
import {
  colorCycle,
  fetchDevices,
  setColor,
  setDeviceLabel,
  setGroupLabel,
  setLocationLabel,
  setPower,
  strobe,
  theater,
  twinkle,
  visualizer
} from './client'
import type {
  ColorCycleRequest,
  HSBK,
  SetDeviceLabelRequest,
  SetGroupLabelRequest,
  SetLocationLabelRequest,
  StrobeRequest,
  TheaterRequest,
  TwinkleRequest,
  VisualizerRequest
} from '../proto/proto/message_service'
import { setupScreenCapture } from './screenCapture'

function createWindow(): void {
  // Create the browser window.
  const mainWindow = new BrowserWindow({
    width: 1560,
    height: 820,
    minWidth: 1560,
    minHeight: 820,
    show: false,
    frame: false,
    autoHideMenuBar: true,
    ...(process.platform === 'linux' ? { icon } : {}),
    webPreferences: {
      preload: join(__dirname, '../preload/index.js'),
      sandbox: false
    }
  })
  mainWindow.setMenuBarVisibility(false)

  mainWindow.on('ready-to-show', () => {
    mainWindow.show()
  })

  mainWindow.webContents.setWindowOpenHandler((details) => {
    shell.openExternal(details.url)
    return { action: 'deny' }
  })

  // HMR for renderer base on electron-vite cli.
  // Load the remote URL for development or the local html file for production.
  if (is.dev && process.env.ELECTRON_RENDERER_URL) {
    mainWindow.loadURL(process.env.ELECTRON_RENDERER_URL)
  } else {
    mainWindow.loadFile(join(__dirname, '../renderer/index.html'))
  }

  // Setup IPC to handle renderer process requests
  ipcMain.handle('fetch-devices', async () => {
    return await fetchDevices()
  })

  ipcMain.handle('set-color', async (_, colors: { [key: string]: HSBK }) => {
    return await setColor({
      colors
    })
  })

  ipcMain.handle('set-power', async (_, powers: { [key: string]: boolean }) => {
    return await setPower({
      powers
    })
  })

  ipcMain.handle('strobe', async (_, request: StrobeRequest) => {
    return await strobe(request)
  })

  ipcMain.handle('color-cycle', async (_, request: ColorCycleRequest) => {
    return await colorCycle(request)
  })

  ipcMain.handle('twinkle', async (_, request: TwinkleRequest) => {
    return await twinkle(request)
  })

  ipcMain.handle('visualizer', async (_, request: VisualizerRequest) => {
    return await visualizer(request)
  })

  ipcMain.handle('theater', async (_, request: TheaterRequest) => {
    return await theater(request)
  })

  ipcMain.handle('set-location-label', async (_, request: SetLocationLabelRequest) => {
    return await setLocationLabel(request)
  })

  ipcMain.handle('set-group-label', async (_, request: SetGroupLabelRequest) => {
    return await setGroupLabel(request)
  })

  ipcMain.handle('set-device-label', async (_, request: SetDeviceLabelRequest) => {
    return await setDeviceLabel(request)
  })

  // IPC Event Handlers
  ipcMain.handle('minimize-window', () => {
    mainWindow.minimize()
  })

  ipcMain.handle('close-window', () => {
    mainWindow.close()
  })
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.whenReady().then(() => {
  // Set app user model id for windows
  electronApp.setAppUserModelId('com.electron')

  setupScreenCapture()
  startGRPCServer()

  // Default open or close DevTools by F12 in development
  // and ignore CommandOrControl + R in production.
  // see https://github.com/alex8088/electron-toolkit/tree/master/packages/utils
  app.on('browser-window-created', (_, window) => {
    optimizer.watchWindowShortcuts(window)
  })

  createWindow()

  app.on('activate', () => {
    // On macOS it's common to re-create a window in the app when the
    // dock icon is clicked and there are no other windows open.
    if (BrowserWindow.getAllWindows().length === 0) createWindow()
  })
})

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit()
  }
})

function startGRPCServer(): void {
  const grpcServerPath = join(process.resourcesPath, 'grpc-server.exe')
  console.log('Attempting to start gRPC server from:', grpcServerPath)

  try {
    const grpcServer = spawn(grpcServerPath)

    grpcServer.stdout.on('data', (data) => {
      console.log(`gRPC server output: ${data}`)
    })

    grpcServer.stderr.on('data', (data) => {
      console.error(`gRPC server error: ${data}`)
    })

    grpcServer.on('error', (error) => {
      console.error('Failed to start gRPC server:', error)
    })

    grpcServer.on('close', (code) => {
      console.log(`gRPC server exited with code ${code}`)
    })
  } catch (error) {
    console.error('Exception when starting gRPC server:', error)
  }
}

import { DeviceServiceClient } from '../proto/proto/message_service.client'
import type {
  ColorCycleRequest,
  ColorCycleResponse,
  GetDevicesRequest,
  GetDevicesResponse,
  SetColorRequest,
  SetColorResponse,
  SetDeviceLabelRequest,
  SetDeviceLabelResponse,
  SetGroupLabelRequest,
  SetGroupLabelResponse,
  SetLocationLabelRequest,
  SetLocationLabelResponse,
  SetPowerRequest,
  SetPowerResponse,
  StrobeRequest,
  StrobeResponse,
  TheaterRequest,
  TheaterResponse,
  TwinkleRequest,
  TwinkleResponse,
  VisualizerRequest,
  VisualizerResponse
} from '../proto/proto/message_service'
import { GrpcTransport } from '@protobuf-ts/grpc-transport'
import { ChannelCredentials } from '@grpc/grpc-js'
const transport = new GrpcTransport({
  host: 'localhost:50051',
  channelCredentials: ChannelCredentials.createInsecure()
})

export async function fetchDevices(): Promise<GetDevicesResponse> {
  const request: GetDevicesRequest = {}
  try {
    const client = new DeviceServiceClient(transport)
    const call = client.getDevices(request)
    const response: GetDevicesResponse = await call.response
    return response
  } catch (error) {
    console.error('Failed to fetch devices:', error)
    throw error
  }
}

export async function setColor(request: SetColorRequest): Promise<SetColorResponse> {
  try {
    const client = new DeviceServiceClient(transport)
    const call = client.setColor(request)
    const response: SetColorResponse = await call.response
    return response
  } catch (error) {
    console.error('Failed to set color:', error)
    throw error
  }
}

export async function setPower(request: SetPowerRequest): Promise<SetPowerResponse> {
  try {
    const client = new DeviceServiceClient(transport)
    const call = client.setPower(request)
    const response: SetPowerResponse = await call.response
    return response
  } catch (error) {
    console.error('Failed to set power:', error)
    throw error
  }
}

export async function setLocationLabel(
  request: SetLocationLabelRequest
): Promise<SetLocationLabelResponse> {
  try {
    const client = new DeviceServiceClient(transport)
    const call = client.setLocationLabel(request)
    const response: SetLocationLabelResponse = await call.response
    return response
  } catch (error) {
    console.error('Failed to set location label:', error)
    throw error
  }
}

export async function setGroupLabel(request: SetGroupLabelRequest): Promise<SetGroupLabelResponse> {
  try {
    const client = new DeviceServiceClient(transport)
    const call = client.setGroupLabel(request)
    const response: SetGroupLabelResponse = await call.response
    return response
  } catch (error) {
    console.error('Failed to set group label:', error)
    throw error
  }
}

export async function setDeviceLabel(
  request: SetDeviceLabelRequest
): Promise<SetGroupLabelResponse> {
  try {
    const client = new DeviceServiceClient(transport)
    const call = client.setDeviceLabel(request)
    const response: SetDeviceLabelResponse = await call.response
    return response
  } catch (error) {
    console.error('Failed to set device label:', error)
    throw error
  }
}

export async function strobe(request: StrobeRequest): Promise<StrobeResponse> {
  try {
    const client = new DeviceServiceClient(transport)
    const call = client.strobe(request)
    const response: StrobeResponse = await call.response
    return response
  } catch (error) {
    console.error('Failed to start strobe:', error)
    throw error
  }
}

export async function colorCycle(request: ColorCycleRequest): Promise<StrobeResponse> {
  try {
    const client = new DeviceServiceClient(transport)
    const call = client.colorCycle(request)
    const response: ColorCycleResponse = await call.response
    return response
  } catch (error) {
    console.error('Failed to start color cycle:', error)
    throw error
  }
}

export async function twinkle(request: TwinkleRequest): Promise<TwinkleResponse> {
  try {
    const client = new DeviceServiceClient(transport)
    const call = client.twinkle(request)
    const response: TwinkleResponse = await call.response
    return response
  } catch (error) {
    console.error('Failed to start twinkle:', error)
    throw error
  }
}

export async function visualizer(request: VisualizerRequest): Promise<VisualizerResponse> {
  try {
    const client = new DeviceServiceClient(transport)
    const call = client.visualizer(request)
    const response: VisualizerResponse = await call.response
    return response
  } catch (error) {
    console.error('Failed to start visualizer:', error)
    throw error
  }
}

export async function theater(request: TheaterRequest): Promise<TheaterResponse> {
  try {
    const client = new DeviceServiceClient(transport)
    const call = client.theater(request)
    const response: TwinkleResponse = await call.response
    return response
  } catch (error) {
    console.error('Failed to start theater:', error)
    throw error
  }
}

import { HsvaEqual, Uint16HsvToHex } from '@renderer/utils/hsv'
import { hexToHsva, type HsvaColor } from '@uiw/react-color'
import type { Device, GetDevicesResponse } from 'src/proto/proto/message_service'
import { create } from 'zustand'

export interface Light extends Device {
  color: HsvaColor
  effect: string | undefined
}

export interface GroupStore {
  id: string
  locationId: string
  label: string
  isOn: boolean
  color: HsvaColor
  lights: Light[]
}

export interface LocationStore {
  id: string
  label: string
  isOn: boolean
  color: HsvaColor
  groups: { [groupId: string]: GroupStore }
}

interface LightsState {
  locations: { [locationId: string]: LocationStore }
  setLocations: (devices: GetDevicesResponse) => void
  updateGroupLights: (locationId: string, groupId: string, newLights: Light[]) => void
  updateLocations: (updatedLocationStore: { [locationId: string]: LocationStore }) => void
  updateLocation: (locationId: string, updates: Partial<LocationStore>) => void
  updateGroup: (locationId: string, groupId: string, updates: Partial<GroupStore>) => void
  updateLight: (lightId: string, updates: Partial<Light>) => void
  emptyLocations: () => void
}

const defaultColor: HsvaColor = {
  a: 1,
  h: 0,
  s: 0,
  v: 100
}

const useLightsStore = create<LightsState>((set) => ({
  locations: {},
  setLocations: (devices: GetDevicesResponse): void => {
    const locations = devices.locations
    const newLocationStore: { [locationId: string]: LocationStore } = {}

    for (const [locationKey, location] of Object.entries(locations)) {
      let locationColor: HsvaColor | null = null
      newLocationStore[locationKey] = {
        isOn: false,
        id: locationKey,
        label: location.label,
        color: defaultColor,
        groups: {}
      }

      for (const [groupKey, group] of Object.entries(location.groups)) {
        const lights: Light[] = [
          {
            target: BigInt(-1),
            color: {
              a: 1,
              h: 0,
              s: 0,
              v: 100
            },
            label: {
              label: 'All Lights'
            },
            address: '',
            port: 0,
            effect: undefined,
            power: {
              level: 0
            }
          }
        ]
        let power = 0
        let groupColor: HsvaColor | null = null
        for (const device of group.devices) {
          const hexColor = Uint16HsvToHex(
            device.light?.hue,
            device.light?.saturation,
            device.light?.brightness
          )
          const hsva = hexToHsva(hexColor)
          if (groupColor === null) {
            groupColor = hsva
            // If locationColor is null, this is also the first group in the location
            locationColor = locationColor === null ? hsva : locationColor
          } else if (!HsvaEqual(groupColor, hsva)) {
            groupColor = defaultColor // Set group color to white
            // If location color isn't already white, set it to white
            if (locationColor && !HsvaEqual(locationColor, defaultColor)) {
              locationColor = defaultColor
            }
          } else if (locationColor === null || !HsvaEqual(locationColor, groupColor)) {
            locationColor = groupColor
          }

          if (device.power !== undefined && device.power?.level > 0) {
            power = 100
          }

          lights.push({
            color: hsva,
            ...device,
            effect: undefined
          })
        }
        if (power > 0) {
          lights[0].power = {
            level: power
          }
          newLocationStore[locationKey].isOn = true
        }
        const finalColor = groupColor ?? defaultColor

        lights[0].color = finalColor

        newLocationStore[locationKey].groups[groupKey] = {
          id: groupKey,
          label: group.label,
          lights: lights,
          isOn: power > 0,
          color: finalColor,
          locationId: locationKey
        }
      }
      newLocationStore[locationKey].color = locationColor ?? defaultColor
    }

    set({ locations: newLocationStore })
  },
  updateGroupLights: (locationId: string, groupId: string, newLights: Light[]): void =>
    set((state) => {
      const location = state.locations[locationId]
      if (!location) return state

      const group = location.groups[groupId]
      if (!group) return state

      const hsva = group.lights.every((light) => HsvaEqual(light.color, group.lights[0].color))
        ? group.lights[0].color
        : defaultColor

      group.color = hsva
      if (!HsvaEqual(location.color, group.color)) {
        location.color = group.color
      }

      return {
        ...state,
        locations: {
          ...state.locations,
          [locationId]: {
            ...location,
            groups: {
              ...location.groups,
              [groupId]: {
                ...group,
                lights: newLights
              }
            }
          }
        }
      }
    }),
  updateLocations: (updatedLocationStore: { [locationId: string]: LocationStore }): void => {
    set({ locations: updatedLocationStore })
  },
  updateLocation: (locationId: string, updates: Partial<LocationStore>): void =>
    set((state) => ({
      locations: {
        ...state.locations,
        [locationId]: {
          ...state.locations[locationId],
          ...updates
        }
      }
    })),
  updateGroup: (locationId: string, groupId: string, updates: Partial<GroupStore>): void =>
    set((state) => ({
      locations: {
        ...state.locations,
        [locationId]: {
          ...state.locations[locationId],
          groups: {
            ...state.locations[locationId].groups,
            [groupId]: {
              ...state.locations[locationId].groups[groupId],
              ...updates
            }
          }
        }
      }
    })),
  updateLight: (lightId: string, updates: Partial<Light>): void =>
    set((state) => {
      const newLocations = { ...state.locations }

      for (const locationId in newLocations) {
        const location = newLocations[locationId]
        for (const groupId in location.groups) {
          const group = location.groups[groupId]
          const updatedLights = group.lights.map((light) =>
            light.target.toString() === lightId ? { ...light, ...updates } : light
          )

          if (updatedLights.some((light) => light.target.toString() === lightId)) {
            newLocations[locationId].groups[groupId].lights = updatedLights
            return { locations: newLocations }
          }
        }
      }

      return state
    }),
  emptyLocations: (): void => {
    set({})
  }
}))

export default useLightsStore

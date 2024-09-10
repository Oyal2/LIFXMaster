import type { GroupStore, Light, LocationStore } from '@renderer/hooks/useLightsStore'
import { useMemo } from 'react'

export const getGroupLights = (locations: {
  [key: string]: LocationStore
}): ((locId: string, grpId: string) => Light[]) => {
  return useMemo(() => {
    return (locId: string, grpId: string): Light[] => {
      return locations[locId]?.groups[grpId]?.lights || []
    }
  }, [locations])
}

export const getLights = (
  locations: { [key: string]: LocationStore },
  locationId: string | undefined,
  groupId: string | undefined
): Light[] => {
  const groupLightsFn = getGroupLights(locations)

  return useMemo(() => {
    if (locationId && groupId) {
      return groupLightsFn(locationId, groupId)
    }
    return []
  }, [locationId, groupId, groupLightsFn])
}

export const checkAnyLightOn = (locations: { [key: string]: LocationStore }): boolean => {
  for (const location of Object.values(locations)) {
    if (location.isOn) {
      return true
    }
  }
  return false
}

export const checkAnyGroupLightOn = (groups: { [groupId: string]: GroupStore }): boolean => {
  for (const group of Object.values(groups)) {
    if (group.isOn) {
      return true
    }
  }
  return false
}

export const checkIfAllLightsExists = (lights: Light[]): boolean => {
  if (lights.length === 0) {
    return false
  }

  return lights[0].target.toString() === '-1'
}

import type React from 'react'

interface WifiIconProps {
  signalStrength: number
}

const WifiIcon: React.FC<WifiIconProps> = ({ signalStrength }) => {
  const getSignalLevel = (signal: number): number => {
    const rssi = Math.floor(10 * Math.log10(signal) + 0.5)
    let status: number

    if (rssi < 0 || rssi === 200) {
      if (rssi === 200) {
        status = 0
      } else if (rssi <= -80) {
        status = 1
      } else if (rssi <= -70) {
        status = 2
      } else if (rssi <= -60) {
        status = 3
      } else {
        status = 4
      }
    } else if (rssi === 4 || rssi === 5 || rssi === 6) {
      status = 1
    } else if (rssi >= 7 && rssi <= 11) {
      status = 2
    } else if (rssi >= 12 && rssi <= 16) {
      status = 3
    } else if (rssi > 16) {
      status = 4
    } else {
      status = 0
    }

    return status
  }

  const level = getSignalLevel(signalStrength)

  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width={48}
      height={48}
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth={2}
      strokeLinecap="round"
      strokeLinejoin="round"
      className="icon icon-tabler icons-tabler-outline icon-tabler-wifi"
    >
      <path stroke="none" d="M0 0h24v24H0z" fill="none" />
      <path d="M12 18l.01 0" stroke={level >= 1 ? '#4caf50' : '#e0e0e0'} />
      <path d="M9.172 15.172a4 4 0 0 1 5.656 0" stroke={level >= 2 ? '#4caf50' : '#e0e0e0'} />
      <path d="M6.343 12.343a8 8 0 0 1 11.314 0" stroke={level >= 3 ? '#4caf50' : '#e0e0e0'} />
      <path
        d="M3.515 9.515c4.686 -4.687 12.284 -4.687 17 0"
        stroke={level >= 4 ? '#4caf50' : '#e0e0e0'}
      />
    </svg>
  )
}

export default WifiIcon

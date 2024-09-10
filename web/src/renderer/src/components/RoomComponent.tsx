import { TbChevronRight } from 'react-icons/tb'
import { MdOutlineEdit } from 'react-icons/md'

import GlowingLightBulb from './LightBulb'
import { CircularProgressBar } from 'react-percentage-bar'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button } from './ui/button'
import { FaCheck } from 'react-icons/fa'
import { TiCancel } from 'react-icons/ti'
import { Input } from './ui/input'
import useEditableLabel from '@renderer/hooks/useEditableLabel'

export interface LightBulb {
  name: string
  color: string
  percentage: number
  isOn: boolean
}

interface RoomProps {
  lightbulb: LightBulb
  uuid: string
  roomType: RoomType
  onClick?: () => void
  onUpdateLabel?: (location: string, newLabel: string) => void
}

type RoomType = 'locations' | 'groups' | 'all_lights'

function RoomComponent(props: RoomProps): JSX.Element {
  const { lightbulb, uuid, roomType, onClick, onUpdateLabel } = props
  const [isHovered, setIsHovered] = useState<boolean>(false)
  const [isAllLightsOn, setIsAllLightsOn] = useState<boolean>(lightbulb.isOn)
  const navigate = useNavigate()

  const { view, value, handleKeyUp, handleChange, handleCancel, handleSave, switchToInput } =
    useEditableLabel({
      initialName: lightbulb.name,
      onUpdateLabel,
      uuid
    })

  const handleLightBulbClick = (): void => {
    if (roomType === 'all_lights') {
      setIsAllLightsOn(!isAllLightsOn)
      if (onClick) onClick()
    } else {
      navigate(`${roomType}/${uuid}`)
    }
  }

  const getBulbColor = (): string => {
    if (roomType === 'all_lights') {
      return isAllLightsOn ? lightbulb.color : 'transparent'
    }
    return lightbulb.color
  }

  const renderView = (): JSX.Element => {
    return view === 'label' ? (
      <span
        className="text-white text-xl py-2"
        onClick={() => {
          if (roomType !== 'all_lights') {
            switchToInput
          }
        }}
        onKeyUp={() => {
          if (roomType !== 'all_lights') {
            switchToInput
          }
        }}
      >
        {lightbulb.name}
      </span>
    ) : (
      <div className="flex items-center justify-center space-x-2">
        <Input
          className="text-white text-xl w-2/3 py-2 h-10 bg-gray-800 border border-gray-700 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          value={value}
          onClick={(e) => {
            e.stopPropagation()
            e.preventDefault()
          }}
          onKeyUp={handleKeyUp}
          onChange={handleChange}
        />
        <Button
          className="p-2 h-10 bg-red-600 hover:bg-red-700 transition-colors duration-200 rounded-md focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-opacity-50"
          onClick={handleCancel}
        >
          <TiCancel size={24} className="text-white" />
        </Button>
        <Button
          className="p-2 h-10 bg-green-600 hover:bg-green-700 transition-colors duration-200 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-opacity-50"
          onClick={handleSave}
        >
          <FaCheck size={24} className="text-white" />
        </Button>
      </div>
    )
  }

  return (
    <>
      <div
        className="cursor-pointer flex flex-col items-center"
        onMouseEnter={() => setIsHovered(true && roomType !== 'all_lights')}
        onMouseLeave={() => setIsHovered(false && roomType !== 'all_lights')}
        onClick={handleLightBulbClick}
        onKeyUp={handleLightBulbClick}
      >
        <div
          className="flex justify-end w-full"
          onClick={(e) => {
            e.stopPropagation()
            e.preventDefault()
            switchToInput()
          }}
          onKeyUp={(e) => {
            e.stopPropagation()
            e.preventDefault()
            switchToInput()
          }}
        >
          <MdOutlineEdit
            size={27}
            className={`text-white icon-transition ${isHovered && view === 'label' ? 'show' : 'opacity-0'}`}
          />
        </div>
        <div className="relative w-full">
          <div className="flex justify-center relative">
            <CircularProgressBar
              showPercentage={false}
              startPosition={'0'}
              percentage={lightbulb.percentage}
              radius={'6rem'}
              size={12}
              color={getBulbColor()}
              trackColor="#171717"
            >
              <div className="relative">
                <GlowingLightBulb
                  glowColor={lightbulb.color}
                  imageType={roomType === 'all_lights' ? 'Power' : 'Lightbulb'}
                  size={roomType === 'all_lights' ? 100 : 120}
                />
              </div>
            </CircularProgressBar>
          </div>
        </div>
        <div className="mt-2 flex justify-between w-full">
          <div className="invisible">
            <TbChevronRight size={27} />
          </div>
          <div className="text-white text-center text-lg">{renderView()}</div>
          {view === 'label' && (
            <TbChevronRight
              size={27}
              className={`text-white icon-transition ${isHovered ? 'show' : 'opacity-0'}`}
            />
          )}
        </div>
      </div>
    </>
  )
}

export default RoomComponent

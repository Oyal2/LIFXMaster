import { useCallback, useState } from 'react'

type ViewType = 'label' | 'input'

interface UseEditableLabelProps {
  initialName: string
  onUpdateLabel?: (uuid: string, newName: string) => void
  uuid: string
}

interface useEditableLabel {
  view: ViewType
  value: string
  handleKeyUp: (e: React.KeyboardEvent<HTMLInputElement>) => void
  handleChange: (e: React.ChangeEvent<HTMLInputElement>) => void
  handleCancel: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void
  handleSave: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void
  switchToInput: () => void
}

const useEditableLabel = ({
  initialName,
  onUpdateLabel,
  uuid
}: UseEditableLabelProps): useEditableLabel => {
  const [view, setView] = useState<ViewType>('label')
  const [value, setValue] = useState(initialName)

  const handleKeyUp = useCallback(
    (e: React.KeyboardEvent<HTMLInputElement>) => {
      e.stopPropagation()
      e.preventDefault()
      if (e.key === 'Enter') {
        setValue(e.currentTarget.value)
        if (onUpdateLabel) {
          onUpdateLabel(uuid, e.currentTarget.value)
        }
        setView('label')
      } else if (e.key === 'Escape') {
        setValue(initialName)
        setView('label')
      }
    },
    [initialName, onUpdateLabel, uuid]
  )

  const handleChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    setValue(e.target.value)
  }, [])

  const handleCancel = useCallback(
    (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
      e.preventDefault()
      e.stopPropagation()
      setValue(initialName)
      setView('label')
    },
    [initialName]
  )

  const handleSave = useCallback(
    (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
      e.preventDefault()
      e.stopPropagation()
      if (onUpdateLabel) {
        onUpdateLabel(uuid, value)
      }
      setView('label')
    },
    [onUpdateLabel, uuid, value]
  )

  const switchToInput = useCallback(() => {
    setView('input')
  }, [])

  return {
    view,
    value,
    handleKeyUp,
    handleChange,
    handleCancel,
    handleSave,
    switchToInput
  }
}

export default useEditableLabel

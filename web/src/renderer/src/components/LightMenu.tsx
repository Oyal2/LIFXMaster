
interface LightMenuProps {
  items: string[]
  setMenuSelection: React.Dispatch<React.SetStateAction<number>>
  menuSelection: number
}

function LightMenu(prop: LightMenuProps): JSX.Element {
  const { items, setMenuSelection, menuSelection } = prop

  return (
    <div className="flex">
      <svg height="65" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
        <polygon points="100,100 100,0 0,0" fill="#1A1C48" />
      </svg>
      {items.map((item, index) => (
        <div
          key={`${index}_${item}`}
          className="text-white bg-[#1A1C48] py-2 px-4 flex-1 h-[65px] font-semibold text-lg flex items-center justify-center cursor-pointer"
          onClick={() => setMenuSelection(index)}
          onKeyUp={() => setMenuSelection(index)}
        >
          <h1 className={`${menuSelection === index ? 'border-b-4' : ''}`}>{item}</h1>
        </div>
      ))}
      <svg height="65" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
        <polygon points="0,100 100,0 0,0" fill="#1A1C48" />
      </svg>
    </div>
  )
}

export default LightMenu

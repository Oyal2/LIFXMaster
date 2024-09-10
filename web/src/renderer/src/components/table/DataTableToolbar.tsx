import type { Table } from '@tanstack/react-table'
import { Input } from '../../components/ui/input'

interface DataTableToolbarProps<TData> {
  table: Table<TData>
}

export function DataTableToolbar<TData>({ table }: DataTableToolbarProps<TData>): JSX.Element {
  return (
    <div className="flex items-center justify-between ml-2">
      <div className="flex flex-1 items-center space-x-2">
        <Input
          placeholder="Filter Lights...."
          value={(table.getColumn('label')?.getFilterValue() as string) ?? ''}
          onChange={(event) => {
            table.getColumn('label')?.setFilterValue(event.target.value)
          }}
          className="h-8 w-[150px] lg:w-[250px] border-[#535353] text-[#9499C3] font-semibold"
        />
      </div>
    </div>
  )
}

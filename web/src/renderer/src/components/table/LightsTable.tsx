'use client'

import type * as React from 'react'
import {
  type ColumnFiltersState,
  type RowSelectionState,
  type SortingState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable
} from '@tanstack/react-table'

import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '../ui/table'
import { DataTableToolbar } from './DataTableToolbar'
import { useState } from 'react'
import useLightsStore, { type Light } from '@renderer/hooks/useLightsStore'
import GetLightColumns from './column'
import { toast } from 'sonner'

interface DataTableProps {
  data: Light[]
  locationId: string | undefined
  groupId: string | undefined
  setRowSelection: React.Dispatch<React.SetStateAction<RowSelectionState>>
  rowSelection: RowSelectionState | undefined
}

export function DataTable(prop: DataTableProps): JSX.Element {
  const { data, locationId, groupId, rowSelection, setRowSelection } = prop
  const [sorting, setSorting] = useState<SortingState>([])
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([])
  const updateGroupLights = useLightsStore((state) => state.updateGroupLights)

  const handleAllLightsClick = async (lights: Light[]): Promise<void> => {
    if (locationId === undefined || groupId === undefined) return
    const powerRequest: { [key: string]: boolean } = {}
    const oldLevel = lights[0].power?.level || 0
    const newLevel = oldLevel > 0 ? 0 : 100

    const updatedLights = data.map((light: Light) => {
      const updatedLight = {
        ...light,
        power: { level: newLevel },
        effect: undefined
      }
      if (updatedLight.target.toString() !== '-1') {
        powerRequest[light.target.toString()] = oldLevel === 0
      }
      return updatedLight
    })
    try {
      await window.electron.ipcRenderer.invoke('set-power', powerRequest)
    } catch (error) {
      console.error("Failed to set devices' power:", error)
      toast.error(`Failed to set devices' power: ${error}`)
    }
    updateGroupLights(locationId, groupId, updatedLights)
  }

  const handleIndividualLightClick = async (light: Light): Promise<void> => {
    if (locationId === undefined || groupId === undefined) return
    const oldLevel = light.power?.level || 0
    const newLevel = oldLevel > 0 ? 0 : 100
    const powerRequest: { [key: string]: boolean } = {
      [light.target.toString()]: oldLevel === 0
    }
    const updatedLights = [...data]
    const index = updatedLights.findIndex((x) => x.address === light.address)
    if (index !== -1) {
      updatedLights[index].power = {
        level: newLevel
      }
      updatedLights[index].effect = undefined
    }
    try {
      await window.electron.ipcRenderer.invoke('set-power', powerRequest)
    } catch (error) {
      console.error("Failed to set device's power:", error)
      toast.error(`Failed to set device's power: ${error}`)
    }
    updateGroupLights(locationId, groupId, updatedLights)
  }

  const columns = GetLightColumns({
    lights: data,
    handleIndividualLightClick,
    handleAllLightsClick
  })

  const table = useReactTable({
    data,
    columns,
    onSortingChange: setSorting,
    onColumnFiltersChange: setColumnFilters,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onRowSelectionChange: setRowSelection,
    enableMultiRowSelection: false,
    state: {
      sorting,
      rowSelection,
      columnFilters,
      columnVisibility: {
        select: false
      }
    }
  })

  return (
    <div className="space-y-4">
      <DataTableToolbar table={table} />
      <div className="rounded-md text-[#94993] p-2 text-white h-[65vh] overflow-auto scrollbar scrollbar-thumb-rounded-full scrollbar-thumb-slate-700 scrollbar-thin">
        <Table className="">
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow key={headerGroup.id}>
                {headerGroup.headers.map((header) => {
                  return (
                    <TableHead
                      key={header.id}
                      colSpan={header.colSpan}
                      className="border border-[#535353] sticky top-0 z-10"
                    >
                      {header.isPlaceholder
                        ? null
                        : flexRender(header.column.columnDef.header, header.getContext())}
                    </TableHead>
                  )
                })}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected()}
                  className={`cursor-pointer hover:bg-blue-500 ${row.getIsSelected() ? 'bg-blue-500' : ''}`}
                  onClick={() => {
                    row.toggleSelected(!row.getIsSelected())
                  }}
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id} className="border border-[#535353]">
                      {flexRender(cell.column.columnDef.cell, cell.getContext())}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell colSpan={columns.length} className="h-24 text-center">
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
    </div>
  )
}

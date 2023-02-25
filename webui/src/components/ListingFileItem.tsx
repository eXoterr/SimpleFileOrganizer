import React, { FC, MouseEventHandler, ReactNode, useEffect, useState } from 'react'

type ListingFileItemProps = {
    filename: string
    icon?: ReactNode
    folder: boolean
    reloadListing: Function
    addToList: Function
    remFromList: Function
}

export default function ListingFileItem({filename, icon, folder, reloadListing, addToList, remFromList}: ListingFileItemProps) {
  
  const [selected, setSelected] = useState<boolean>(false)
  
  useEffect(() => {
    selected ? addToList(filename) : remFromList(filename)
  }, [selected])

  return (
    <div className='flex text-xl items-center gap-2 p-1 bg-slate-500 rounded-md w-full h-fit'>
        <input checked={selected} onChange={() => setSelected(!selected)} className='' type="checkbox" name="" id="" />
        {icon}
        {folder ? 
          <a href="#" onClick={() => reloadListing(filename)}>{filename}</a>
          :
          <label onClick={() => {setSelected(!selected)}}>{filename}</label>
        }
    </div>
  )
}
import React, { MouseEventHandler, useState } from 'react'
import { IDirectory } from '../types'
import ListingFileItem from './ListingFileItem'
import {AiFillFolderOpen, AiFillFile} from 'react-icons/ai'

type Props = {
    directory?: IDirectory
    name: string
    reloadListing: Function
    upFunction: () => void
    addToList: (filename: string) => void
    remFromList: (filename: string) => void
}

export default function Listing({directory, name, reloadListing, upFunction, addToList, remFromList}: Props) {
  

  return (
    <div className='w-full h-full'>
        <h2 className='text-center text-xl'>{name}</h2>
        <div className=' bg-slate-600 h-full overflow-scroll shadow shadow-slate-800 rounded-lg border border-slate-800 p-2 flex flex-col gap-2'>
            <ListingFileItem
                reloadListing={upFunction}
                filename='..'
                icon={<AiFillFolderOpen />}
                folder={true}
                addToList={() => {}}
                remFromList={() => {}}
            />
            {directory?.files?.map(
                f => <ListingFileItem
                        key={f.filename}
                        addToList={addToList}
                        remFromList={remFromList}
                        reloadListing={reloadListing}
                        filename={f.filename}
                        folder={f.directory}
                        icon={f.directory ? <AiFillFolderOpen /> : <AiFillFile />}
                    />
            )}
        </div>
    </div>
  )
}
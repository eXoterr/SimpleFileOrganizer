import React, { useEffect, useState } from 'react'
import Listing from './components/Listing'
import { SlMagicWand } from 'react-icons/sl'
import { IDirectory } from './types'
import { getFiles, getLibrary, organizerFilm } from './api'

type Props = {}

export default function App({}: Props) {

const [currentFilesPath, setCurrentFilesPath] = useState<string[]>([''])
const [currentLibraryPath, setCurrentLibraryPath] = useState<string[]>([''])
const [files, setFiles] = useState<IDirectory>()
const [library, setLibrary] = useState<IDirectory>()

useEffect(() => {
  getFiles(currentFilesPath).then(r => setFiles(r))
  getLibrary(currentLibraryPath).then(r => setLibrary(r))
}, [currentFilesPath, currentLibraryPath])

const goToDirectoryInFiles = (path: string) => {
  setCurrentFilesPath([...currentFilesPath, path])
}

const upDirectoryInFiles = () => {
  const arr = currentFilesPath
  arr.pop()
  setCurrentFilesPath([...arr])
}

const goToDirectoryInLibrary = (path: string) => { 
  setCurrentLibraryPath([...currentLibraryPath, path])
}

const orgFilm = (filename: string) => {
  organizerFilm([...currentFilesPath, filename])
}

const upDirectoryInLibrary = () => {
  const arr = currentLibraryPath
  arr.pop()
  setCurrentLibraryPath([...arr])
}

const [selectedFiles, setSelectedFiles] = useState<string[]>([]) 

  const addToList = (filename: string) => {
    setSelectedFiles([...selectedFiles, filename])
  }

  const remFromList = (filename: string) => {
    setSelectedFiles(selectedFiles.filter((file) => {
        return file != filename
    }))
  }


  return (
    <div className='bg-gray-800 text-white'>
      <div className='container mx-auto h-screen overflow-hidden'>
        <h1 className='text-3xl font-bold text-center mt-4'>Simple File Organizer</h1>

        <div className="flex gap-10 h-5/6 mt-4 items-center">
            <Listing
              name='Файлы'
              reloadListing={goToDirectoryInFiles}
              upFunction={upDirectoryInFiles}
              addToList={addToList}
              remFromList={remFromList}
              directory={files}
            />
            <div>
              <button className='text-xl p-3 bg-slate-600 rounded-lg' onClick={() => {
                selectedFiles.map(
                  f => orgFilm(f)
                )
                getLibrary(currentLibraryPath).then(r => setLibrary(r))
              }}>
                <SlMagicWand />
              </button>
            </div>
            <Listing
              name='Медиатека'
              reloadListing={goToDirectoryInLibrary}
              upFunction={upDirectoryInLibrary}
              directory={library}
              addToList={addToList}
              remFromList={remFromList}
            />
        </div>
      </div>
    </div>
  )
}
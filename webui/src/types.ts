export interface IFile{
    filename: string
    size: number
    directory: boolean
}

export interface IDirectory{
    files: IFile[]
}
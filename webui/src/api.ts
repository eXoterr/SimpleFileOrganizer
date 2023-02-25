import ky from "ky";
import { IDirectory } from "./types";


export async function getFiles(pathArr: string[]): Promise<IDirectory>{
    let path = createPath(pathArr)
    console.log(path)
    let files = await ky.get(`http://127.0.0.1:8080/files/list?path=${path}`).text()
    let result: IDirectory = JSON.parse(files)
    return result
}

export async function getLibrary(pathArr: string[]): Promise<IDirectory>{
    let path = createPath(pathArr)
    console.log(path)
    let files = await ky.get(`http://127.0.0.1:8080/files/library?path=${path}`).text()
    let result: IDirectory = JSON.parse(files)
    return result
}

export async function organizerFilm(pathArr: string[]): Promise<string>{
    let path = createPath(pathArr)
    console.log(path)
    const formdata = new FormData()
    formdata.append("path", path)
    let files = await ky.post(`http://127.0.0.1:8080/organize/film`, {body: formdata}).text()
    return files
}

function createPath(path: string[]): string{
    return path.join('/')
}
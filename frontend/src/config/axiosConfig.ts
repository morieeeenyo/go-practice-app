import axios from "axios"
import { Arguments } from "swr"

type Args ={
  arg: Arguments
}

const axiosDefault = axios.create({
  baseURL: "http://localhost:8080",
})

export const get = (path: string) =>
  axiosDefault.get(path).then((res) => res.data)

export const post = async (url: string, { arg }: Args) => {
  const response = await axiosDefault.post(url, arg)

  return response.data
}

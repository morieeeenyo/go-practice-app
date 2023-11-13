import axios from "axios"

const axiosDefault = axios.create({
  baseURL: "http://localhost:8080",
})

export const get = (path: string) => axiosDefault.get(path).then((res) => res.data)

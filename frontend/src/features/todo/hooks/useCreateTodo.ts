"use client"
import { post } from "@/config/axiosConfig"
import { useCallback } from "react"
import useSWRMutation from "swr/mutation"
import { FormValues } from "../components/TodoForm"
import { AxiosError, AxiosResponse } from "axios"
import { Todo } from "../types"

export const useCreateTodo = () => {
  const { trigger: createTodo, isMutating } = useSWRMutation<
    AxiosResponse<Todo>,
    AxiosError
  >("/todos", post)

  const onSubmit = async (data: FormValues) => {
    createTodo(data)
  }

  return {
    onSubmit,
    isMutating
  }
}

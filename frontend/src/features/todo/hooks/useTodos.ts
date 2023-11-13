"use client"
import useSWR from "swr"
import { Todo } from "../types"
import { get } from "@/config/axiosConfig"

export const useTodos = () => {
  const { data: todos, isLoading, error } = useSWR<Todo[]>("/todos", get)

  return {
    todos: todos ?? [],
    isLoading,
    error,
  }
}

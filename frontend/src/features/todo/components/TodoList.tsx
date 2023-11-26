"use client"
import React from "react"
import { useTodos } from "../hooks/useTodos"
import { TodoForm } from "./TodoForm"
import { TodoItem } from "./TodoItem"
import { useCreateTodo } from "../hooks/useCreateTodo"

export const TodoList = () => {
  const { todos, isLoading, error } = useTodos()
  const { onCreateTodo, isMutating } = useCreateTodo()

  if (isLoading) return <div>Now Loading...</div>
  if (error) return <div>Unexpected Error</div>
  return (
    <div
      style={{
        minHeight: "50%",
        width: "50%",
        border: "1px solid #000",
        padding: "1rem",
      }}
    >
      <h4
        style={{
          margin: "0 0 1rem",
        }}
      >
        やることリスト
      </h4>
      <TodoForm onCreateTodo={onCreateTodo} isMutating={isMutating} />
      <ul
        style={{
          margin: "auto",
          paddingLeft: "2rem",
        }}
      >
        {todos.map((todo) => (
          <TodoItem key={todo.id} todo={todo} />
        ))}
      </ul>
    </div>
  )
}

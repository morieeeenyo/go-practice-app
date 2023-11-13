"use client"
import React from "react"
import { useTodos } from "../hooks/useTodos"
import { TodoItem } from "./TodoItem"

export const TodoList = () => {
  const { todos, isLoading, error } = useTodos()

  if (isLoading) return <div>Now Loading...</div>
  if (error) return <div>Unexpected Error</div>
  return (
    <div
      style={{
        height: "50%",
        width: "50%",
        border: "1px solid #000",
      }}
    >
      <h4>やることリスト</h4>
      <ul
        style={{
          margin: "auto",
        }}
      >
        {todos.map((todo) => (
          <TodoItem key={todo.id} todo={todo} />
        ))}
      </ul>
    </div>
  )
}

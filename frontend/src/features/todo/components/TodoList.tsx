import React from "react"
import { useTodos } from "../hooks/useTodos"
import { TodoItem } from "./TodoItem"

export const TodoList = () => {
  const { todos } = useTodos()
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

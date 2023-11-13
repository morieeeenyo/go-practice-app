import { TodoList } from "@/features/todo/components/TodoList"
import React from "react"

const TodoListPage = () => {
  return (
    <div
      style={{
        width: "100%",
        height: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <TodoList />
    </div>
  )
}

export default TodoListPage

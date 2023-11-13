"use client"
import React from "react"
import { Todo } from "../types"

type Props = {
  todo: Todo
}

export const TodoItem: React.FC<Props> = ({ todo }: Props) => {
  return (
    <li>
      <span>{todo.title}</span>
      <input type="checkbox" checked={todo.isCompleted} onChange={() => null} />
    </li>
  )
}

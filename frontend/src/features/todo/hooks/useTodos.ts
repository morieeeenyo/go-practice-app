import { Todo } from "../types"

export const useTodos = () => {
  const todos: Todo[] = [
    {
      id: 1,
      title: "todo1",
      isCompleted: false,
    },
    {
      id: 2,
      title: "todo2",
      isCompleted: false,
    },
    {
      id: 3,
      title: "todo3",
      isCompleted: true,
    },
  ]

  return {
    todos
  }
}
import React from "react"

export const TodoForm = () => {
  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        padding: "0 0 1rem",
      }}
    >
      <input
        type="text"
        placeholder="TODOのタイトルを入力してください"
        style={{
          marginRight: "2rem",
          minWidth: "20rem",
        }}
      />
      <button>登録</button>
    </div>
  )
}

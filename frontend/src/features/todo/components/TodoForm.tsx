import React from "react"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"

const validationSchema = z.object({
  title: z.string().refine((value) => Boolean(value.trim().length), {
    message: "タイトルを入力してください",
  }),
})

type FormValues = z.infer<typeof validationSchema>

export const TodoForm = () => {
  const form = useForm<FormValues>({
    mode: "onChange",
    resolver: zodResolver(validationSchema),
  })

  const {
    register,
    formState: { errors, isValid },
  } = form
  return (
    <div
      style={{
        display: "flex",
        alignItems: "flex-start",
        padding: "0 0 1rem",
      }}
    >
      <div>
        <input
          {...register("title")}
          name="title"
          type="text"
          placeholder="TODOのタイトルを入力してください"
          style={{
            marginRight: "2rem",
            minWidth: "20rem",
            display: "block",
          }}
        />
        {errors.title && (
          <span
            style={{
              color: "red",
              fontSize: "0.75rem",
            }}
          >
            {errors.title?.message}
          </span>
        )}
      </div>
      <button disabled={!isValid}>登録</button>
    </div>
  )
}

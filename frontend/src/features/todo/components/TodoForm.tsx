import React from "react"
import { useForm, Controller } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"

const validationSchema = z.object({
  title: z.string().refine((value) => Boolean(value.trim().length), {
    message: "タイトルを入力してください",
  }),
})

export type FormValues = z.infer<typeof validationSchema>

type Props = {
  onCreateTodo: (data: FormValues) => Promise<void>
  isMutating: boolean
}

export const TodoForm: React.FC<Props> = ({ onCreateTodo, isMutating }) => {
  const form = useForm<FormValues>({
    mode: "onChange",
    resolver: zodResolver(validationSchema),
    defaultValues: {
      title: "",
    },
  })

  const {
    reset,
    control,
    register,
    handleSubmit,
    formState: { errors, isValid },
  } = form

  const onSubmit = async (data: FormValues) => {
    await onCreateTodo(data).then(() => {
      reset()
    })
  }

  return (
    <form
      style={{
        display: "flex",
        alignItems: "flex-start",
        padding: "0 0 1rem",
      }}
      onSubmit={handleSubmit(onSubmit)}
    >
      <div>
        <Controller
          name="title"
          control={control}
          rules={{ required: true }}
          render={({ field }) => (
            <input
              {...field}
              type="text"
              placeholder="TODOのタイトルを入力してください"
              style={{
                marginRight: "2rem",
                minWidth: "20rem",
                display: "block",
              }}
            />
          )}
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
      <button disabled={!isValid || isMutating}>登録</button>
    </form>
  )
}

import type { Todo as TodoType } from '@/types'

interface TodoProps {
  todo: TodoType
}

export const Todo = ({ todo }: TodoProps) => {
  return (
    <li>
      <p>{todo.id}</p>
      <p>{todo.title}</p>
      <p>{todo.done}</p>
      <p>{todo.priority}</p>
      <p>{todo.due_date}</p>
      <p>{todo.created_at}</p>
      <p>{todo.updated_at}</p>
    </li>
  )
}

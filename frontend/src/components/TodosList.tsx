import { useLoaderData } from 'react-router-dom'
import { Todo } from './Todo'

import type { Todos, Todo as TodoType } from '@/types'

export const TodosList = () => {
  const todosData: Todos = useLoaderData()
  const todos = todosData.todos

  return (
    <>
      {todos.length > 0 && (
        <ul>
          {todos.map((todo: TodoType) => (
            <Todo key={todo.id} todo={todo} />
          ))}
        </ul>
      )}
    </>
  )
}

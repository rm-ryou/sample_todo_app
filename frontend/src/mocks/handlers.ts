import { http, HttpResponse } from 'msw'

import type { Todo, Todos } from '@/types'

const endpoint = import.meta.env.VITE_API_ENDPOINT

const todos: Todo[] = [
  {
    id: 1,
    title: '終了したタスク',
    done: true,
    priority: 0,
    due_date: null,
    created_at: new Date().toISOString(),
    updated_at: new Date().toISOString(),
  },
  {
    id: 2,
    title: '進行中のタスク',
    done: false,
    priority: 2,
    due_date: new Date(2025, 4, 3).toISOString(),
    created_at: new Date(2025, 4, 1).toISOString(),
    updated_at: new Date(2025, 4, 1).toISOString(),
  },
]

export const handlers = [
  http.get(`${endpoint}/v1/todos`, () => {
    return HttpResponse.json<Todos>({ todos: todos })
  }),

  http.get(`${endpoint}/v1/todos/:id`, () => {
    return HttpResponse.json<Todo>({
      id: 1,
      title: '終了したタスク',
      done: true,
      priority: 0,
      due_date: null,
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
    })
  }),
]

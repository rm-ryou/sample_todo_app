export type Todo = {
  id: number
  title: string
  done: boolean
  priority: number
  due_date: string | null
  created_at: string
  updated_at: string
}

export type Todos = {
  todos: Todo[]
}

import type { LoaderFunctionArgs } from 'react-router'

export const todosLoader = async () => {
  const endpoint = import.meta.env.VITE_API_ENDPOINT
  const response = await fetch(`${endpoint}/v1/todos/`)
  const resData = await response.json()

  return resData
}

export const todoDetailsLoader = async ({ params }: LoaderFunctionArgs) => {
  const endpoint = import.meta.env.VITE_API_ENDPOINT
  const response = await fetch(`${endpoint}/v1/todos/${params.id}`)
  const resData = await response.json()

  return resData
}

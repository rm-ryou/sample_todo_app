export const loader = async () => {
  const endpoint = import.meta.env.VITE_API_ENDPOINT
  const response = await fetch(`${endpoint}/v1/todos/`)
  const resData = await response.json()

  return resData
}

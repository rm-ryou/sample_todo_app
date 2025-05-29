import type { Rooms } from '@/types'

export const roomsLoader = async () => {
  const endpoint = import.meta.env.VITE_API_ENDPOINT

  try {
    const response = await fetch(`${endpoint}/v1/rooms`)
    if (!response.ok) {
      // TODO: handle error
    }
    const resData: Rooms = await response.json()
    return resData
  } catch (error) {
    console.error(error)
    // TODO: handle error
  }
}

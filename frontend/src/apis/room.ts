import type { UniqueIdentifier } from '@dnd-kit/core'
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

export const deleteRoom = async (id: UniqueIdentifier) => {
  const endpoint = import.meta.env.VITE_API_ENDPOINT

  try {
    const response = await fetch(`${endpoint}/v1/rooms/${id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
    })
    if (!response.ok) {
      // FIXME: Make the caller catch the error and display a message
      const errorData = await response.json().catch(() => null)
      const errorMessage =
        errorData?.message || response.statusText || 'Unexpected error'
      console.error(
        `Failed to delete room. ${response.status}: ${errorMessage}`,
      )
    }
  } catch (error) {
    console.error(error)
  }
}

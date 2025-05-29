import { http, HttpResponse } from 'msw'

import { rooms } from './data/room'
import type { Rooms } from '@/types'

const endpoint = import.meta.env.VITE_API_ENDPOINT

export const handlers = [
  http.get(`${endpoint}/v1/rooms`, () => {
    return HttpResponse.json<Rooms>({ rooms: rooms })
  }),
]

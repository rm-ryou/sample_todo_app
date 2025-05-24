import { useLoaderData } from 'react-router'
import Room from './Room'
import type { Rooms } from '@/types'

import styles from './RoomsList.module.css'

const RoomsList = () => {
  const roomsData: Rooms = useLoaderData()
  const rooms = roomsData.rooms

  return (
    <>
      {rooms.length > 0 && (
        <ul className={styles.rooms}>
          {rooms.map((room) => (
            <Room key={room.id} room={room} />
          ))}
        </ul>
      )}
    </>
  )
}

export default RoomsList

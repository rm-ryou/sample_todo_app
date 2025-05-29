import { useLoaderData } from 'react-router'
import Room from './Room'
import styles from './RoomsList.module.css'
import type { Rooms } from '@/types'

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

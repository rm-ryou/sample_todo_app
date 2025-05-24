import type { Room as RoomType } from '@/types'

import styles from './Room.module.css'

interface RoomProps {
  room: RoomType
}

const Room = (props: RoomProps) => {
  const { room } = props
  return (
    <li className={styles.room}>
      <p className={styles.name}>{room.name}</p>
    </li>
  )
}

export default Room

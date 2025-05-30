import { useDraggable } from '@dnd-kit/core'

import type { Room as RoomType } from '@/types'
import styles from './Room.module.css'

interface RoomProps {
  room: RoomType
}

const Room = (props: RoomProps) => {
  const { room } = props
  const { attributes, listeners, setNodeRef, transform } = useDraggable({
    id: room.id,
    data: {
      type: 'room',
    },
  })
  const style = transform
    ? {
        transform: `translate3d(${transform.x}px, ${transform.y}px, 0)`,
      }
    : undefined

  return (
    <div ref={setNodeRef} style={style} {...listeners} {...attributes}>
      <li className={styles.room}>
        <p className={styles.name}>{room.name}</p>
      </li>
    </div>
  )
}

export default Room

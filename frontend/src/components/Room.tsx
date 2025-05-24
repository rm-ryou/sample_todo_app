import { useState } from 'react'

import type { Room as RoomType } from '@/types'
import styles from './Room.module.css'

interface RoomProps {
  room: RoomType
}

const Room = (props: RoomProps) => {
  const { room } = props
  const [isEditing, setIsEditing] = useState(false)
  const [name, setName] = useState(room.name)

  const handleDoubleClick = () => {
    setIsEditing(true)
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value)
  }

  const handleBlur = () => {
    setIsEditing(false)
    // TODO: send data
  }

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      setIsEditing(false)
      // TODO: send data
    }
  }

  return (
    <li className={styles.room}>
      {isEditing ? (
        <input
          type='text'
          autoFocus
          value={name}
          className={styles.roominput}
          onChange={handleChange}
          onBlur={handleBlur}
          onKeyDown={handleKeyDown}
        />
      ) : (
        <p className={styles.name} onDoubleClick={handleDoubleClick}>
          {name}
        </p>
      )}
    </li>
  )
}

export default Room

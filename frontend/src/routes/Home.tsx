import { Outlet } from 'react-router'
import { DndContext } from '@dnd-kit/core'
import type { DragEndEvent } from '@dnd-kit/core'
import { deleteRoom } from '@/apis/room'

import Header from '@/components/Header'

const Home = () => {
  const handleDragEnd = (e: DragEndEvent) => {
    const { active, over } = e
    if (over && over.id === 'trashArea') {
      // TODO: Switch process by room, board, todo
      deleteRoom(active.id)
    }
  }

  return (
    <DndContext onDragEnd={handleDragEnd}>
      <Header />
      <Outlet />
    </DndContext>
  )
}

export default Home

import { Outlet } from 'react-router'
import { DndContext } from '@dnd-kit/core'
import type { DragEndEvent } from '@dnd-kit/core'

import Header from '@/components/Header'

const Home = () => {
  const handleDragEnd = (e: DragEndEvent) => {
    if (e.over && e.over.id === 'trashArea') {
      console.log('To be trashed!!')
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

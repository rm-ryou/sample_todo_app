import { Outlet } from 'react-router'

import RoomsList from '@/components/RoomsList'

const Rooms = () => {
  return (
    <>
      <RoomsList />
      <Outlet />
    </>
  )
}

export default Rooms

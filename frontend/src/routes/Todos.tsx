import { Outlet } from 'react-router-dom'
import { TodosList } from '../components/TodosList'

export const Todos = () => {
  return (
    <>
      <TodosList />
      <Outlet />
    </>
  )
}

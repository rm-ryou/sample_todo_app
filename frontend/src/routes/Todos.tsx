import { Outlet } from 'react-router'

import TodosList from '../components/TodosList'

const Todos = () => {
  return (
    <>
      <TodosList />
      <Outlet />
    </>
  )
}

export default Todos

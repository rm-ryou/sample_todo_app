import { Outlet } from 'react-router'

import Header from '@/components/Header'

const Home = () => {
  return (
    <>
      <Header />
      <Outlet />
    </>
  )
}

export default Home

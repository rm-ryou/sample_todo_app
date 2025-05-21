import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router'

import Home from './routes/Home'
import Todos from './routes/Todos'
import TodoDetails from './routes/TodoDetails'
import { todosLoader, todoDetailsLoader } from './apis/todo'
import { enableMocking } from './mocks/browser'
import './index.css'

// NOTE: Since the loader in react-router runs before the worker is configured,
// the worker is started in browser.ts
enableMocking().then(() => {
  const router = createBrowserRouter([
    {
      path: '/',
      element: <Home />,
      children: [
        {
          path: '/',
          element: <Todos />,
          loader: todosLoader,
        },
        {
          path: '/:id',
          element: <TodoDetails />,
          loader: todoDetailsLoader,
        },
      ],
    },
  ])

  createRoot(document.getElementById('root')!).render(
    <StrictMode>
      <RouterProvider router={router} />
    </StrictMode>,
  )
})

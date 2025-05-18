import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router'

import { Todos } from './routes/Todos'
import { loader as todosLoader } from './apis/todo'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Todos />,
    loader: todosLoader,
    children: [
      {
        path: '/new',
        element: '<h1>Hello World!</h1>',
      },
    ],
  },
])

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)

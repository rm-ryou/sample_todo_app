import { useLoaderData } from 'react-router'
import { Box, Grid } from '@mui/material'

import Todo from './Todo'
import type { Todos, Todo as TodoType } from '@/types'

const TodosList = () => {
  const todosData: Todos = useLoaderData()
  const todos = todosData.todos

  return (
    <Box sx={{ flexGrow: 1, paddig: 3 }}>
      {todos.length > 0 && (
        <Grid
          container
          spacing={{ xs: 2, md: 3 }}
          columns={{ xs: 4, sm: 8, md: 12 }}
          sx={{
            mt: 2,
            justifyContent: 'center',
            alignItems: 'center',
          }}
        >
          {todos.map((todo: TodoType) => (
            <Todo key={todo.id} todo={todo} />
          ))}
        </Grid>
      )}
    </Box>
  )
}

export default TodosList

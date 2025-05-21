import { useLoaderData, useNavigate } from 'react-router'
import { Box, Typography, Modal } from '@mui/material'

import type { Todo } from '@/types'

const TodoDetails = () => {
  const todo: Todo = useLoaderData()
  const navigate = useNavigate()
  const handleClose = () => {
    navigate('..')
  }

  return (
    <Modal
      open={true}
      onClose={handleClose}
      aria-labelledby='modal-modal-title'
      aria-describedby='modal-modal-description'
    >
      <Box
        sx={{
          position: 'absolute',
          top: '50%',
          left: '50%',
          color: '#000000',
          transform: 'translate(-50%, -50%)',
          width: 400,
          bgcolor: 'white',
          boxShadow: 12,
          p: 4,
        }}
      >
        <Typography id='modal-modal-title' variant='h6' component='h2'>
          {todo.title}
        </Typography>
      </Box>
    </Modal>
  )
}

export default TodoDetails

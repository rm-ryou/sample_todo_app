import { useState } from 'react'
import { Link } from 'react-router'
import {
  Box,
  Card,
  CardContent,
  IconButton,
  Menu,
  MenuItem,
} from '@mui/material'
import MoreVertIcon from '@mui/icons-material/MoreVert'

import type { Todo as TodoType } from '@/types'

interface TodoProps {
  todo: TodoType
}

const Todo = ({ todo }: TodoProps) => {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null)
  const open = Boolean(anchorEl)

  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget)
  }

  const handleClose = () => {
    setAnchorEl(null)
  }

  return (
    <Card sx={{ width: '100%', marginBottom: 1, boxShadow: 2 }}>
      <Box
        sx={{
          display: 'flex',
          justifyContent: 'space-between',
          alignItems: 'flex-start',
          pr: 1,
        }}
      >
        <CardContent>{todo.title}</CardContent>

        <IconButton
          onClick={handleClick}
          sx={{ mt: 1, mr: 1 }}
          aria-controls={open ? `todo-menu-${todo.id}` : undefined}
          aria-haspopup='true'
          aria-expanded={open ? 'true' : undefined}
        >
          <MoreVertIcon />
        </IconButton>
        <Menu
          anchorEl={anchorEl}
          id={`todo-menu-${todo.id}`}
          open={open}
          onClose={handleClose}
          onClick={handleClose}
        >
          <MenuItem onClick={handleClose} component={Link} to={`${todo.id}`}>
            Details
          </MenuItem>
          {/* TODO: patch request */}
          <MenuItem onClick={handleClose}>Mark as finished</MenuItem>
          {/* TODO: delete request */}
          <MenuItem onClick={handleClose}>Delete</MenuItem>
        </Menu>
      </Box>
    </Card>
  )
}

export default Todo

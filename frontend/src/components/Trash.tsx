import { useDroppable } from '@dnd-kit/core'
import DeleteIcon from '@mui/icons-material/Delete'

const Trash = () => {
  const { setNodeRef } = useDroppable({
    id: 'trashArea',
  })

  return (
    <div ref={setNodeRef}>
      <DeleteIcon fontSize='large' />
    </div>
  )
}

export default Trash

import { Link } from 'react-router'
import DeleteIcon from '@mui/icons-material/Delete'

import styles from './Header.module.css'

const Header = () => {
  return (
    <header className={styles.header}>
      <Link to='/' className={styles.logo}>
        <h1>Sample todo app</h1>
      </Link>
      <DeleteIcon fontSize='large' />
    </header>
  )
}

export default Header

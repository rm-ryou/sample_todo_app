import { Link } from 'react-router'
import Trash from './Trash'

import styles from './Header.module.css'

const Header = () => {
  return (
    <header className={styles.header}>
      <Link to='/' className={styles.logo}>
        <h1>Sample todo app</h1>
      </Link>
      <Trash />
    </header>
  )
}

export default Header

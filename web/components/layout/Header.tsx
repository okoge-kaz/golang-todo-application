import Link from 'next/link'
import styles from './style.module.scss'

const Header = () => {
  return (
    <header>
      <div className={styles.header_bar}>
        <button className={styles.button}>
          <Link href="/user/new">
            Sign Up
          </Link>
        </button>
        <button className={styles.button}>
          <Link href="/login">
            Sign In
          </Link>
        </button>
      </div>
    </header>
  )
}

export default Header

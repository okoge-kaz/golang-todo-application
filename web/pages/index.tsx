import Header from 'components/layout/Header'
import Head from 'next/head'
import Link from 'next/link'
import styles from 'styles/Home.module.css'

const Home = () => {
  return (
    <div className="container mx-auto">
      <Head>
        <title>Todo Application</title>
        <meta name="description" content="todo application" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Header />

      <main className={styles.main}>
        <h1 className="text-4xl font-bold text-center mt-10">
          Welcome to Todo Application
        </h1>

        <div>
          <div className={styles.description}>
            <p>Tokyo Institute of Technology B3 (2022 3Q)</p>
            <p>Workshop on System Design</p>
            <p>Todo Application (Backend: Go Gin, Frontend: Typescript Next.js)</p>
          </div>

          <div className={styles.grid}>
            <Link href="/task">
              <div className={styles.card}>
                <h3>Tasks &rarr;</h3>
                <p>You can view all tasks.</p>
              </div>
            </Link>

            <Link href="/task">
              <div className={styles.card}>
                <h3>Categories &rarr;</h3>
                <p>You can view all categories.</p>
              </div>
            </Link>

            {/* 以下はログイン済みのときのみ表示 */}
            <Link href="/user/change_password">
              <div className={styles.card}>
                <h3>Change Account Password &rarr;</h3>
              </div>
            </Link>

            <Link href="/user/change_name">
              <div className={styles.card}>
                <h3>Change Account Username &rarr;</h3>
              </div>
            </Link>

            <Link href="/task/new">
              <div className={styles.card}>
                <h3>New Task &rarr;</h3>
              </div>
            </Link>

            <Link href="/category/new">
              <div className={styles.card}>
                <h3>New Categories &rarr;</h3>
              </div>
            </Link>

            <Link href="/user/delete">
              <div className={styles.card}>
                <h3>Delete Account &rarr;</h3>
              </div>
            </Link>

          </div>
        </div>
      </main>
    </div>
  )
}

export default Home

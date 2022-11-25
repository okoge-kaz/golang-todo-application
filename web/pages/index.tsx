import Header from 'components/layout/Header'
import Head from 'next/head'

const Home = () => {
  return (
    <div className="container mx-auto">
      <Head>
        <title>Todo Application</title>
        <meta name="description" content="todo application" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Header />

      <main>
        <h1 className="text-4xl font-bold text-center mt-10">
          Welcome to Todo Application
        </h1>
      </main>
    </div>
  )
}

export default Home

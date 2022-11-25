import Header from 'components/layout/Header'
import Task from 'interfaces/task'
import { GetServerSideProps } from "next"
import router from 'next/router'

type Props = {
  task: Task
}

const Task = (props: Props) => {
  return (
    <div className="container mx-auto">
      <Header />

      <main>
        <h1 className="text-4xl font-bold text-center mt-10">
          Task
        </h1>

        <div>
          <h2 className="text-2xl font-bold">{props.task.title}</h2>
          <p>{props.task.description}</p>
          <p>{props.task.status}</p>
        </div>

      </main>
    </div>
  )
}

export default Task

export const getServerSideProps: GetServerSideProps = async () => {
  const id = router.query.id
  const res = await fetch(`http://localhost:8000/task/${id}`)
  const task: Task = await res.json()

  return {
    props: {
      task,
    },
  }
}

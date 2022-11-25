import Header from 'components/layout/Header'
import Task from 'interfaces/task'
import { GetServerSideProps } from "next"
import Link from 'next/link'

type Props = {
  tasks: Task[]
}


const TaskList = (props: Props) => {

  return (
    <div className="container mx-auto">
      <Header />

      <main>
        <h1 className="text-4xl font-bold text-center mt-10">
          Task List
        </h1>

        <div>
          {props.tasks.map((task) => (
            <Link href={`/task/${task.id}`} key={task.id}>
              <div key={task.id}>
                <h2 className="text-2xl font-bold">{task.title}</h2>
                <p>{task.description}</p>
              </div>
            </Link>
          ))}
        </div>

        <div>
          <Link href="/task/new">
            <a className="text-blue-500">New Task</a>
          </Link>
        </div>

      </main>
    </div>
  )
}

export default TaskList

export const getServerSideProps: GetServerSideProps = async () => {
  const res = await fetch('http://localhost:8000/task')
  const tasks: Task[] = await res.json()

  return {
    props: {
      tasks,
    },
  }
}

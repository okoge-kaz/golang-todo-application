export default interface Task {
  id: number
  title: string
  deadline: Date
  status: string
  description: string
  is_done: boolean
  category: string
  created_at: Date
  updated_at: Date
}

import { Todo } from "../entities/Todo"

type Props = {
  todos: Todo[]
}

export const TodoList: React.FC<Props> = ({ todos }) => {
  return(
    <ul className="todo-list">
      {
        todos.map((todo, i) => {
          return(
            <li key={i}>{todo.title}</li>
          )
        })
      }
    </ul>
  )
}
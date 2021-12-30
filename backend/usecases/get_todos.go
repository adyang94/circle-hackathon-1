package usecases

import "github.com/adyang94/react-go-todo-app/backend/entities"

func GetTodos(repo TodosRepository) ([]entities.Todo, error) {
	todos, err := repo.GetAllTodos()
	if err != nil {
		return nil, ErrInternal

	}
	return todos, nil
}

package usecases

import "github.com/adyang94/react-go-todo-app/backend/entities"

type TodosRepository interface {
	GetAllTodos () ([]entities.Todo, error)
}
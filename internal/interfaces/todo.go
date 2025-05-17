package interfaces

import (
	"time"

	"github.com/rm-ryou/sample_todo_app/internal/entities"
)

type TodoRepository interface {
	GetAll() ([]*entities.Todo, error)
	GetById(id int) (*entities.Todo, error)
	Create(todo *entities.Todo) error
	Update(todo *entities.Todo) error
	Delete(id int) error
}

type TodoServicer interface {
	GetAll() ([]*entities.Todo, error)
	GetById(id int) (*entities.Todo, error)
	Create(title string, done bool, priority int, dueDate *time.Time) error
	Update(id int, title string, done bool, priority int, dueDate *time.Time) error
	Delete(id int) error
}

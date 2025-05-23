package services

import (
	"time"

	"github.com/rm-ryou/sample_todo_app/internal/entities"
	"github.com/rm-ryou/sample_todo_app/internal/interfaces"
)

type TodoService struct {
	repo interfaces.TodoRepository
}

func NewTodoService(repo interfaces.TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (ts *TodoService) GetAll() ([]*entities.Todo, error) {
	return ts.repo.GetAll()
}

func (ts *TodoService) GetById(id int) (*entities.Todo, error) {
	return ts.repo.GetById(id)
}

func (ts *TodoService) Create(title string, done bool, priority int, dueDate *time.Time) error {
	todo := entities.NewTodo(title, done, priority, dueDate)
	if err := todo.Validate(); err != nil {
		return err
	}

	return ts.repo.Create(todo)
}

func (ts *TodoService) Update(id int, title string, done bool, priority int, dueDate *time.Time) error {
	todo, err := ts.repo.GetById(id)
	if err != nil {
		return err
	}

	todo.UpdateAttributes(title, done, priority, dueDate)
	if err := todo.Validate(); err != nil {
		return err
	}

	return ts.repo.Update(todo)
}

func (ts *TodoService) Delete(id int) error {
	if _, err := ts.repo.GetById(id); err != nil {
		return err
	}

	return ts.repo.Delete(id)
}

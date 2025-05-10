package todo

import "github.com/rm-ryou/sample_todo_app/internal/entity"

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) CreateTodo(todo *entity.Todo) error {
	if err := todo.Validate(); err != nil {
		return err
	}
	return s.r.Create(todo)
}

func (s *Service) GetTodo(id int) (*entity.Todo, error) {
	return s.r.Get(id)
}

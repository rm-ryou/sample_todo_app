package todo

import "github.com/rm-ryou/sample_todo_app/internal/entity"

type TodoService struct {
	r Repository
}

func NewService(r Repository) *TodoService {
	return &TodoService{
		r: r,
	}
}

func (ts *TodoService) GetAll() ([]*entity.Todo, error) {
	return ts.r.GetAll()
}

func (ts *TodoService) GetById(id int) (*entity.Todo, error) {
	return ts.r.GetById(id)
}

func (ts *TodoService) Create(todo *entity.Todo) error {
	if err := todo.Validate(); err != nil {
		return err
	}
	return ts.r.Create(todo)
}

func (ts *TodoService) Update(todo *entity.Todo) error {
	_, err := ts.r.GetById(todo.Id)
	if err != nil {
		return err
	}

	if err := todo.Validate(); err != nil {
		return err
	}
	return ts.r.Update(todo)
}

func (ts *TodoService) Delete(id int) error {
	_, err := ts.r.GetById(id)
	if err != nil {
		return err
	}

	return ts.r.Delete(id)
}

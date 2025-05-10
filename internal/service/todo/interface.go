package todo

import "github.com/rm-ryou/sample_todo_app/internal/entity"

type Getter interface {
	Get(id int) (*entity.Todo, error)
}

type Modifier interface {
	Create(todo *entity.Todo) error
}

type Repository interface {
	Getter
	Modifier
}

type Servicer interface {
	GetTodo(id int) (*entity.Todo, error)
	CreateTodo(todo *entity.Todo) error
}

package interfaces

import "github.com/rm-ryou/sample_todo_app/internal/entities"

type BoardRepository interface {
	GetAll() ([]*entities.Board, error)
	GetById(id int) (*entities.Board, error)
	Create(board *entities.Board) error
	Update(board *entities.Board) error
	Delete(id int) error
}

type BoardServicer interface {
	GetAll() ([]*entities.Board, error)
	Create(name string, priority, roomId int) error
	Update(id int, name string, priority int) error
	Delete(id int) error
}

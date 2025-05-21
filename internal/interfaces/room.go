package interfaces

import (
	"github.com/rm-ryou/sample_todo_app/internal/entities"
)

type RoomRepository interface {
	GetAll() ([]*entities.Room, error)
	Create(room *entities.Room) error
	Update(room *entities.Room) error
	Delete(id int) error
}

type RoomServicer interface {
	GetAll() ([]*entities.Room, error)
	Create(name string) error
	Update(id int, name string) error
	Delete(id int) error
}

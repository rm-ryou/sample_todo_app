package services

import (
	"github.com/rm-ryou/sample_todo_app/internal/entities"
	"github.com/rm-ryou/sample_todo_app/internal/interfaces"
)

type RoomService struct {
	repo interfaces.RoomRepository
}

func NewRoomService(repo interfaces.RoomRepository) *RoomService {
	return &RoomService{
		repo: repo,
	}
}

func (rs *RoomService) GetAll() ([]*entities.Room, error) {
	return rs.repo.GetAll()
}

func (rs *RoomService) Create(name string) error {
	room := entities.NewRoom(name)
	if err := room.Validate(); err != nil {
		return err
	}

	return rs.repo.Create(room)
}

func (rs *RoomService) Update(id int, name string) error {
	room, err := rs.repo.GetById(id)
	if err != nil {
		return err
	}

	room.UpdateAttributes(name)
	if err := room.Validate(); err != nil {
		return err
	}

	return rs.repo.Update(room)
}

func (rs *RoomService) Delete(id int) error {
	if _, err := rs.repo.GetById(id); err != nil {
		return err
	}

	return rs.repo.Delete(id)
}

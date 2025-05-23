package services

import (
	"github.com/rm-ryou/sample_todo_app/internal/entities"
	"github.com/rm-ryou/sample_todo_app/internal/interfaces"
)

type BoardService struct {
	repo interfaces.BoardRepository
}

func NewBoardService(repo interfaces.BoardRepository) *BoardService {
	return &BoardService{
		repo: repo,
	}
}

func (bs *BoardService) GetAll() ([]*entities.Board, error) {
	return bs.repo.GetAll()
}

func (bs *BoardService) Create(name string, priority, roomId int) error {
	board := entities.NewBoard(name, priority, roomId)
	if err := board.Validate(); err != nil {
		return err
	}

	return bs.repo.Create(board)
}

func (bs *BoardService) Update(id int, name string, priority int) error {
	board, err := bs.repo.GetById(id)
	if err != nil {
		return err
	}

	board.UpdateAttributes(name, priority)
	if err := board.Validate(); err != nil {
		return err
	}

	return bs.repo.Update(board)
}

func (bs *BoardService) Delete(id int) error {
	if _, err := bs.repo.GetById(id); err != nil {
		return err
	}

	return bs.repo.Delete(id)
}

package entities

import (
	"errors"
	"time"
)

type Todo struct {
	Id        int
	BoardId   int
	Title     string
	Done      bool
	Priority  int
	DueDate   *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTodo(boardId int, title string, done bool, priority int, dueDate *time.Time) *Todo {
	return &Todo{
		BoardId:  boardId,
		Title:    title,
		Done:     done,
		Priority: priority,
		DueDate:  dueDate,
	}
}

func (t *Todo) Validate() error {
	if t.Title == "" || len(t.Title) > 50 {
		return errors.New("Invalid title")
	}

	if t.Priority < 0 {
		return errors.New("Invalid priority size")
	}
	return nil
}

func (t *Todo) UpdateAttributes(title string, done bool, priority int, dueDate *time.Time) {
	t.Title = title
	t.Done = done
	t.Priority = priority
	t.DueDate = dueDate
}

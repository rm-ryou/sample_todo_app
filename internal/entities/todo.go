package entities

import (
	"errors"
	"time"
)

type Todo struct {
	Id        int
	Title     string
	Done      bool
	Priority  int
	DueDate   *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTodo(title string, done bool, priority int, dueDate *time.Time) *Todo {
	return &Todo{
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

func (t *Todo) UpdateTitle(title string) error {
	t.Title = title
	return t.Validate()
}

func (t *Todo) UpdateDone(done bool) {
	t.Done = done
}

func (t *Todo) UpdatePriority(priority int) error {
	t.Priority = priority
	return t.Validate()
}

func (t *Todo) UpdateDueDate(dueDate *time.Time) {
	t.DueDate = dueDate
}

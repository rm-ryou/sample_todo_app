package entity

import (
	"errors"
	"time"
)

type Todo struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Done      bool       `json:"done"`
	Priority  int        `json:"priority"`
	DueDate   *time.Time `json:"due_date,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (t *Todo) Validate() error {
	if err := t.validateTitle(); err != nil {
		return err
	}
	return nil
}

// FIXME: use custom error
func (t *Todo) validateTitle() error {
	if t.Title == "" {
		return errors.New("Title is required")
	}

	if len(t.Title) > 50 {
		return errors.New("Title must be 50 characters or less")
	}
	return nil
}

package entity

import "time"

type Todo struct {
	Id        int
	Title     string
	Done      bool
	Priority  int
	DueDate   *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

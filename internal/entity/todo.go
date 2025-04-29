package entity

import "time"

type Todo struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Done      bool       `json:"done"`
	Priority  int        `json:"priority"`
	DueDate   *time.Time `json:"due_date,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

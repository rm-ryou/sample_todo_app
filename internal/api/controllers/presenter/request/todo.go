package request

import "time"

type Todo struct {
	BoardId  int        `json:"board_id" validate:"required"`
	Title    string     `json:"title" validate:"required,max=50"`
	Done     bool       `json:"done"`
	Priority int        `json:"priority"`
	DueDate  *time.Time `json:"due_date,omitempty"`
}

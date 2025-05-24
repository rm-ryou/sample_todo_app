package response

import (
	"time"

	"github.com/rm-ryou/sample_todo_app/internal/entities"
)

type ListTodo struct {
	Todos []*Todo `json:"todos"`
}

type Todo struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Done      bool       `json:"done"`
	Priority  int        `json:"priority"`
	BoardId   int        `json:"board_id"`
	DueDate   *time.Time `json:"due_date,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func ConvertTodoResponse(todo *entities.Todo) *Todo {
	return &Todo{
		Id:        todo.Id,
		Title:     todo.Title,
		Done:      todo.Done,
		Priority:  todo.Priority,
		DueDate:   todo.DueDate,
		BoardId:   todo.BoardId,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

func ConvertoTodosResponse(todos []*entities.Todo) *ListTodo {
	listTodo := []*Todo{}

	for _, todo := range todos {
		listTodo = append(listTodo, ConvertTodoResponse(todo))
	}
	return &ListTodo{Todos: listTodo}
}

package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/rm-ryou/sample_todo_app/internal/entity"
	"github.com/rm-ryou/sample_todo_app/internal/service/todo"
)

type Todo struct {
	s todo.Servicer
}

func NewTodo(s todo.Servicer) *Todo {
	return &Todo{
		s: s,
	}
}

func (t *Todo) CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := &entity.Todo{}

	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err := t.s.CreateTodo(todo)
	if err != nil {
		// TODO: エラーの内容によってステータスコードを変えられるような構造体の定義
		ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	CommonResponse(w, 200, "OK")
}

func (t *Todo) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	todo, err := t.s.GetTodo(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ErrorResponse(w, http.StatusNotFound, err)
			return
		}
		ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err)
	}
}

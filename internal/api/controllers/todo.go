package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/rm-ryou/sample_todo_app/internal/api/controllers/presenter/request"
	"github.com/rm-ryou/sample_todo_app/internal/interfaces"
)

type TodoController struct {
	service interfaces.TodoServicer
}

func NewTodoController(service interfaces.TodoServicer) *TodoController {
	return &TodoController{
		service: service,
	}
}

func (tc *TodoController) GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := tc.service.GetAll()
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if todos == nil {
		ErrorResponse(w, http.StatusNotFound, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err)
	}
}

func (tc *TodoController) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	todo, err := tc.service.GetById(id)
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

func (tc *TodoController) Create(w http.ResponseWriter, r *http.Request) {
	var req request.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err := tc.service.Create(req.Title, req.Done, req.Priority, req.DueDate)
	if err != nil {
		// TODO: エラーの内容によってステータスコードを変えられるような構造体の定義
		ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	CommonResponse(w, 200, "OK")
}

func (tc *TodoController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	var req request.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = tc.service.Update(id, req.Title, req.Done, req.Priority, req.DueDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ErrorResponse(w, http.StatusNotFound, err)
			return
		}
		ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	CommonResponse(w, 200, "OK")
}

func (tc *TodoController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = tc.service.Delete(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ErrorResponse(w, http.StatusNotFound, err)
			return
		}
		ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	CommonResponse(w, 200, "OK")
}

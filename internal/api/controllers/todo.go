package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/rm-ryou/sample_todo_app/internal/api/controllers/presenter/request"
	"github.com/rm-ryou/sample_todo_app/internal/api/controllers/presenter/response"
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

func (tc *TodoController) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	todo, err := tc.service.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(w, http.StatusNotFound, err)
			return
		}
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	res := response.ConvertTodoResponse(todo)
	response.Basic(w, http.StatusOK, res)
}

func (tc *TodoController) Create(w http.ResponseWriter, r *http.Request) {
	boardIdStr := r.PathValue("boardId")
	boardId, err := strconv.Atoi(boardIdStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	var req request.Todo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := tc.service.Create(boardId, req.Title, req.Done, req.Priority, req.DueDate); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.Basic(w, http.StatusOK, response.BasicResponse{Message: "OK"})
}

func (tc *TodoController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	var req request.Todo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	err = tc.service.Update(id, req.Title, req.Done, req.Priority, req.DueDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(w, http.StatusNotFound, err)
			return
		}
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.Basic(w, http.StatusOK, response.BasicResponse{Message: "OK"})
}

func (tc *TodoController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	err = tc.service.Delete(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(w, http.StatusNotFound, err)
			return
		}
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.Basic(w, http.StatusOK, response.BasicResponse{Message: "OK"})
}

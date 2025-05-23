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

type RoomController struct {
	service interfaces.RoomServicer
}

func NewRoomController(service interfaces.RoomServicer) *RoomController {
	return &RoomController{
		service: service,
	}
}

func (rc *RoomController) GetAll(w http.ResponseWriter, r *http.Request) {
	rooms, err := rc.service.GetAll()
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	res := response.ConvertRoomsResponse(rooms)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err)
	}
}

func (rc *RoomController) Create(w http.ResponseWriter, r *http.Request) {
	req := request.Room{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err := rc.service.Create(req.Name)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	CommonResponse(w, 200, "OK")
}

func (rc *RoomController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	req := request.Room{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = rc.service.Update(id, req.Name)
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

func (rc *RoomController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = rc.service.Delete(id)
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

package controllers

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/rm-ryou/sample_todo_app/internal/entities"
	mock_service "github.com/rm-ryou/sample_todo_app/internal/interfaces/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetByIdTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockTodoServicer(ctrl)
	controller := NewTodoController(mockService)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/boards/{boardId}/todos/{id}", controller.GetById)

	testCases := []struct {
		name           string
		idParam        string
		boardIdParam   string
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:         "Success to Get todo",
			idParam:      "1",
			boardIdParam: "1",
			setupMock: func() {
				mockService.EXPECT().GetById(1).
					Return(&entities.Todo{
						Id:        1,
						Title:     "test",
						Done:      false,
						Priority:  1,
						DueDate:   nil,
						BoardId:   1,
						CreatedAt: time.Date(2025, 5, 1, 10, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2025, 5, 1, 10, 0, 0, 0, time.UTC),
					}, nil)
			},
			expectedStatus: 200,
			expectedBody: `{
				"id":1,
				"title":"test",
				"done":false,
				"priority":1,
				"board_id":1,
				"created_at":"2025-05-01T10:00:00Z",
				"updated_at":"2025-05-01T10:00:00Z"
			}`,
		},
		{
			name:           "Failed with invalid request - Due to non-numeric id",
			idParam:        "invalid",
			boardIdParam:   "1",
			setupMock:      func() {},
			expectedStatus: 400,
			expectedBody:   `{"message":"Bad Request"}`,
		},
		{
			name:         "Failed with not found - Due to no todo with id",
			idParam:      "999",
			boardIdParam: "1",
			setupMock: func() {
				mockService.EXPECT().GetById(999).
					Return(nil, sql.ErrNoRows)
			},
			expectedStatus: 404,
			expectedBody:   `{"message":"Not Found"}`,
		},
		{
			name:         "Failed with internal server error - Due to unexpected errors",
			idParam:      "1",
			boardIdParam: "1",
			setupMock: func() {
				mockService.EXPECT().GetById(1).
					Return(nil, errors.New("unexpected error"))
			},
			expectedStatus: 500,
			expectedBody:   `{"message":"Internal Server Error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			path := "/v1/boards/" + tc.boardIdParam + "/todos/" + tc.idParam
			req := httptest.NewRequest(http.MethodGet, path, nil)
			res := httptest.NewRecorder()

			mux.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.JSONEq(t, tc.expectedBody, res.Body.String())
		})
	}
}

func TestCreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockTodoServicer(ctrl)
	controller := NewTodoController(mockService)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/boards/{boardId}/todos", controller.Create)

	testCases := []struct {
		name           string
		boardIdParam   string
		requestBody    string
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:         "Success to Create new todo",
			boardIdParam: "1",
			requestBody:  `{"title":"TestTodo","done":false,"priority":0,"board_id":1}`,
			setupMock: func() {
				mockService.EXPECT().Create(1, "TestTodo", false, 0, nil).Return(nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"message":"OK"}`,
		},
		{
			name:           "Failed with bad request - Due to number of characters in the title is more than 50",
			boardIdParam:   "1",
			requestBody:    fmt.Sprintf(`{"title":"%s","done":false,"priority":0,"board_id":1}`, strings.Repeat("a", 51)),
			setupMock:      func() {},
			expectedStatus: 400,
			expectedBody:   `{"message":"Bad Request"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			path := "/v1/boards/" + tc.boardIdParam + "/todos"
			body := bytes.NewBufferString(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, path, body)
			res := httptest.NewRecorder()

			mux.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.JSONEq(t, tc.expectedBody, res.Body.String())
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockTodoServicer(ctrl)
	controller := NewTodoController(mockService)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/boards/{boardId}/todos/{id}", controller.Update)

	testCases := []struct {
		name           string
		idParam        string
		boardIdParam   string
		requestBody    string
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:         "Success to Update todo",
			idParam:      "1",
			boardIdParam: "1",
			requestBody:  `{"title":"UpdateTitle!","done":true,"priority":0,"board_id":1}`,
			setupMock: func() {
				mockService.EXPECT().Update(1, "UpdateTitle!", true, 0, nil).
					Return(nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"message":"OK"}`,
		},
		{
			name:           "Failed with invalid request - Due to non-numeric id",
			idParam:        "invalid",
			boardIdParam:   "1",
			requestBody:    `{"title":"UpdateTitle!","done":true,"priority":0,"board_id":1}`,
			setupMock:      func() {},
			expectedStatus: 400,
			expectedBody:   `{"message":"Bad Request"}`,
		},
		{
			name:           "Failed with bad request - Due to number of characters in the title is more than 50",
			idParam:        "1",
			boardIdParam:   "1",
			requestBody:    fmt.Sprintf(`{"title":"%s","done":true,"priority":0,"board_id":1}`, strings.Repeat("a", 51)),
			setupMock:      func() {},
			expectedStatus: 400,
			expectedBody:   `{"message":"Bad Request"}`,
		},
		{
			name:         "Failed with not found - Due to no todo with id",
			idParam:      "999",
			boardIdParam: "1",
			requestBody:  `{"title":"UpdateTitle!","done":false,"priority":1,"board_id":1}`,
			setupMock: func() {
				mockService.EXPECT().Update(999, "UpdateTitle!", false, 1, nil).
					Return(sql.ErrNoRows)
			},
			expectedStatus: 404,
			expectedBody:   `{"message":"Not Found"}`,
		},
		{
			name:         "Failed with internal server error - Due to unexpected errors",
			idParam:      "1",
			boardIdParam: "1",
			requestBody:  `{"title":"UpdateTitle!","done":false,"priority":1,"board_id":1}`,
			setupMock: func() {
				mockService.EXPECT().Update(1, "UpdateTitle!", false, 1, nil).
					Return(errors.New("unexpected error"))
			},
			expectedStatus: 500,
			expectedBody:   `{"message":"Internal Server Error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			path := "/v1/boards/" + tc.boardIdParam + "/todos/" + tc.idParam
			body := bytes.NewBufferString(tc.requestBody)
			req := httptest.NewRequest(http.MethodPut, path, body)
			res := httptest.NewRecorder()

			mux.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.JSONEq(t, tc.expectedBody, res.Body.String())
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockTodoServicer(ctrl)
	controller := NewTodoController(mockService)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/boards/{boardId}/todos/{id}", controller.Delete)

	testCases := []struct {
		name           string
		idParam        string
		boardIdParam   string
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:         "Success to Delete todo",
			idParam:      "1",
			boardIdParam: "1",
			setupMock: func() {
				mockService.EXPECT().Delete(1).
					Return(nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"message":"OK"}`,
		},
		{
			name:           "Failed with invalid request - Due to non-numeric id",
			idParam:        "invalid",
			boardIdParam:   "1",
			setupMock:      func() {},
			expectedStatus: 400,
			expectedBody:   `{"message":"Bad Request"}`,
		},
		{
			name:         "Failed with not found - Due to no todo with id",
			idParam:      "999",
			boardIdParam: "1",
			setupMock: func() {
				mockService.EXPECT().Delete(999).
					Return(sql.ErrNoRows)
			},
			expectedStatus: 404,
			expectedBody:   `{"message":"Not Found"}`,
		},
		{
			name:         "Failed with internal server error - Due to unexpected errors",
			idParam:      "1",
			boardIdParam: "1",
			setupMock: func() {
				mockService.EXPECT().Delete(1).
					Return(errors.New("unexpected error"))
			},
			expectedStatus: 500,
			expectedBody:   `{"message":"Internal Server Error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			path := "/v1/boards/" + tc.boardIdParam + "/todos/" + tc.idParam
			req := httptest.NewRequest(http.MethodDelete, path, nil)
			res := httptest.NewRecorder()

			mux.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.JSONEq(t, tc.expectedBody, res.Body.String())
		})
	}
}

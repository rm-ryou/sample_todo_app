package handler

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

	"github.com/rm-ryou/sample_todo_app/internal/entity"
	mock_service "github.com/rm-ryou/sample_todo_app/internal/service/todo/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockServicer(ctrl)

	testCases := []struct {
		name           string
		requestBody    string
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "Success to Create new todo",
			requestBody: `{"title":"TestTodo"}`,
			setupMock: func() {
				mockService.EXPECT().CreateTodo(gomock.Any()).Return(nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"message":"OK"}`,
		},
		{
			name:        "Failed with bad request - Due to number of characters in the title is more than 50",
			requestBody: fmt.Sprintf(`{"title":"%s"}`, strings.Repeat("a", 51)),
			setupMock: func() {
				mockService.EXPECT().CreateTodo(gomock.Any()).
					// FIXME: use custom error
					Return(errors.New("Title must be 50 characters or less"))
			},
			// TODO: response 400 error
			expectedStatus: 500,
			expectedBody:   `{"message":"Internal Server Error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			handler := NewTodo(mockService)

			mux := http.NewServeMux()
			mux.HandleFunc("POST /v1/todos", handler.Create)

			body := bytes.NewBufferString(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/v1/todos", body)
			res := httptest.NewRecorder()

			mux.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.JSONEq(t, tc.expectedBody, res.Body.String())
		})
	}
}

func TestGetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockServicer(ctrl)

	testCases := []struct {
		name           string
		idParam        string
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "Success to Get todo",
			idParam: "1",
			setupMock: func() {
				mockService.EXPECT().GetTodo(1).
					Return(&entity.Todo{
						Id:        1,
						Title:     "test",
						Done:      false,
						Priority:  1,
						DueDate:   nil,
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
				"created_at":"2025-05-01T10:00:00Z",
				"updated_at":"2025-05-01T10:00:00Z"
			}`,
		},
		{
			name:           "Failed with invalid request - Due to non-numeric id",
			idParam:        "invalid",
			setupMock:      func() {},
			expectedStatus: 400,
			expectedBody:   `{"message":"Bad Request"}`,
		},
		{
			name:    "Failed with not found - Due to no todo with id",
			idParam: "999",
			setupMock: func() {
				mockService.EXPECT().GetTodo(999).
					Return(nil, sql.ErrNoRows)
			},
			expectedStatus: 404,
			expectedBody:   `{"message":"Not Found"}`,
		},
		{
			name:    "Failed with internal server error - Due to unexpected errors",
			idParam: "1",
			setupMock: func() {
				mockService.EXPECT().GetTodo(1).
					Return(nil, errors.New("unexpected error"))
			},
			expectedStatus: 500,
			expectedBody:   `{"message":"Internal Server Error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			handler := NewTodo(mockService)

			mux := http.NewServeMux()
			mux.HandleFunc("GET /v1/todos/{id}", handler.GetById)

			req := httptest.NewRequest(http.MethodGet, "/v1/todos/"+tc.idParam, nil)
			res := httptest.NewRecorder()

			mux.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.JSONEq(t, tc.expectedBody, res.Body.String())
		})
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockServicer(ctrl)

	testCases := []struct {
		name           string
		idParam        string
		requestBody    string
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "Success to Update todo",
			idParam:     "1",
			requestBody: `{"title":"UpdateTitle!"}`,
			setupMock: func() {
				paramTodo := &entity.Todo{
					Id:    1,
					Title: "UpdateTitle!",
				}
				mockService.EXPECT().UpdateTodo(paramTodo).
					Return(nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"message":"OK"}`,
		},
		{
			name:           "Failed with invalid request - Due to non-numeric id",
			idParam:        "invalid",
			requestBody:    `{"title":"UpdateTitle!"}`,
			setupMock:      func() {},
			expectedStatus: 400,
			expectedBody:   `{"message":"Bad Request"}`,
		},
		{
			name:        "Failed with not found - Due to no todo with id",
			idParam:     "999",
			requestBody: `{"title":"UpdateTitle!"}`,
			setupMock: func() {
				paramTodo := &entity.Todo{
					Id:    999,
					Title: "UpdateTitle!",
				}
				mockService.EXPECT().UpdateTodo(paramTodo).
					Return(sql.ErrNoRows)
			},
			expectedStatus: 404,
			expectedBody:   `{"message":"Not Found"}`,
		},
		{
			name:        "Failed with internal server error - Due to unexpected errors",
			idParam:     "1",
			requestBody: `{"title":"UpdateTitle!"}`,
			setupMock: func() {
				paramTodo := &entity.Todo{
					Id:    1,
					Title: "UpdateTitle!",
				}
				mockService.EXPECT().UpdateTodo(paramTodo).
					Return(errors.New("unexpected error"))
			},
			expectedStatus: 500,
			expectedBody:   `{"message":"Internal Server Error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			handler := NewTodo(mockService)

			mux := http.NewServeMux()
			mux.HandleFunc("PUT /v1/todos/{id}", handler.Update)

			body := bytes.NewBufferString(tc.requestBody)
			req := httptest.NewRequest(http.MethodPut, "/v1/todos/"+tc.idParam, body)
			res := httptest.NewRecorder()

			mux.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.JSONEq(t, tc.expectedBody, res.Body.String())
		})
	}
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockServicer(ctrl)

	testCases := []struct {
		name           string
		idParam        string
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "Success to Delete todo",
			idParam: "1",
			setupMock: func() {
				mockService.EXPECT().DeleteTodo(1).
					Return(nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"message":"OK"}`,
		},
		{
			name:           "Failed with invalid request - Due to non-numeric id",
			idParam:        "invalid",
			setupMock:      func() {},
			expectedStatus: 400,
			expectedBody:   `{"message":"Bad Request"}`,
		},
		{
			name:    "Failed with not found - Due to no todo with id",
			idParam: "999",
			setupMock: func() {
				mockService.EXPECT().DeleteTodo(999).
					Return(sql.ErrNoRows)
			},
			expectedStatus: 404,
			expectedBody:   `{"message":"Not Found"}`,
		},
		{
			name:    "Failed with internal server error - Due to unexpected errors",
			idParam: "1",
			setupMock: func() {
				mockService.EXPECT().DeleteTodo(1).
					Return(errors.New("unexpected error"))
			},
			expectedStatus: 500,
			expectedBody:   `{"message":"Internal Server Error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			handler := NewTodo(mockService)

			mux := http.NewServeMux()
			mux.HandleFunc("DELETE /v1/todos/{id}", handler.Delete)

			req := httptest.NewRequest(http.MethodDelete, "/v1/todos/"+tc.idParam, nil)
			res := httptest.NewRecorder()

			mux.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.JSONEq(t, tc.expectedBody, res.Body.String())
		})
	}
}

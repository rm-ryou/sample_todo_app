package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rm-ryou/sample_todo_app/internal/entity"
	mock_service "github.com/rm-ryou/sample_todo_app/internal/service/todo/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

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

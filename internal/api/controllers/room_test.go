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

func TestGetAllRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockRoomServicer(ctrl)
	controller := NewRoomController(mockService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/rooms/", controller.GetAll)

	testCases := []struct {
		name           string
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success to Get all room",
			setupMock: func() {
				mockService.EXPECT().GetAll().
					Return([]*entities.Room{
						{
							Id:        1,
							Name:      "test room",
							CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
						},
						{
							Id:        2,
							Name:      "sample room",
							CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2025, 5, 1, 10, 0, 0, 0, time.UTC),
						},
					}, nil)
			},
			expectedStatus: 200,
			expectedBody: `{
				"rooms":[
					{
						"id":1,
						"name":"test room",
						"created_at":"2025-01-01T10:00:00Z",
						"updated_at":"2025-01-01T10:00:00Z"
					},
					{
						"id":2,
						"name":"sample room",
						"created_at":"2025-04-01T10:00:00Z",
						"updated_at":"2025-05-01T10:00:00Z"
					}
				]
			}`,
		},
		{
			name: "If there is no record, return empty json",
			setupMock: func() {
				mockService.EXPECT().GetAll().
					Return(nil, nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"rooms":[]}`,
		},
		{
			name: "Failed with internal server error - Due to unexpected errors",
			setupMock: func() {
				mockService.EXPECT().GetAll().
					Return(nil, errors.New("unexpected error"))
			},
			expectedStatus: 500,
			expectedBody:   `{"message":"Internal Server Error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			req := httptest.NewRequest(http.MethodGet, "/v1/rooms/", nil)
			res := httptest.NewRecorder()

			mux.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.JSONEq(t, tc.expectedBody, res.Body.String())
		})
	}
}

func TestCreateRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockRoomServicer(ctrl)
	controller := NewRoomController(mockService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/rooms", controller.Create)

	testCases := []struct {
		name           string
		requestBody    string
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "Success to Create new room",
			requestBody: `{"Name":"test room"}`,
			setupMock: func() {
				mockService.EXPECT().Create("test room").Return(nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"message":"OK"}`,
		},
		{
			name:           "Failed with bad request - Due to number of characters in the name is more than 50",
			requestBody:    fmt.Sprintf(`{"name":"%s"}`, strings.Repeat("a", 51)),
			setupMock:      func() {},
			expectedStatus: 400,
			expectedBody:   `{"message":"Bad Request"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			body := bytes.NewBufferString(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/v1/rooms", body)
			res := httptest.NewRecorder()

			mux.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.JSONEq(t, tc.expectedBody, res.Body.String())
		})
	}
}

func TestUpdateRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockRoomServicer(ctrl)
	controller := NewRoomController(mockService)

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /v1/rooms/{id}", controller.Update)

	testCases := []struct {
		name           string
		idParam        string
		requestBody    string
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "Success to Update room",
			idParam:     "1",
			requestBody: `{"name":"update name"}`,
			setupMock: func() {
				mockService.EXPECT().Update(1, "update name").
					Return(nil)
			},
			expectedStatus: 200,
			expectedBody:   `{"message":"OK"}`,
		},
		{
			name:           "Failed with invalid request - Due to non-numeric id",
			idParam:        "invalid",
			requestBody:    `{"name":"update name"}`,
			setupMock:      func() {},
			expectedStatus: 400,
			expectedBody:   `{"message":"Bad Request"}`,
		},
		{
			name:           "Failed with bad request - Due to number of characters in the title is more than 50",
			idParam:        "1",
			requestBody:    fmt.Sprintf(`{"name":"%s"}`, strings.Repeat("a", 51)),
			setupMock:      func() {},
			expectedStatus: 400,
			expectedBody:   `{"message":"Bad Request"}`,
		},
		{
			name:        "Failed with not found - Due to no room with id",
			idParam:     "999",
			requestBody: `{"name":"update name"}`,
			setupMock: func() {
				mockService.EXPECT().Update(999, "update name").
					Return(sql.ErrNoRows)
			},
			expectedStatus: 404,
			expectedBody:   `{"message":"Not Found"}`,
		},
		{
			name:        "Failed with internal server error - Due to unexpected errors",
			idParam:     "1",
			requestBody: `{"name":"update name"}`,
			setupMock: func() {
				mockService.EXPECT().Update(1, "update name").
					Return(errors.New("unexpected error"))
			},
			expectedStatus: 500,
			expectedBody:   `{"message":"Internal Server Error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			body := bytes.NewBufferString(tc.requestBody)
			req := httptest.NewRequest(http.MethodPut, "/v1/rooms/"+tc.idParam, body)
			res := httptest.NewRecorder()

			mux.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.JSONEq(t, tc.expectedBody, res.Body.String())
		})
	}
}

func TestDeleteRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockRoomServicer(ctrl)
	controller := NewRoomController(mockService)

	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /v1/rooms/{id}", controller.Delete)

	testCases := []struct {
		name           string
		idParam        string
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "Success to Delete room",
			idParam: "1",
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
			setupMock:      func() {},
			expectedStatus: 400,
			expectedBody:   `{"message":"Bad Request"}`,
		},
		{
			name:    "Failed with not found - Due to no room with id",
			idParam: "999",
			setupMock: func() {
				mockService.EXPECT().Delete(999).
					Return(sql.ErrNoRows)
			},
			expectedStatus: 404,
			expectedBody:   `{"message":"Not Found"}`,
		},
		{
			name:    "Failed with internal server error - Due to unexpected errors",
			idParam: "1",
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

			req := httptest.NewRequest(http.MethodDelete, "/v1/rooms/"+tc.idParam, nil)
			res := httptest.NewRecorder()

			mux.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			assert.JSONEq(t, tc.expectedBody, res.Body.String())
		})
	}
}

package handler

import (
	"bytes"
	"cinema_diary"
	"cinema_diary/pkg/service"
	mock_service "cinema_diary/pkg/service/mocks"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user cinema_diary.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           cinema_diary.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"id": 1, "name": "Test", "login":"login", "password":"123"}`,
			inputUser: cinema_diary.User{
				Id:       1,
				Name:     "Test",
				Login:    "login",
				Password: "123",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user cinema_diary.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},

		{
			name:                "Empty fields",
			inputBody:           `{"login":"login", "password":"123"}`,
			mockBehavior:        func(s *mock_service.MockAuthorization, user cinema_diary.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid input body"}`,
		},

		{
			name:      "Service failure",
			inputBody: `{"id": 1, "name": "Test", "login":"login", "password":"123"}`,
			inputUser: cinema_diary.User{
				Id:       1,
				Name:     "Test",
				Login:    "login",
				Password: "123",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user cinema_diary.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("Service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `"Service failure"`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

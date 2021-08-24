package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserViewHandler(t *testing.T) {
	type want struct {
		userParam    string
		user         string
		responseCode int
		contentType  string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "test 1",
			want: want{
				userParam:    "user_id",
				user:         "user1",
				responseCode: 200,
				contentType:  "application/json",
			},
		},
	}

	users := map[string]User{
		"user1": {
			Name:     "Test",
			LastName: "Test",
		},
		"user2": {
			Name:     "Test 2",
			LastName: "Test 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", "/user?user_id=user1", nil)
			w := httptest.NewRecorder()
			h := http.Handler(UserViewHandler(users))
			h.ServeHTTP(w, request)
			res := w.Result()
			// возвращаемый код ответа
			assert.Equal(t, tt.want.responseCode, res.StatusCode)
			//поле user в map
			assert.NotNil(t, users[tt.want.user])
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
			assert.NotEmpty(t, request.FormValue(tt.want.userParam))
		})
	}
}

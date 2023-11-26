package authentication

import (
	"bytes"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserLogin(t *testing.T) {
	t.Run("Successful login", func(t *testing.T) {
		requestBody := []byte(`{"username": "testuser", "password": "password123"}`)
		request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()

		UserLogin(response, request)

		// Check the response status code
		if response.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.Code)
		}
	})
}

func TestUserRegister(t *testing.T) {
	t.Run("Successful login", func(t *testing.T) {
		requestBody := []byte(`{"username": "testuser", "password": "password123"}`)
		request, err := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()

		UserRegister(response, request)

		// Check the response status code
		if response.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.Code)
		}
	})
}

func Test_createToken(t *testing.T) {
	user := User{
		Username: "user",
		Password: "123",
	}
	token := createToken(user)
	TokenAuth = jwtauth.New("HS256", []byte("Secret key"), nil)
	_, err := jwtauth.VerifyToken(TokenAuth, token)
	if err != nil {
		t.Errorf("Function does not work, %v", err)
	}
}

func Test_hashAndStorePassword(t *testing.T) {
	user := User{
		Username: "user",
		Password: "123",
	}
	testMap := make(map[string]string)
	err := hashAndStorePassword(user, testMap)
	if err != nil {
		t.Errorf("Function does not work, %v", err)

	}
	if testMap[user.Username] == "" {
		t.Errorf("Error! got nil password")
	}
}

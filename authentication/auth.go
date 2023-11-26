package authentication

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
)

var TokenAuth *jwtauth.JWTAuth
var users = make(map[string]string)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("Secret key"), nil)
}

// @Summary UserRegister
// @Tags аутентификация
// @Description Регистрирует юзера в базе
// @Accept json
// @Produce json
// @Param input body User true "логин и пароль"
// @Success 200
// @Failure 400 "Неверные данные"
// @Failure 500 "Сервер не работает"
// @Router /api/register/ [post]
func UserRegister(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	err = hashAndStorePassword(user, users)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = w.Write([]byte("Регистрация прошла успешно"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Println("Новый клиент зарегистрировался")
}

// @Summary UserLogin
// @Tags аутентификация
// @Description Выдаёт клиенту токен jwt
// @Accept json
// @Produce json
// @Param input body User true "логин и пароль"
// @Success 200
// @Failure 400 "Неверные данные"
// @Failure 500 "Сервер не работает"
// @Router /api/login/ [post]
func UserLogin(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(500)
	}
	//переводим на регистрацию, если такой пользователь отсутствует
	if _, ok := users[user.Username]; !ok {
		http.Redirect(w, r, "/api/register/", 200)
	}

	getPas, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if users[user.Username] != string(getPas) {
		w.WriteHeader(403)
	}
	token := createToken(user)

	//отсылаем токен
	_, err = w.Write([]byte(token))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Println("Клиент получил токен")
}

func createToken(user User) string {
	_, tokenString, _ := TokenAuth.Encode(map[string]interface{}{user.Username: user.Password})

	log.Println("Создан новый токен")
	return tokenString
}

func hashAndStorePassword(user User, users map[string]string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Сохраняем хешированный пароль в карту
	users[user.Username] = string(hashedPassword)

	return nil
}

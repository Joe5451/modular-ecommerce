package domain

import (
	"net/http"

	"github.com/Joe5451/modular-ecommerce/internal/errorx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

var (
	ErrUserNameCannotBeBlank     = errorx.New(http.StatusBadRequest, "VALIDATION_ERROR", "the user name cannot be blank")
	ErrUserIDCannotBeBlank       = errorx.New(http.StatusBadRequest, "VALIDATION_ERROR", "the user id cannot be blank")
	ErrUserEmailCannotBeBlank    = errorx.New(http.StatusBadRequest, "VALIDATION_ERROR", "the user email cannot be blank")
	ErrUserPasswordCannotBeBlank = errorx.New(http.StatusBadRequest, "VALIDATION_ERROR", "the user password cannot be blank")
	ErrUserNotAuthorized         = errorx.New(http.StatusUnauthorized, "UNAUTHORIZED", "user is not authorized")
)

func RegisterUser(id, name, email, password string) (*User, error) {
	if id == "" {
		return nil, ErrUserIDCannotBeBlank
	}

	if name == "" {
		return nil, ErrUserNameCannotBeBlank
	}

	if email == "" {
		return nil, ErrUserEmailCannotBeBlank
	}

	if password == "" {
		return nil, ErrUserEmailCannotBeBlank
	}

	password, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

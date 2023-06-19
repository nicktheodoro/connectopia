package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// UserModel represents a user model.
type UserModel struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// Prepare will call the methods to validate and format the received user
func (user *UserModel) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	if err := user.format(); err != nil {
		return err
	}

	return nil
}

func (user *UserModel) validate() error {
	if user.Name == "" {
		return errors.New("name is required and cannot be blank")
	}

	if user.Username == "" {
		return errors.New("username is required and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("email is required and cannot be blank")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("the inserted email is invalid")
	}

	if user.Password == "" {
		return errors.New("password is required and cannot be blank")
	}

	return nil
}

func (user *UserModel) format() error {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	return nil
}

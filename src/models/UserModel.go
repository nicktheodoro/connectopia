package models

import (
	"connectopia-api/src/security"
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
func (user *UserModel) Prepare(actionType string) error {
	if err := user.validate(actionType); err != nil {
		return err
	}

	if err := user.format(actionType); err != nil {
		return err
	}

	return nil
}

func (user *UserModel) validate(actionType string) error {
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

	if actionType == "insert" && user.Password == "" {
		return errors.New("password is required and cannot be blank")
	}

	return nil
}

func (user *UserModel) format(actionType string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	if actionType == "insert" {
		hashedPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)
	}

	return nil
}

package users

import (
	"database/sql"
	"errors"

	"github.com/go-panton/mcre/users/models"
)

// Service is the interface of User API
type Service interface {
	SignUp(username, password string) error
}
type service struct {
	repo models.UserRepository
}

// NewService instantiates new user-service.
func NewService(repo models.UserRepository) Service {
	return &service{repo}
}

func (svc *service) SignUp(username, password string) error {
	if username == "" {
		return BadRequestError{errors.New("The username is empty.")}
	} else if password == "" {
		return BadRequestError{errors.New("The password is empty.")}
	}

	searchUser, err := svc.repo.Find(username)
	if err != nil && err != sql.ErrNoRows {
		return BadRequestError{err}
	}
	if searchUser != nil {
		return BadRequestError{errors.New("The username has already been taken.")}
	}

	return svc.repo.Insert(username, password)
}

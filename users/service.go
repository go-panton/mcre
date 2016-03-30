package users

import (
	"errors"
	"github.com/go-panton/mcre/users/model"
)

// Service is the interface of User API
type Service interface {
	SignUp(username, password string) error
}
type service struct{
	repo models.UserRepository
}
// NewService instantiates new user-service.
func NewService(repo models.UserRepository) Service {
	return &service{repo}
}

func (svc *service) SignUp(username, password string) error {
	if username == "" || password == "" {
		return errors.New("Username is : " + username + " Password is: " + password)
	}

	return svc.repo.Insert(username,password)
}

package users

import "errors"

// Service is the interface of User API
type Service interface {
	User(username, password string) (bool, error)
}

// NewService instantiates new user-service.
func NewService() Service {
	return &service{}
}

type service struct{}

func (svc *service) User(username, password string) (bool, error) {
	if username != "" && password != "" {
		return true, nil
	}
	return false, errors.New("Username is : " + username + " Password is: " + password)
}

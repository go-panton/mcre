package users

import "errors"

// Service is the interface of User API
type Service interface {
	User(username, password string) error
}

// NewService instantiates new user-service.
func NewService() Service {
	return &service{}
}

type service struct{}

func (svc *service) User(username, password string) error {
	if username != "" && password != "" {
		return nil
	}
	return errors.New("Username is : " + username + " Password is: " + password)
}

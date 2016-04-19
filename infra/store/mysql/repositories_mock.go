package mysql

import (
	"errors"

	id "github.com/go-panton/mcre/id/model"
	user "github.com/go-panton/mcre/users/model"
)

type mockUserRepository struct {
	UserArray []user.User
}

type mockSeqRepository struct {
}

//NewMockUserRepository returns a mock userRepository
func NewMockUserRepository() user.UserRepository {
	return &mockUserRepository{}
}

//NewMockSeqRepository returns a mock seqRepository
func NewMockSeqRepository() id.SeqRepository {
	return &mockSeqRepository{}
}

func (m *mockUserRepository) Find(username string) (*user.User, error) {
	if username == "alex" {
		return &user.User{Username: "alex", Password: "root"}, nil
	}
	return nil, nil
}

func (m *mockUserRepository) Insert(username, password string) error {
	return nil
}

func (m *mockUserRepository) Verify(username, password string) (*user.User, error) {
	if username == "alex" && password == "root" {
		return &user.User{Username: "alex", Password: "root"}, nil
	}
	return nil, errors.New("Invalid User")
}

func (m *mockSeqRepository) Find(query string) (int, error) {
	if query == "" {
		return 0, errors.New("The query string is empty.")
	} else if query == "NODE" {
		return 200899, nil
	} else if query == "FILENAME" {
		return 761, nil
	}
	return 0, nil
}

func (m *mockSeqRepository) Update(query string, value int) error {
	if query == "" {
		return errors.New("The query string is empty")
	}
	if value < 1 {
		return errors.New("The value should not be less than 1")
	}
	return nil
}

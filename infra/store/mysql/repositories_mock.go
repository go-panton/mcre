package mysql

import (
	"errors"

	"github.com/go-panton/mcre/id/model"
	"github.com/go-panton/mcre/users/model"
)

type mockUserRepository struct {
	UserArray []models.User
}

type mockSeqRepository struct {
}

func NewMockUserRepository() models.UserRepository {
	return &mockUserRepository{}
}

func NewMockSeqRepository() model.SeqRepository {
	return &mockSeqRepository{}
}

func (m *mockUserRepository) Find(username string) (*models.User, error) {
	if username == "alex" {
		return &models.User{Username: "alex", Password: "root"}, nil
	}
	return nil, nil
}

func (m *mockUserRepository) Insert(username, password string) error {
	return nil
}

func (m *mockUserRepository) Verify(username, password string) (*models.User, error) {
	if username == "alex" && password == "root" {
		return &models.User{Username: "alex", Password: "root"}, nil
	}
	return nil, errors.New("Invalid User")
}

func (m *mockSeqRepository) Get(query string) (int, error) {
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

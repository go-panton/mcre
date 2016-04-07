package mysql

import (
	"errors"

	"github.com/go-panton/mcre/users/model"
)

type mockUserRepository struct {
	UserArray []models.User
}

type mockSeqRepository struct {
	SeqStruct models.Seq
}

func NewMockUserRepository() models.UserRepository {
	return &mockUserRepository{}
}

func NewMockSeqRepository() models.SeqRepository {
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
	}
	return 0, nil
}

func (m *mockSeqRepository) Update(val int, query string) error {
	if query == "" {
		return errors.New("The query string is empty.")
	}
	return nil
}

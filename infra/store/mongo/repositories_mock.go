package mongo

import "github.com/go-panton/mcre/users/model"

type mockUserRepository struct {
	UserArray []model.User
}

func NewMockUserRepository() model.UserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) Find(username string) (*model.User, error) {
	if username == "alex" {
		return &model.User{Username: "alex", Password: "root"}, nil
	}
	return nil, nil
}
func (m *mockUserRepository) Insert(username, password string) error {
	return nil
}

func (m *mockUserRepository) Verify(username, password string) (*model.User, error) {
	return nil, nil
}

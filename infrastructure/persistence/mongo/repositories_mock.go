package mongo

import "github.com/go-panton/mcre/users/model"

type mockUserRepository struct{
	UserArray []models.User
}

func NewMockUserRepository() models.UserRepository{
	return &mockUserRepository{}
}

func (m *mockUserRepository) Find(username string) (*models.User, error){
	if username == "alex" {
		return &models.User{Username:"alex",Password:"root"}, nil
	}
	return nil,nil
}
func (m *mockUserRepository) Insert(username,password string) (error){
	return nil
}
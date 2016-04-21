package models

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

// Repository provides method to query database.
type UserRepository interface {
	Find(username string) (*User, error)
	Insert(username, password string) error
}

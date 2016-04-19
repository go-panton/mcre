package model

//User struct define the user
type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

// UserRepository provided method to query the database
type UserRepository interface {
	Find(username string) (*User, error)
	Insert(username, password string) error
	Verify(username, password string) (*User, error)
}

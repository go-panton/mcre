package models

//User struct define the data that stores inside database
type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

//UserRepository provides method to query database.
type UserRepository interface {
	Find(username string) (*User, error)
	Insert(username, password string) error
}

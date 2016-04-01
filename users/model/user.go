package models

type User struct{
	Username    string  `bson:"username"`
	Password    string  `bson:"password"`
}

// Repository provides access a cargo store.
type UserRepository interface {
	Find(userId string) (*User, error)
	Insert(username,password string) (error)
}
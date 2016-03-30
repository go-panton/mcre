package mongo

import(
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/go-panton/mcre/users/model"
)

type userRepository struct {
	col *mgo.Collection
}

func ConnectDatabase(dbName,colName string) *mgo.Collection {
	return session.DB(dbName).C(colName)
}

func NewUser(col *mgo.Collection) models.UserRepository {
	return &userRepository{col}
}

func (r *userRepository)Insert(username,password string) error {
	newUser := models.User{Username:username,Password:password}

	err := r.col.Insert(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Find(userId string) (*models.User, error) {
	result := &models.User{}
	err := r.col.Find(bson.M{"username":"alex"}).One(result)
	if err != nil {
		return nil,err
	}

	return result, nil
}
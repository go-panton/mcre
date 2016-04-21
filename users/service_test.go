package users

import (
	"testing"

	"errors"

	"github.com/go-panton/mcre/infra/store/mongo"
	"github.com/go-panton/mcre/users/models"
)

type testPair struct {
	TestData models.User
	Expected BadRequestError
}

var tests = []testPair{
	{models.User{Username: "alex", Password: ""}, BadRequestError{errors.New("The password is empty.")}},
	{models.User{Username: "", Password: "root"}, BadRequestError{errors.New("The username is empty.")}},
	{models.User{Username: "", Password: ""}, BadRequestError{errors.New("The username is empty.")}},
	{models.User{Username: "alex", Password: "root"}, BadRequestError{errors.New("The username has already been taken.")}},
	{models.User{Username: "Rex", Password: "Gear"}, BadRequestError{nil}},
}

func TestSignUp(t *testing.T) {
	for _, test := range tests {
		err := NewService(mongo.NewMockUserRepository()).SignUp(test.TestData.Username, test.TestData.Password)
		if err != nil {
			if test.Expected.Error() != err.Error() {
				t.Errorf("Want: \n%#v \nGot: \n%#v", test.Expected, err)
			}
		}
	}
}

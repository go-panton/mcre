package users

import (
	"testing"
	"fmt"
	"github.com/go-panton/mcre/users/model"
	"github.com/go-panton/mcre/infrastructure/persistence/mongo"
)

var tests = []models.User{
	{"alex", ""},//password empty
	{"","root"},//username empty
	{"",""},//both field empty
	{"alex","root"},//username already exist in database
	{"Rex","Gear"},//success case
}

//func TestService(t *testing.T){
//	var repo models.UserRepository
//	for _, pair := range tests {
//		err := NewService(repo).SignUp(pair.Username,pair.Password)
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//		return
//	}
//}

func TestSignUp(t *testing.T){
	for _,pair := range tests{
		err := NewService(mongo.NewMockUserRepository()).SignUp(pair.Username,pair.Password)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

//func TestCaseMongo(t *testing.T) {
//	session, err := mgo.Dial("localhost")
//	defer session.Close()
//	if err != nil {
//		fmt.Println(err)
//	}
//	handle := MakeHandler(context.Background(),NewService(mongo.NewUser(session.DB("go_panton").C("user"))))
//	w := httptest.NewRecorder()
//
//	for _, pair := range tests{
//		bytePair, err1 := json.Marshal(pair)
//		if err1 != nil {
//			fmt.Println(err1)
//		}
//		r, err2 := http.NewRequest("POST", "/users", bytes.NewReader(bytePair))
//		r.Header.Set("Content-Type","application/json")
//		if err2 != nil {
//			log.Fatal(err2)
//		}
//		handle.ServeHTTP(w,r)
//
//		fmt.Println(w.Code)
//		fmt.Println(w.Body.String())
//	}
//}
package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func TestUserRepo(t *testing.T) {
	db, err := sql.Open("mysql", "root:root123@/go_panton")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	NewUser(db).Insert("test", "random")

	result, err := NewUser(db).Find("test")
	if err != nil || result == nil {
		t.Errorf("No result from database")
	}
	fmt.Println(result)
}

func TestSeqFind(t *testing.T) {
	tests := []struct {
		Key   string
		Want  int
		Error error
	}{
		{"NODE", 200905, nil},
		{"FILENAME", 766, nil},
		{"", 0, errors.New("The key is empty.")},
	}

	db, err := sql.Open("mysql", "root:root123@/ptd_new")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	for _, test := range tests {
		res, err := NewSeq(db).Find(test.Key)
		if err != nil {
			if err.Error() != test.Error.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Error.Error())
			}
		}
		if res != test.Want {
			t.Errorf("Got: %v,Want: %v", res, test.Want)
		}
	}
}

func TestSeqUpdate(t *testing.T) {
	db, err := sql.Open("mysql", "root:root123@/ptd_new")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	tests := []struct {
		Key   string
		Value int
		Error error
	}{
		{"NODE", 200905, nil},
		{"FILENAME", 766, nil},
		{"NODE", 0, errors.New("Update Value cannot be less than 1")},
		{"FILENAME", 0, errors.New("Update Value cannot be less than 1")},
		{"", 0, errors.New("The key is empty.")},
	}

	for _, test := range tests {
		err := NewSeq(db).Update(test.Key, test.Value)

		if err != nil {
			if err.Error() != test.Error.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Error.Error())
			}
		}
	}
}

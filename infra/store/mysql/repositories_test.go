package mysql

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMySQL(t *testing.T) {
	db, err := sql.Open("mysql", "root:root123@/go_panton")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	NewUser(db).Insert("test", "random")

	result, err := NewUser(db).Find("test")
	if err != nil || result == nil {
		t.Errorf("No result from database %v", err)
	}
	fmt.Println(result)
}

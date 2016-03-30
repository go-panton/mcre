package mysql

import (
	"testing"
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

)

func TestMySQL(t *testing.T) {
	db, err := sql.Open("mysql", "root:root123@/go_panton")
	if err != nil {
		fmt.Println("Error: %v",err)
	}

	NewUser(db).Insert("test","random")
}

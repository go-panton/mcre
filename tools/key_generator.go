package tools

import (
	"github.com/go-panton/mcre/infra/store/mysql"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// IDGenerator return the value based on query input
//
// Return error when :-
// - query has no result
func IDGenerator(query string) (int, error) {
	seq := mysql.NewSeq(mysql.ConnectDatabase("root:root123@/ptd_new"))

	//seq := mysql.NewMockSeqRepository()

	//query from database
	val, err := seq.Find(query)
	if err != nil {
		return 0, err
	}

	nextNumber := val + 1

	//update the value
	//TODO: make it thread safe
	err1 := seq.Update(nextNumber, query)
	if err1 != nil {
		return 0, err1
	}
	return nextNumber, nil
}

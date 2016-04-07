package tools

import (
	"github.com/go-panton/mcre/infra/store/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func IDGenerator(query string) (int, error) {
	seq := mysql.NewSeq(mysql.ConnectDatabase("root:root123@/ptd_new"))

	//seq := mysql.NewMockSeqRepository()

	//query from database
	val, err := seq.Get(query)
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

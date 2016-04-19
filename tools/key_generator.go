package tools

import (
	"github.com/go-panton/mcre/id/model"
	_ "github.com/go-sql-driver/mysql"
)

//IDGenerator get the value from table based on query, increment it by 1(indicating a new number)
//and return the new ID
func IDGenerator(seq model.SeqRepository, query string) (int, error) {
	//query from database
	val, err := seq.Get(query)
	if err != nil {
		return 0, err
	}

	return val + 1, nil
}

//UpdateValue update the ID back to table to keep track of the ID
func UpdateValue(seq model.SeqRepository, query string, value int) error {
	//update the value
	//TODO: make it thread safe
	err := seq.Update(query, value)
	if err != nil {
		return err
	}
	return nil
}

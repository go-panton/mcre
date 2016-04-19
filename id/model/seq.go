package model

//SeqRepository is the interface to use method
type SeqRepository interface {
	// Find return value based on the key provided
	//
	// return error when :
	// - key is empty
	// - there is no result return
	Find(Query string) (int, error)
	// Update update the value corresponding to the key in database
	//
	// return error when :
	// - key is empty.
	// - value is less than 1
	// - update has failed
	Update(Query string, Value int) error
}

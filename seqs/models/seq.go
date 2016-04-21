package models

//SeqRepository define the interface method to be called
type SeqRepository interface {
	// Find return value based on the key provided
	//
	// return error when :
	// - key is empty
	// - there is no result return
	Find(string) (int, error)
	// Update update the value corresponding to the key in database
	//
	// return error when :
	// - key is empty.
	// - value is less than 1
	// - update has failed
	Update(string, int) error
}

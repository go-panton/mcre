package models

type SeqRepository interface {
	Get(Query string) (int, error)
	Update(Query string, Value int) error
}

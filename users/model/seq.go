package models

type Seq struct {
	Name  string
	Value int
}

type SeqRepository interface {
	Get(string) (int, error)
	Update(int, string) error
}

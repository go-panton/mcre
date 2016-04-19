package id

import (
	"testing"

	"errors"

	"github.com/go-panton/mcre/infra/store/mysql"
)

func TestFind(t *testing.T) {
	var tests = []struct {
		key  string
		want int
	}{
		{"NODE", 200899},
		{"", 0},
		{"FILENAME", 761},
	}

	seq := mysql.NewMockSeqRepository()

	for _, test := range tests {
		value, _ := seq.Find(test.key)
		if value != test.want {
			t.Errorf("Want: %v, Got: %v", test.want, value)
		}
	}
}

func TestUpdate(t *testing.T) {
	var tests = []struct {
		key   string
		value int
		want  error
	}{
		{"NODE", 80877, nil},
		{"NODE", 0, errors.New("The value should not be less than 1")},
		{"FILENAME", 34343, nil},
		{"FILENAME", 0, errors.New("The value should not be less than 1")},
		{"", 0, errors.New("The key is empty")},
		{"", 34353, errors.New("The key is empty")},
	}

	seq := mysql.NewMockSeqRepository()

	for _, test := range tests {
		err := seq.Update(test.key, test.value)
		if err != nil {
			if err.Error() != test.want.Error() {
				t.Errorf("Want: %v, Got: %v", test.want, err)
			}
		}
	}
}

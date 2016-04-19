package tools

import (
	"testing"

	"github.com/go-panton/mcre/infra/store/mysql"
)

type TestCase struct {
	TestData string
	Expected int
}

func TestKeygen(t *testing.T) {
	tests := []TestCase{
		{"NODE", 200900},
		{"", 0},
		{"FILENAME", 762},
	}

	for _, test := range tests {
		seq := mysql.NewMockSeqRepository()
		value, _ := IDGenerator(seq, test.TestData)
		if value != test.Expected {
			t.Errorf("Want: %v, Got: %v", test.Expected, value)
		}
	}
}

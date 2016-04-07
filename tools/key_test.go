package tools

import "testing"

type TestCase struct {
	TestData string
	Expected int
}

func TestKeygen(t *testing.T) {
	tests := []TestCase{
		{"NODE", 200900},
		{"", 0},
	}

	for _, test := range tests {
		value, _ := IDGenerator(test.TestData)
		if value != test.Expected {
			t.Errorf("Want: %v, Got: %v", test.Expected, value)
		}
	}
}

package tools

import "testing"

func TestMockFile(t *testing.T) {
	err := mockInsertFile()
	if err != nil {
		t.Errorf(err.Error())
	}
}

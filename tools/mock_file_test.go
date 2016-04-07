package tools

import "testing"

func TestMockFile(t *testing.T) {
	err := mock_insert_file()
	if err != nil {
		t.Errorf(err.Error())
	}
}

package files

import (
	"fmt"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	newFile, err := os.Create("D:\\PantonSys\\PTD\\ISOStorage\\1\\test.txt")
	if err != nil {
		fmt.Errorf("Unable to create file")
	}
	newFile.Close()
}

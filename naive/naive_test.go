package naive

import (
	"fmt"
	"strings"
	"testing"
)

func setup() {

}
func teardown() {

}

func TestNaive(t *testing.T) {
	mcre := New()
	_, err := mcre.PutObject("gen", "b/b/a.txt", strings.NewReader("abc"))
	fmt.Println(err)
}

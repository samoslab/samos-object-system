package ticket

import (
	"testing"
	"fmt"
)

func TestGenerate(t *testing.T) {
	r := Generate()
	fmt.Println(r)
}

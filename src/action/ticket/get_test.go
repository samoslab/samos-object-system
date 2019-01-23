package ticket

import (
	"testing"

	"github.com/samoslab/samos-object-system/src/model/ticket"
)

func TestG(t *testing.T) {
	r := ticket.Generate()
	t.Log(r)
}

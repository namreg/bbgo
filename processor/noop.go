package processor

import (
	"io"

	"github.com/namreg/bbgo/node"
)

// Noop processor does nothing.
func Noop(t node.Tag, w io.Writer) {
	io.WriteString(w, t.String())
}

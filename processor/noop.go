package processor

import (
	"io"

	"github.com/namreg/bbgo/context"
	"github.com/namreg/bbgo/node"
)

// Noop processor does nothing.
func Noop(ctx *context.Context, t node.Tag, w io.Writer) {
	io.WriteString(w, t.String())
}

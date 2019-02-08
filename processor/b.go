package processor

import (
	"io"

	"github.com/namreg/bbgo/context"
	"github.com/namreg/bbgo/node"
)

// B processes [b] bbcode.
func B(ctx *context.Context, tag node.Tag, w io.Writer) {
	switch tag.(type) {
	case *node.OpeningTag:
		io.WriteString(w, "<strong>")
	case *node.ClosingTag:
		io.WriteString(w, "</strong>")
	}
}

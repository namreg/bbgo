package processor

import (
	"io"

	"github.com/namreg/bbgo/context"
	"github.com/namreg/bbgo/node"
)

// Simple processes simple bbcodes such as [i], [b], [u], [s].
func Simple(ctx *context.Context, tag node.Tag, w io.Writer) {
	switch t := tag.(type) {
	case *node.OpeningTag:
		io.WriteString(w, "<")
		io.WriteString(w, t.TagName())
		io.WriteString(w, ">")
	case *node.ClosingTag:
		io.WriteString(w, "</")
		io.WriteString(w, t.TagName())
		io.WriteString(w, ">")
	}
}

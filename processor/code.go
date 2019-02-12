package processor

import (
	"io"

	"github.com/namreg/bbgo/context"
	"github.com/namreg/bbgo/node"
)

// Code processes [code] bbcode.
func Code(ctx *context.Context, tag node.Tag, w io.Writer) {
	switch tag.(type) {
	case *node.OpeningTag:
		io.WriteString(w, "<pre>")
		ctx.BeginRawMode(tag)
	case *node.ClosingTag:
		io.WriteString(w, "</pre>")
		ctx.EndRawMode(tag)
	}
}

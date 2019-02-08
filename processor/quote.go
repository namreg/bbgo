package processor

import (
	"io"

	"github.com/namreg/bbgo/context"
	"github.com/namreg/bbgo/node"
)

// Quote processes [quote] bbcode.
func Quote(ctx *context.Context, tag node.Tag, w io.Writer) {
	switch t := tag.(type) {
	case *node.OpeningTag:
		io.WriteString(w, `<blockquote>`)
		if n, ok := t.Attrs()["name"]; ok {
			io.WriteString(w, `<cite>`)
			io.WriteString(w, n)
			io.WriteString(w, ` said:</cite>`)
		}
	case *node.ClosingTag:
		io.WriteString(w, `</blockquote>`)
	}
}

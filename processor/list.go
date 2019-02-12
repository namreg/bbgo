package processor

import (
	"io"

	"github.com/namreg/bbgo/context"
	"github.com/namreg/bbgo/node"
)

// List processes [list] bbcode.
func List(ctx *context.Context, tag node.Tag, w io.Writer) {
	switch tag.(type) {
	case *node.OpeningTag:
		io.WriteString(w, "<ul>")
	case *node.ClosingTag:
		// closing list item
		io.WriteString(w, "</li>")
		io.WriteString(w, "</ul>")
	}
}

// Asterisk processes [*] list item bbcode.
func Asterisk(ctx *context.Context, tag node.Tag, w io.Writer) {
	switch tag.(type) {
	case *node.OpeningTag:
		prev2 := ctx.Prev2Node()
		if prev2 != nil && prev2.Token().Literal == "*" {
			// closing previous list item
			io.WriteString(w, "</li>")
		}
		io.WriteString(w, "<li>")
	}
}

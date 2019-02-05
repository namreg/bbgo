package processor

import (
	"io"

	"github.com/namreg/bbgo/node"
)

// B processor process [b] bbcode.
func B(tag node.Tag, w io.Writer) {
	switch tag.(type) {
	case *node.OpeningTag:
		io.WriteString(w, "<strong>")
	case *node.ClosingTag:
		io.WriteString(w, "</strong>")
	}
}

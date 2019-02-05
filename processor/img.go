package processor

import (
	"io"

	"github.com/namreg/bbgo/node"
)

// Img processes [img] bbcode.
func Img(tag node.Tag, w io.Writer) {
	switch t := tag.(type) {
	case *node.OpeningTag:
		io.WriteString(w, `<img `)
		if a := t.Attr(); a != "" {
			io.WriteString(w, `title="`)
			io.WriteString(w, a)
			io.WriteString(w, `" `)
		}
		io.WriteString(w, `src="`)
	case *node.ClosingTag:
		io.WriteString(w, `" />`)
	}
}

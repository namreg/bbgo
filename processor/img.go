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
		if v := t.Value(); v != "" {
			io.WriteString(w, `title="`)
			io.WriteString(w, v)
			io.WriteString(w, `" `)
		}
		io.WriteString(w, `src="`)
	case *node.ClosingTag:
		io.WriteString(w, `" />`)
	}
}

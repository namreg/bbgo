package processor

import (
	"io"
	"strings"

	"github.com/namreg/bbgo/context"
	"github.com/namreg/bbgo/node"
)

// Color processes [color] bbcode.
func Color(ctx *context.Context, tag node.Tag, w io.Writer) {
	switch t := tag.(type) {
	case *node.OpeningTag:
		sanitize := func(r rune) rune {
			if r == '#' || r == ',' || r == '.' || r == '(' || r == ')' || r == '%' {
				return r
			} else if r >= '0' && r <= '9' {
				return r
			} else if r >= 'a' && r <= 'z' {
				return r
			} else if r >= 'A' && r <= 'Z' {
				return r
			}
			return -1
		}
		color := strings.Map(sanitize, t.Value())
		io.WriteString(w, `<span`)

		if color != "" {
			io.WriteString(w, ` style="color: `)
			io.WriteString(w, color)
			io.WriteString(w, `;"`)
		}
		io.WriteString(w, `>`)

	case *node.ClosingTag:
		io.WriteString(w, `</span>`)
	}
}

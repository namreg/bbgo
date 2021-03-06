package node

import (
	"strings"

	"github.com/namreg/bbgo/token"
)

// Text is a text node.
type Text struct {
	tok token.Token // the first token that comprises text node
	sb  strings.Builder
}

// NewText creates a new text node.
func NewText(tok token.Token, val string) *Text {
	t := &Text{tok: tok}
	t.sb.WriteString(val)
	return t
}

// Token satisfy the Node interface.
func (t *Text) Token() token.Token {
	return t.tok
}

// String satisfies to the Node interface.
func (t *Text) String() string {
	return t.sb.String()
}

// Append appends a string to the value.
func (t *Text) Append(str string) {
	t.sb.WriteString(str)
}

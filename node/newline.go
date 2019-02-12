package node

import (
	"github.com/namreg/bbgo/token"
)

// Newline is a new line node.
type Newline struct {
	tok token.Token
}

// NewLine creates a new line node.
func NewLine(tok token.Token) *Newline {
	return &Newline{tok: tok}
}

// Token satisfy the Node interface.
func (n *Newline) Token() token.Token {
	return n.tok
}

// String satisfies to the Node interface.
func (n *Newline) String() string {
	return "\\n"
}

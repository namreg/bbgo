package node

import (
	"sync"

	"github.com/namreg/bbgo/token"
)

var newLinePool = sync.Pool{
	New: func() interface{} {
		return &Newline{}
	},
}

// Newline is a new line node.
type Newline struct {
	tok token.Token
}

// NewLine creates a new line node.
func NewLine(tok token.Token) *Newline {
	nl := newLinePool.Get().(*Newline)
	nl.tok = tok
	return nl
}

// Token satisfy the Node interface.
func (n *Newline) Token() token.Token {
	return n.tok
}

// String satisfies to the Node interface.
func (n *Newline) String() string {
	return "\\n"
}

package node

import (
	"fmt"

	"github.com/namreg/bbgo/token"
)

var _ Tag = (*ClosingTag)(nil)

// ClosingTag is a bbcode closing tag, i.e [/b].
type ClosingTag struct {
	tok token.Token
}

// Token satisfies to the Node interface.
func (ct *ClosingTag) Token() token.Token {
	return ct.tok
}

// String satisfies to the Node interface.
func (ct *ClosingTag) String() string {
	return fmt.Sprintf("[/%s]", ct.TagName())
}

// TagName satifies to the Node interface.
func (ct *ClosingTag) TagName() string {
	return ct.tok.Literal
}

// NewClosingTag creates a new closing tag.
func NewClosingTag(tok token.Token) *ClosingTag {
	return &ClosingTag{tok: tok}
}

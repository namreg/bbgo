package node

import (
	"fmt"

	"github.com/namreg/bbgo/token"
)

var _ Tag = (*OpeningTag)(nil)

// OpeningTag is a bbcode opening tag, i.e. [b]
type OpeningTag struct {
	tok  token.Token
	attr string
}

// NewOpeningTag creates a new opening tag.
func NewOpeningTag(tok token.Token, attr string) *OpeningTag {
	return &OpeningTag{tok: tok, attr: attr}
}

// Token satisfies to the Node interface.
func (ot *OpeningTag) Token() token.Token {
	return ot.tok
}

// String satisfies to the Node interface.
func (ot *OpeningTag) String() string {
	if ot.attr != "" {
		return fmt.Sprintf(`[%s="%s"]`, ot.TagName(), ot.Attr())
	}
	return fmt.Sprintf("[%s]", ot.TagName())
}

// TagName satifies to the Node interface.
func (ot *OpeningTag) TagName() string {
	return ot.tok.Literal
}

// Attr returns a bbcode tag attribute.
func (ot *OpeningTag) Attr() string {
	return ot.attr
}

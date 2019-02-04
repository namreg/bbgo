package node

import (
	"fmt"

	"github.com/namreg/bbgo/token"
)

var _ Tag = (*SelfClosingTag)(nil)

// SelfClosingTag is self-closing bbcode tag, i.e. [url="https://google.com" /]
type SelfClosingTag struct {
	tok  token.Token
	attr string
}

// NewSelfClosingTag creates a new opening tag.
func NewSelfClosingTag(tok token.Token, attr string) *SelfClosingTag {
	return &SelfClosingTag{tok: tok, attr: attr}
}

// Token satisfies to the Node interface.
func (ot *SelfClosingTag) Token() token.Token {
	return ot.tok
}

// String satisfies to the Node interface.
func (ot *SelfClosingTag) String() string {
	if ot.attr != "" {
		return fmt.Sprintf(`[%s="%s" /]`, ot.TagName(), ot.Attr())
	}
	return fmt.Sprintf("[%s /]", ot.TagName())
}

// TagName satifies to the Node interface.
func (ot *SelfClosingTag) TagName() string {
	return ot.tok.Literal
}

// Attr returns a bbcode tag attribute.
func (ot *SelfClosingTag) Attr() string {
	return ot.attr
}

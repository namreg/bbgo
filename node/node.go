package node

import (
	"fmt"
	"strings"

	"github.com/namreg/bbgo/token"
)

// Node is a parsing result.
type Node interface {
	// Token returns a main token of the node.
	Token() token.Token
	// String returns a string represention of the node.
	String() string
}

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
	return fmt.Sprintf("[%s]", ot.tok.Literal)
}

// Attr returns a bbcode tag attribute.
func (ot *OpeningTag) Attr() string {
	return ot.attr
}

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
	return fmt.Sprintf("[/%s]", ct.tok.Literal)
}

// NewClosingTag creates a new closing tag.
func NewClosingTag(tok token.Token) *ClosingTag {
	return &ClosingTag{tok: tok}
}

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

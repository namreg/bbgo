package node

import (
	"github.com/namreg/bbgo/token"
)

// Node is the smallest piece of the parsing result.
type Node interface {
	// Token returns a main token of the node.
	Token() token.Token
	// String returns a string represention of the node.
	String() string
}

// Tag is a bbcode tag. It can be either opening, closing or self-closing.
type Tag interface {
	Node
	// TagName returns a tag name.
	TagName() string
}

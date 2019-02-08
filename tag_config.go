package bbgo

import (
	"io"

	"github.com/namreg/bbgo/context"
	"github.com/namreg/bbgo/node"
)

// TagConfig is a tag config.
type TagConfig struct {
	name        string
	escape      bool
	selfclosing bool
	processor   func(*context.Context, node.Tag, io.Writer)
}

// TagOpt is tag option.
type TagOpt func(*TagConfig)

// Escape determines whether the content inside a tag should be escaped.
func Escape(b bool) TagOpt {
	return func(tg *TagConfig) {
		tg.escape = b
	}
}

// SelfClsosing determines whether a tag can be self-closing.
func SelfClsosing(b bool) TagOpt {
	return func(tg *TagConfig) {
		tg.selfclosing = b
	}
}

// Processor sets a tag processor.
func Processor(p func(*context.Context, node.Tag, io.Writer)) TagOpt {
	return func(tg *TagConfig) {
		tg.processor = p
	}
}

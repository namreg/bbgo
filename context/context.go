package context

import (
	basecontext "context"

	"github.com/namreg/bbgo/node"
)

// Context is a wrapper around context.Context with convinient methods.
type Context struct {
	basecontext.Context

	rawModeTag node.Tag

	prevNode  node.Node
	prev2Node node.Node
}

// New creates a new context.
func New(ctx basecontext.Context) *Context {
	return &Context{Context: ctx}
}

// BeginRawMode begins a raw mode for tag.
func (ctx *Context) BeginRawMode(t node.Tag) {
	ctx.rawModeTag = t
}

// EndRawMode ends a raw mode for tag.
func (ctx *Context) EndRawMode(t node.Tag) {
	ctx.rawModeTag = nil
}

// RawModeTag returns a tag in raw mode.
func (ctx *Context) RawModeTag() node.Tag {
	return ctx.rawModeTag
}

// InRawMode checks whether the given tag is in raw mode.
func (ctx *Context) InRawMode(t node.Tag) bool {
	return ctx.rawModeTag != nil && ctx.rawModeTag.TagName() == t.TagName()
}

// SetPrevNode sets a previous processed node.
func (ctx *Context) SetPrevNode(n node.Node) {
	ctx.prev2Node = ctx.prevNode
	ctx.prevNode = n
}

// PrevNode returns a previous processed node.
func (ctx *Context) PrevNode() node.Node {
	return ctx.prevNode
}

// Prev2Node returns a node before previous.
func (ctx *Context) Prev2Node() node.Node {
	return ctx.prev2Node
}

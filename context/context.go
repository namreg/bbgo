package context

import (
	basecontext "context"

	"github.com/namreg/bbgo/node"
)

// Context is a wrapper around context.Context with convinient methods.
type Context struct {
	basecontext.Context

	prevNode  node.Node
	prev2Node node.Node
}

// New creates a new context.
func New(ctx basecontext.Context) *Context {
	return &Context{Context: ctx}
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

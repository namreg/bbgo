package bbgo

import (
	"html"
	"io"
	"strings"

	basecontext "context"

	"github.com/namreg/bbgo/context"
	"github.com/namreg/bbgo/lexer"
	"github.com/namreg/bbgo/node"
	"github.com/namreg/bbgo/parser"
	"github.com/namreg/bbgo/processor"
	"github.com/namreg/bbgo/token"
)

// Processor process a bbcode tag and writes result to the given Writer.
type Processor func(*context.Context, node.Tag, io.Writer)

// BBGO is a main object that contains tag configs, parsing context and etc.
type BBGO struct {
	tags map[string]Processor
}

// New creates a new BBGO and registers default processors.
func New() *BBGO {
	b := &BBGO{
		tags: make(map[string]Processor),
	}
	b.registerDefaultProcessors()
	return b
}

// RegisterTag registers a new tag.
func (b *BBGO) RegisterTag(name string, p Processor) {
	b.tags[name] = p
	token.RegisterIdentifiers(name)
}

// Parse parses the given input.
func (b *BBGO) Parse(input string) string {
	ctx := context.New(basecontext.Background())
	sb := new(strings.Builder)

	l := lexer.New(input)
	p := parser.New(l)

	for _, n := range p.Parse() {
		if t, ok := n.(node.Tag); ok && (ctx.RawModeTag() == nil || ctx.InRawMode(t)) {
			if proc, ok := b.tags[t.TagName()]; ok {
				proc(ctx, t, sb)
			}
		} else if _, ok := n.(*node.Newline); ok {
			io.WriteString(sb, "<br>")
		} else {
			io.WriteString(sb, html.EscapeString(n.String()))
		}
		ctx.SetPrevNode(n)
	}

	return sb.String()
}

func (b *BBGO) registerDefaultProcessors() {
	b.RegisterTag("code", Processor(processor.Code))
	b.RegisterTag("img", Processor(processor.Img))
	b.RegisterTag("quote", Processor(processor.Quote))
	b.RegisterTag("url", Processor(processor.URL))
	b.RegisterTag("list", Processor(processor.List))
	b.RegisterTag("*", Processor(processor.Asterisk))

	for _, t := range []string{"i", "b", "u", "s"} {
		b.RegisterTag(t, Processor(processor.Simple))
	}
}

package bbgo

import (
	basecontext "context"
	"io"
	"strings"

	"github.com/namreg/bbgo/context"
	"github.com/namreg/bbgo/lexer"
	"github.com/namreg/bbgo/node"
	"github.com/namreg/bbgo/parser"
	"github.com/namreg/bbgo/processor"
	"github.com/namreg/bbgo/token"
)

// BBGO is a main object that contains tag configs, parsing context and etc.
type BBGO struct {
	tags map[string]TagConfig
}

// New creates a new BBGO and registers default processors.
func New( /*todo: strict mode and etc.*/ ) *BBGO {
	b := &BBGO{
		tags: make(map[string]TagConfig),
	}
	b.registerDefaultProcessors()
	return b
}

// RegisterTag registers a new tag.
func (b *BBGO) RegisterTag(name string, opts ...TagOpt) {
	tc := TagConfig{name: name}

	for _, o := range opts {
		o(&tc)
	}

	if tc.processor == nil {
		tc.processor = processor.Noop
	}

	b.tags[name] = tc

	token.RegisterIdentifiers(name)
}

// Parse parses the given input.
func (b *BBGO) Parse(input string) string {
	ctx := context.New(basecontext.Background())
	sb := new(strings.Builder)

	l := lexer.New(input)
	p := parser.New(l)

	for _, n := range p.Parse() {
		if t, ok := n.(node.Tag); ok {
			if tc, ok := b.tags[t.TagName()]; ok {
				tc.processor(ctx, t, sb)
			}
		} else {
			// todo(namreg): handle escapes
			io.WriteString(sb, n.String())
		}
		ctx.SetPrevNode(n)
	}

	return sb.String()
}

func (b *BBGO) registerDefaultProcessors() {
	b.RegisterTag("b", Processor(processor.B))
	b.RegisterTag("img", Processor(processor.Img))
	b.RegisterTag("quote", Processor(processor.Quote))
	b.RegisterTag("url", Processor(processor.URL))
}

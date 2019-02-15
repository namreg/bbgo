package node

import (
	"sort"
	"strings"
	"sync"

	"github.com/namreg/bbgo/token"
)

var openingTagPool = sync.Pool{
	New: func() interface{} {
		return &OpeningTag{}
	},
}

var _ Tag = (*OpeningTag)(nil)

// OpeningTag is a bbcode opening tag, i.e. [b]
type OpeningTag struct {
	tok   token.Token
	value string
	attrs map[string]string
}

// NewOpeningTag creates a new opening tag.
func NewOpeningTag(tok token.Token, value string, attrs map[string]string) *OpeningTag {
	ot := openingTagPool.Get().(*OpeningTag)
	ot.tok = tok
	ot.value = value
	ot.attrs = attrs
	return ot
}

// Token satisfies to the Node interface.
func (ot *OpeningTag) Token() token.Token {
	return ot.tok
}

// String satisfies to the Node interface.
func (ot *OpeningTag) String() string {
	var sb strings.Builder
	sb.WriteByte('[')
	sb.WriteString(ot.TagName())

	if ot.value != "" {
		sb.WriteByte('=')
		sb.WriteString(ot.value)
	}

	akeys := make([]string, len(ot.attrs))
	for k := range ot.attrs {
		akeys = append(akeys, k)
	}

	sort.Strings(akeys)

	for _, k := range akeys {
		sb.WriteByte(' ')
		sb.WriteString(k)
		sb.WriteString(`="`)
		sb.WriteString(ot.attrs[k])
		sb.WriteByte('"')
	}

	sb.WriteByte(']')

	return sb.String()
}

// TagName satifies to the Node interface.
func (ot *OpeningTag) TagName() string {
	return ot.tok.Literal
}

// Value returns a bbcode tag value (string after =).
func (ot *OpeningTag) Value() string {
	return ot.value
}

// Attrs returns a tag attributes.
func (ot *OpeningTag) Attrs() map[string]string {
	return ot.attrs
}

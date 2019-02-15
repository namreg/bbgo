package node

import (
	"sort"
	"strings"
	"sync"

	"github.com/namreg/bbgo/token"
)

var selfClosingTagPool = sync.Pool{
	New: func() interface{} {
		return &SelfClosingTag{}
	},
}

var _ Tag = (*SelfClosingTag)(nil)

// SelfClosingTag is self-closing bbcode tag, i.e. [url="https://google.com" /]
type SelfClosingTag struct {
	tok   token.Token
	value string
	attrs map[string]string
}

// NewSelfClosingTag creates a new opening tag.
func NewSelfClosingTag(tok token.Token, value string, attrs map[string]string) *SelfClosingTag {
	sct := selfClosingTagPool.Get().(*SelfClosingTag)
	sct.tok = tok
	sct.value = value
	sct.attrs = attrs
	return sct
}

// Token satisfies to the Node interface.
func (ot *SelfClosingTag) Token() token.Token {
	return ot.tok
}

// String satisfies to the Node interface.
func (ot *SelfClosingTag) String() string {
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

	sb.Write([]byte{' ', '/', ']'})

	return sb.String()
}

// TagName satifies to the Node interface.
func (ot *SelfClosingTag) TagName() string {
	return ot.tok.Literal
}

// Value returns a bbcode tag value (string after =).
func (ot *SelfClosingTag) Value() string {
	return ot.value
}

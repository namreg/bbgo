package parser_test

import (
	"reflect"
	"testing"

	"github.com/namreg/bbgo/lexer"
	"github.com/namreg/bbgo/node"
	"github.com/namreg/bbgo/parser"
	"github.com/namreg/bbgo/token"
)

func TestParse(t *testing.T) {
	token.RegisterIdentifiers("url", "b", "size")

	input := `[url="https://google.com"  foo=bar buzz="bar bar" /][b]text[[/b][/b][/foo]bar[][size=111]`

	expectedNodes := []node.Node{
		node.NewSelfClosingTag(token.Token{Kind: token.IDENT, Literal: "url"}, `https://google.com`, map[string]string{"foo": "bar", "buzz": "bar bar"}),
		node.NewOpeningTag(token.Token{Kind: token.IDENT, Literal: "b"}, "", nil),
		node.NewText(token.Token{Kind: token.STRING, Literal: "text"}, "text"),
		node.NewText(token.Token{Kind: token.STRING, Literal: "["}, "["),
		node.NewClosingTag(token.Token{Kind: token.IDENT, Literal: "b"}),
		node.NewClosingTag(token.Token{Kind: token.IDENT, Literal: "b"}),
		node.NewText(token.Token{Kind: token.STRING, Literal: "["}, "[/foo"),
		node.NewText(token.Token{Kind: token.STRING, Literal: "]"}, "]"),
		node.NewText(token.Token{Kind: token.STRING, Literal: "bar"}, "bar"),
		node.NewText(token.Token{Kind: token.STRING, Literal: "["}, "["),
		node.NewText(token.Token{Kind: token.STRING, Literal: "]"}, "]"),
		node.NewOpeningTag(token.Token{Kind: token.IDENT, Literal: "size"}, "111", nil),
		node.NewText(token.Token{Kind: token.STRING, Literal: "["}, `[size="300%]`),
	}

	lex := lexer.New(input)
	p := parser.New(lex)

	i := 0
	for actual := range p.Parse() {
		if i >= len(expectedNodes) {
			t.Fatalf("unexpected node: %+v", actual)
		}
		expected := expectedNodes[i]
		if expected.String() != actual.String() {
			t.Fatalf("Unexpected value node #%d. Want = %s, got = %s", i, expected, actual)
		}
		if !reflect.DeepEqual(expected.Token(), actual.Token()) {
			t.Fatalf("Unexpected token node #%d. Want = %+v, got = %+v", i, expected.Token(), actual.Token())
		}
		i++
	}
}

func TestParse2(t *testing.T) {
	token.RegisterIdentifiers("url", "b", "size")

	input := "[b]hello\nworld[/b]"

	expectedNodes := []node.Node{
		node.NewOpeningTag(token.Token{Kind: token.IDENT, Literal: "b"}, "", nil),
		node.NewText(token.Token{Kind: token.STRING, Literal: "hello"}, "hello"),
		node.NewLine(token.Token{Kind: token.NL, Literal: "\n"}),
		node.NewText(token.Token{Kind: token.STRING, Literal: "world"}, "world"),
		node.NewClosingTag(token.Token{Kind: token.IDENT, Literal: "b"}),
	}

	lex := lexer.New(input)
	p := parser.New(lex)

	i := 0
	for actual := range p.Parse() {
		if i >= len(expectedNodes) {
			t.Fatalf("unexpected node: %+v", actual)
		}
		expected := expectedNodes[i]
		if expected.String() != actual.String() {
			t.Fatalf("Unexpected value node #%d. Want = %s, got = %s", i, expected, actual)
		}
		if !reflect.DeepEqual(expected.Token(), actual.Token()) {
			t.Fatalf("Unexpected token node #%d. Want = %+v, got = %+v", i, expected.Token(), actual.Token())
		}
		i++
	}
}

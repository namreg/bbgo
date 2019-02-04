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
	input := `[b]text[[/b][/b][/foo]bar[][size=111][size="300%]`

	expectedNodes := []node.Node{
		node.NewOpeningTag(token.Token{Kind: token.IDENT, Literal: "b"}, ""),
		node.NewText(token.Token{Kind: token.STRING, Literal: "text"}, "text["),
		node.NewClosingTag(token.Token{Kind: token.IDENT, Literal: "b"}),
		node.NewClosingTag(token.Token{Kind: token.IDENT, Literal: "b"}),
		node.NewText(token.Token{Kind: token.STRING, Literal: "["}, "[/foo]bar[]"),
		node.NewOpeningTag(token.Token{Kind: token.IDENT, Literal: "size"}, "111"),
		node.NewText(token.Token{Kind: token.STRING, Literal: "["}, `[size="300%]`),
	}

	lex := lexer.New(input)
	p := parser.New(lex)

	actualNodes := p.Parse()

	for i, expected := range expectedNodes {
		actual := actualNodes[i]
		if expected.String() != actual.String() {
			t.Fatalf("Unexpected value node #%d. Want = %s, got = %s", i, expected, actual)
		}
		if !reflect.DeepEqual(expected.Token(), actual.Token()) {
			t.Fatalf("Unexpected token node #%d. Want = %+v, got = %+v", i, expected.Token(), actual.Token())
		}

	}
}

package lexer_test

import (
	"testing"

	"github.com/namreg/bbgo/lexer"
	"github.com/namreg/bbgo/token"
)

func TestNextToken(t *testing.T) {
	input := `[foo=bar][url=https://google.com /][quote="автор цитаты"]цитата[/quote][size=400%]hi[/size][b]bold][[/b][b]text`

	tests := []struct {
		expectedKind    token.Kind
		expectedLiteral string
	}{
		{token.LBRACKET, "["},
		{token.STRING, "foo"},
		{token.EQUAL, "="},
		{token.STRING, "bar"},
		{token.RBRACKET, "]"},

		{token.LBRACKET, "["},
		{token.IDENT, "url"},
		{token.EQUAL, "="},
		{token.STRING, "https://google.com "},
		{token.SLASH, "/"},
		{token.RBRACKET, "]"},

		{token.LBRACKET, "["},
		{token.IDENT, "quote"},
		{token.EQUAL, "="},
		{token.QUOTE, `"`},
		{token.STRING, "автор цитаты"},
		{token.QUOTE, `"`},
		{token.RBRACKET, "]"},
		{token.STRING, "цитата"},
		{token.LBRACKET, "["},
		{token.SLASH, "/"},
		{token.IDENT, "quote"},
		{token.RBRACKET, "]"},

		{token.LBRACKET, "["},
		{token.IDENT, "size"},
		{token.EQUAL, "="},
		{token.STRING, "400%"},
		{token.RBRACKET, "]"},
		{token.STRING, "hi"},
		{token.LBRACKET, "["},
		{token.SLASH, "/"},
		{token.IDENT, "size"},
		{token.RBRACKET, "]"},

		{token.LBRACKET, "["},
		{token.IDENT, "b"},
		{token.RBRACKET, "]"},
		{token.STRING, "bold"},
		{token.RBRACKET, "]"},
		{token.LBRACKET, "["},
		{token.LBRACKET, "["},
		{token.SLASH, "/"},
		{token.IDENT, "b"},
		{token.RBRACKET, "]"},

		{token.LBRACKET, "["},
		{token.IDENT, "b"},
		{token.RBRACKET, "]"},
		{token.STRING, "text"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tt.expectedKind != tok.Kind {
			t.Fatalf("Test #%d failed (Unexpected kind). Want = %v, got = %v", i, tt.expectedKind, tok.Kind)
		}
		if tt.expectedLiteral != tok.Literal {
			t.Fatalf("Test #%d failed (Unexpected literal). Want = %v, got = %v", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken2(t *testing.T) {
	input := `[b]text[url="https://google.com" /][size="300%]`

	tests := []struct {
		expectedKind    token.Kind
		expectedLiteral string
	}{
		{token.LBRACKET, "["},
		{token.IDENT, "b"},
		{token.RBRACKET, "]"},
		{token.STRING, "text"},

		{token.LBRACKET, "["},
		{token.IDENT, "url"},
		{token.EQUAL, "="},
		{token.QUOTE, `"`},
		{token.STRING, "https://google.com"},
		{token.QUOTE, `"`},
		{token.SLASH, "/"},
		{token.RBRACKET, "]"},

		{token.LBRACKET, "["},
		{token.IDENT, "size"},
		{token.EQUAL, "="},
		{token.QUOTE, `"`},
		{token.STRING, "300%]"},

		{token.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tt.expectedKind != tok.Kind {
			//t.Fatalf("Test #%d failed (Unexpected kind). Want = %v, got = %v", i, tt.expectedKind, tok.Kind)
		}
		if tt.expectedLiteral != tok.Literal {
			t.Fatalf("Test #%d failed (Unexpected literal). Want = %v, got = %v", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

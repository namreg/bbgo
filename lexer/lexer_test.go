package lexer_test

import (
	"testing"

	"github.com/namreg/bbgo/lexer"
	"github.com/namreg/bbgo/token"
)

func TestNextToken(t *testing.T) {
	input := `[b]bold][[/b]`

	tests := []struct {
		expectedKind    token.Kind
		expectedLiteral string
	}{
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

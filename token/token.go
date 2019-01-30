package token

// Kind is a token kind.
type Kind string

// Token is produced by Lexer.
type Token struct {
	Kind    Kind
	Literal string
}

const (
	// EOF indicates that we are at the end.
	EOF Kind = "EOF"

	// IDENT is a token that represents a tag name.
	IDENT Kind = "IDENT"

	// STRING is string token.
	STRING Kind = "STRING"

	// LBRACKET is a `[` token.
	LBRACKET Kind = "["
	// RBRACKET is a `]` token.
	RBRACKET Kind = "]"
	// SLASH is `/` token.
	SLASH Kind = "/"
	// EQUAL is a `=` token.
	EQUAL Kind = "="
	// QUOTE is a `"` token.
	QUOTE Kind = `"`
)

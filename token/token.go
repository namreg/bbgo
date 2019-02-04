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

var identifiers = map[string]struct{}{
	"b":     {},
	"quote": {},
	"url":   {},
	"img":   {},
	"size":  {},
}

// RegisterIdentifier registers a new identifier.
func RegisterIdentifier(ident string) {
	identifiers[ident] = struct{}{}
}

// IsValidIndetifier determines whether the given identifier is valid.
func IsValidIndetifier(ident string) bool {
	_, ok := identifiers[ident]
	return ok
}

// IsEmpty determines whether the given token is empty.
func IsEmpty(t Token) bool {
	return t == Token{}
}

package lexer

import (
	"github.com/namreg/bbgo/token"
)

// Lexer retreives tokens from input string.
type Lexer struct {
	input          []rune
	position       int  // current position in input (points to current char)
	readPosition   int  // current reading position in input (after current char)
	ch             rune // current char under examination
	insideBrackets bool // indicates whether the current position inside the brackets
}

// New creates a new Lexer and initialized it with the input string.
func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}

// NextToken returns next available token from the underlying input.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '[':
		l.insideBrackets = true
		tok = l.newToken(token.LBRACKET, l.ch)
	case ']':
		l.insideBrackets = false
		tok = l.newToken(token.RBRACKET, l.ch)
	case '=':
		tok = l.newToken(token.EQUAL, l.ch)
	case '"':
		tok = l.newToken(token.QUOTE, l.ch)
	case '/':
		tok = l.newToken(token.SLASH, l.ch)
	case 0:
		tok.Kind = token.EOF
	default:
		if l.insideBrackets && l.isValidIdentifierRune(l.ch) {
			tok.Kind = token.IDENT
			tok.Literal = l.readIdentifier()
			return tok
		}

		tok.Kind = token.STRING
		tok.Literal = l.readString()
		return tok
	}

	l.readChar()
	return tok
}

func (l *Lexer) runeInSlice(r rune, s []rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}

func (l *Lexer) readString() string {
	position := l.position
	for !l.runeInSlice(l.ch, []rune{'[', ']', '/', '=', '"'}) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.isValidIdentifierRune(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) isValidIdentifierRune(r rune) bool {
	// TODO(namreg): this is enough?
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '-'
}

func (l *Lexer) newToken(kind token.Kind, ch rune) token.Token {
	return token.Token{Kind: kind, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

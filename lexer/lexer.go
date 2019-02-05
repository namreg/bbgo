package lexer

import (
	"unicode"

	"github.com/namreg/bbgo/token"
)

// Lexer retreives tokens from input string.
type Lexer struct {
	input          []rune
	position       int  // current position in input (points to current rune)
	readPosition   int  // current reading position in input (after current rune)
	ch             rune // current rune under examination
	insideBrackets bool // indicates whether the current position inside the brackets
	insideQuote    bool // indicates whether the current position inside the quote
	insideAttr     bool // indicates whether the current position inside the attribute
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

	if l.insideBrackets && !l.insideAttr {
		l.skipWhitespaces() // ignore whitwspaces inside brackets
	}

	switch l.ch {
	case '[':
		l.insideBrackets = true
		tok = l.newToken(token.LBRACKET, l.ch)
	case ']':
		l.insideBrackets = false
		tok = l.newToken(token.RBRACKET, l.ch)
	case '=':
		l.insideAttr = l.insideBrackets
		tok = l.newToken(token.EQUAL, l.ch)
	case '"':
		l.insideQuote = !l.insideQuote
		tok = l.newToken(token.QUOTE, l.ch)
	case '/':
		tok = l.newToken(token.SLASH, l.ch)
	// TODO(namreg): NEWLINE token
	case 0:
		tok.Kind = token.EOF
	default:
		if l.insideBrackets {
			if l.insideAttr {
				tok.Kind = token.STRING
				tok.Literal = l.readAttrValue()
				l.insideAttr = false
				return tok
			}

			if l.isValidIdentifierRune(l.ch) {
				kind := token.IDENT
				ident := l.readIdentifier()
				if !token.IsValidIndetifier(ident) { // check whether the readed indentifier defined as allowed
					kind = token.STRING
				}
				tok.Kind = kind
				tok.Literal = ident
				return tok
			}
		}

		tok.Kind = token.STRING
		tok.Literal = l.readString()
		return tok
	}

	l.readChar()
	return tok
}

func (l *Lexer) readUntil(r ...rune) string {
	position := l.position
	for !l.runeInSlice(l.ch, r) && l.ch != 0 {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readAttrValue() string {
	if l.insideQuote {
		return l.readUntil('"')
	}
	position := l.position
LOOP:
	for {
		switch {
		case l.ch == 0 || l.ch == ' ':
			break LOOP
		case l.peekChar() == ']' || (l.peekChar() == '/' && l.peek2Char() == ']'):
			l.readChar() // due to the peekChar
			break LOOP
		}
		l.readChar()
	}
	return string(l.input[position:l.position])
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
	for !l.runeInSlice(l.ch, []rune{0, '[', ']', '/', '=', '"'}) {
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

func (l *Lexer) skipWhitespaces() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
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

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) peek2Char() rune {
	if l.readPosition+1 >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition+1]
}

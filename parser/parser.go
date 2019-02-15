package parser

import (
	"strings"

	"github.com/namreg/bbgo/lexer"
	"github.com/namreg/bbgo/node"
	"github.com/namreg/bbgo/token"
)

const internalBufCap = 20

// Parser parses tokens into nodes.
type Parser struct {
	lex *lexer.Lexer

	prevToken token.Token
	currToken token.Token
	peekToken token.Token

	buf []token.Token // internal buffer
}

// New creates a new parser.
func New(lex *lexer.Lexer) *Parser {
	p := &Parser{
		lex: lex,
		buf: make([]token.Token, 0, internalBufCap),
	}

	// set the currToken and the peekToken
	p.nextToken()
	p.nextToken()

	return p
}

// Parse parses an entire input and returns resulting nodes.
func (p *Parser) Parse() chan node.Node {
	ch := make(chan node.Node)

	go p.parse(ch)

	return ch
}

func (p *Parser) parse(ch chan node.Node) {
LOOP:
	for !p.currTokenIs(token.EOF) {
	SWITCH:
		switch {
		case p.currTokenIs(token.LBRACKET) && p.peekTokenIs(token.SLASH): // seems to be a closing tag
			p.nextToken()
			p.nextToken()
			if !p.currTokenIs(token.IDENT) || !p.peekTokenIs(token.RBRACKET) {
				p.drainBuf(ch)
				break SWITCH
			}
			ch <- node.NewClosingTag(p.currToken)
			p.nextToken()
		case p.currTokenIs(token.LBRACKET) && p.peekTokenIs(token.IDENT): // seems to be an opening tag
			p.nextToken()

			tagToken := p.currToken

			val := p.readTagValue()

			attrs := make(map[string]string)

			for !p.currTokenIs(token.RBRACKET) {
				if ak, av := p.readAttrKeyValue(); ak != "" && av != "" {
					attrs[ak] = av
				}
				if !p.nextToken() {
					break LOOP // we are at the end of the input. (EOF)
				}
			}
			if p.prevTokenIs(token.SLASH) { // seems to be a self-closing tag
				ch <- node.NewSelfClosingTag(tagToken, val, attrs)
			} else {
				ch <- node.NewOpeningTag(tagToken, val, attrs)
			}
		case p.currTokenIs(token.NL):
			ch <- node.NewLine(p.currToken)
		default:
			p.drainBuf(ch)
		}
		p.resetBuf()
		p.nextToken()
	}
	p.drainBuf(ch)
	close(ch)
}

func (p *Parser) readTagValue() string {
	if p.peekTokenIs(token.EQUAL) { // seems to be a tag value
		p.nextToken()
		if p.peekTokenIs(token.QUOTE) { // a value within quotes
			p.nextToken()
			if p.peekTokenIs(token.STRING) {
				return p.peekToken.Literal
			}
		} else if p.peekTokenIs(token.STRING) {
			return p.peekToken.Literal
		}
	}
	return ""
}

func (p *Parser) readAttrKeyValue() (string, string) {
	if p.currTokenIs(token.EQUAL) && !p.prevTokenIs(token.IDENT) { // seems to an attribute
		ak := p.prevToken.Literal
		if p.peekTokenIs(token.QUOTE) {
			p.nextToken()
		}
		if p.peekTokenIs(token.STRING) {
			p.nextToken()
			av := p.currToken.Literal
			return ak, av
		}
		return ak, ""
	}
	return "", ""
}

func (p *Parser) resetBuf() {
	p.buf = p.buf[:0]
}

func (p *Parser) bufToString() string {
	var sb strings.Builder

	for _, t := range p.buf {
		sb.WriteString(t.Literal)
	}

	return sb.String()
}

func (p *Parser) drainBuf(ch chan node.Node) {
	if len(p.buf) == 0 {
		return
	}
	defer p.resetBuf()

	str := p.bufToString()
	tok := token.Token{Kind: token.STRING, Literal: p.buf[0].Literal}

	ch <- node.NewText(tok, str)
}

func (p *Parser) currTokenIs(k token.Kind) bool {
	return p.currToken.Kind == k
}

func (p *Parser) peekTokenIs(k token.Kind) bool {
	return p.peekToken.Kind == k
}

func (p *Parser) prevTokenIs(k token.Kind) bool {
	if token.IsEmpty(p.prevToken) {
		return false
	}
	return k == p.prevToken.Kind
}

func (p *Parser) nextToken() bool {
	p.prevToken = p.currToken
	p.currToken = p.peekToken
	p.peekToken = p.lex.NextToken()

	if !token.IsEmpty(p.currToken) && !p.currTokenIs(token.EOF) {
		p.buf = append(p.buf, p.currToken)
	}

	return p.currToken.Kind != token.EOF
}

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
	lex       *lexer.Lexer
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
func (p *Parser) Parse() []node.Node {
	nodes := make([]node.Node, 0)
LOOP:
	for p.currToken.Kind != token.EOF {
		switch {
		case p.currTokenIs(token.LBRACKET) && p.peekTokenIs(token.SLASH): // seems to be a closing tag
			p.nextToken()
			p.nextToken()
			if p.currToken.Kind != token.IDENT || p.peekToken.Kind != token.RBRACKET {
				nodes = p.drainBuf(nodes)
				break
			}
			nodes = append(nodes, node.NewClosingTag(p.currToken))
			p.nextToken()
		case p.currTokenIs(token.LBRACKET) && p.peekTokenIs(token.IDENT): // seems to be an opening tag
			tag := p.peekToken
			var attr string
			for !p.currTokenIs(token.RBRACKET) {
				if !p.nextToken() {
					break LOOP
				}
			}
			nodes = append(nodes, node.NewOpeningTag(tag, attr))
		default:
			nodes = p.drainBuf(nodes)
		}
		p.resetBuf()
		p.nextToken()
	}
	return p.drainBuf(nodes)
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

func (p *Parser) drainBuf(nodes []node.Node) []node.Node {
	if len(p.buf) == 0 {
		return nodes
	}
	defer p.resetBuf()
	str := p.bufToString()
	if len(nodes) > 0 {
		last := nodes[len(nodes)-1]
		if tn, ok := last.(*node.Text); ok {
			tn.Append(str)
			return nodes
		}
	}
	tok := token.Token{Kind: token.STRING, Literal: p.buf[0].Literal}
	return append(nodes, node.NewText(tok, str))
}

func (p *Parser) currTokenIs(k token.Kind) bool {
	return p.currToken.Kind == k
}

func (p *Parser) peekTokenIs(k token.Kind) bool {
	return p.peekToken.Kind == k
}

func (p *Parser) nextToken() bool {
	p.currToken = p.peekToken
	p.peekToken = p.lex.NextToken()

	if (p.currToken != token.Token{}) && (p.currToken.Kind != token.EOF) {
		p.buf = append(p.buf, p.currToken)
	}

	return p.currToken.Kind != token.EOF
}

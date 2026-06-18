package parser

import (
	"fmt"


	"jsonParser/lexer"
	"jsonParser/token"
)

type Parser struct {
	lexer *lexer.Lexer
	currentToken token.Token
	peekToken token.Token 
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: l,
	}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseObject() error {
	currentToken := p.currentToken

	if currentToken.Type != token.LEFT_BRACE {
		return fmt.Errorf("Current Token is not {")
	} 

	p.nextToken()

	currentToken = p.currentToken

	if currentToken.Type !=  token.RIGHT_BRACE {
		return fmt.Errorf("Current Token is not Valid")
	}

	return nil 
}

func (p *Parser) Parse() error{
	err := p.parseObject()

	if err != nil {
		return err
	}

	p.nextToken()

	if p.currentToken.Type != token.EOF {
		return fmt.Errorf("Current Input is not a valid json string")
	}

	return nil
}
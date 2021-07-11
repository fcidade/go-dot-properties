package parser

import "fmt"

type parser struct {
	tokens []Token
	cursor int
}

func NewParser(tokens []Token) *parser {
	return &parser{
		tokens: tokens,
		cursor: 0,
	}
}

func (p *parser) ParseToMap() (out map[string]string, err error) {
	out = make(map[string]string)

	for p.cursor < len(p.tokens) {

		token, err := p.currToken()
		if err != nil {
			return out, ExpectingTokenError(TypeIdentifier)
		}
		if token.Type != TypeIdentifier {
			return out, InvalidTokenTypeError(TypeIdentifier, token.Type)
		}
		key := token.Text
		p.nextToken()

		token, err = p.currToken()
		if err != nil {
			return out, ExpectingTokenError(TypeSeparator)
		}
		if token.Type != TypeSeparator {
			return out, InvalidTokenTypeError(TypeSeparator, token.Type)
		}
		p.nextToken()

		token, err = p.currToken()
		if err != nil {
			return out, ExpectingTokenError(TypeValue)
		}
		if token.Type != TypeValue {
			return out, InvalidTokenTypeError(TypeValue, token.Type)
		}
		out[key] = token.Text
		p.nextToken()

	}

	return out, nil
}

func (p *parser) currToken() (token Token, err error) {
	if p.cursor >= len(p.tokens) {
		return token, fmt.Errorf("out of bounds")
	}
	return p.tokens[p.cursor], nil
}

func (p *parser) nextToken() {
	p.cursor++
}

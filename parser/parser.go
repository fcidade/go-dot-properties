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

	for p.cursor <= len(p.tokens) {

		token, err := p.currToken()

		if err != nil {
			return out, fmt.Errorf("expecting identifier, found nothing")
		}
		if token.Type != TypeIdentifier {
			return out, fmt.Errorf("expecting identifier, found %s", token.Type) // TODO, custom error w/ p.cursor
		}
		key := token.Text
		p.nextToken()

		token, err = p.currToken()

		if err != nil {
			return out, fmt.Errorf("expecting missing, found nothing")
		}
		if token.Type != TypeSeparator {
			return out, fmt.Errorf("separator missing, found %s", token.Type) // TODO, custom error w/ p.cursor
		}
		p.nextToken()

		token, err = p.currToken()

		if err != nil {
			return out, fmt.Errorf("expecting value, found nothing")
		}
		if token.Type != TypeValue {
			return out, fmt.Errorf("expecting value, found %s", token.Type) // TODO, custom error w/ p.cursor
		}
		out[key] = token.Text
		p.nextToken()

	}

	return
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

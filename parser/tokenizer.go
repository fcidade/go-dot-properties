package parser

import "strconv"

const (
	TypeIdentifier = "identifier"
	TypeValue      = "value"
	TypeSeparator  = "separator"
)

const (
	EOF = 0x00
)

type Token struct {
	Text string
	Type string
}

type tokenizer struct {
	text   string
	cursor int
	tokens []Token
}

func (t *tokenizer) Tokenize() []Token {
	for t.cursor < len(t.text) {
		t.ignoreWhitespaces()
		t.ignoreComments()
		t.recognizeIdentifier()
		t.recognizeValue()

		t.nextChar()
	}

	return t.tokens
}

func (t *tokenizer) ignoreWhitespaces() {
	// TODO: maybe change to !isAlphaNumeric?
	for t.currChar() == ' ' || t.currChar() == '\n' || t.currChar() == '\r' {
		t.nextChar()
	}
}

func (t *tokenizer) ignoreComments() {
	if t.currChar() == '#' || t.currChar() == '!' {
		t.nextChar()

		for t.currChar() != '\n' {
			t.nextChar()
		}
	}
}

func (t *tokenizer) recognizeIdentifier() {
	if t.isAlpha(t.currChar()) {
		begin := t.cursor
		t.nextChar()

		identifier := ""

		for {
			if t.isAlpha(t.currChar()) {
				t.nextChar()
			} else if t.currChar() == '\\' {
				identifier += t.text[begin:t.cursor]
				t.nextChar()
				begin = t.cursor
				t.nextChar()
			} else {
				break
			}
		}

		identifier += t.text[begin:t.cursor]
		t.createToken(identifier, TypeIdentifier)
	}
}

func (t *tokenizer) recognizeValue() {

	if t.currChar() == '=' || t.currChar() == ':' {
		separator := string(t.currChar())
		t.createToken(separator, TypeSeparator)
		t.nextChar()

		t.ignoreWhitespaces()

		begin := t.cursor
		value := ""

		// for t.currChar() != '\n' && t.currChar() != '\\' {
		for {

			if t.currChar() == EOF {
				break
			}

			// if t.currChar() == '\r'{
			// 	value += t.text[begin:t.cursor]
			// 	t.nextChar()
			// 	begin = t.cursor
			// }

			if t.currChar() == '\n' {
				break
			}

			if t.currChar() == '\\' {
				value += t.text[begin:t.cursor]
				t.nextChar()
				begin = t.cursor

				if t.currChar() == 'u' {
					t.nextChar()
					unicodeStr := ""
					for t.isNumeric(t.currChar()) {
						unicodeStr += string(rune(t.currChar()))
						t.nextChar()
					}
					if len(unicodeStr) != 0 {
						value = t.handleUnicode(unicodeStr)
					}
					begin = t.cursor
				}

				for t.currChar() == '\n' || t.currChar() == '\t' {
					t.nextChar()
					begin = t.cursor
				}
			}

			t.nextChar()
		}

		value += t.text[begin:t.cursor]
		t.createToken(value, TypeValue)
	}
}

func (t *tokenizer) handleUnicode(in string) (out string) {
	numeric, err := strconv.Atoi(in)
	if err != nil {
		return ""
	}
	return string(rune(numeric))
}

func (t *tokenizer) currChar() byte {
	if t.cursor >= len(t.text) {
		return EOF
	}
	return t.text[t.cursor]
}

func (t *tokenizer) nextChar() {
	if t.cursor < len(t.text) {
		t.cursor++
	}
}

func (t *tokenizer) createToken(text, tokenType string) {
	t.tokens = append(t.tokens, Token{
		Text: text,
		Type: tokenType,
	})
}

func (t *tokenizer) isAlpha(currChar byte) bool {
	return (currChar >= 'a' && currChar <= 'z') || currChar >= 'A' && currChar <= 'Z'
}

func (t *tokenizer) isNumeric(currChar byte) bool {
	return currChar >= '0' && currChar <= '9'
}

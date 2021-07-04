package parser

const (
	TypeIdentifier = "identifier"
	TypeValue      = "value"
	TypeSeparator  = "separator"
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

func (t *tokenizer) currChar() byte {
	return t.text[t.cursor]
}

func (t *tokenizer) nextChar() {
	t.cursor++
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

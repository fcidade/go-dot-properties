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

		t.ignoreComments()
		t.recognizeIdentifier()

		// 	// for t.currChar() == ' ' {
		// 	// 	t.nextChar()
		// 	// }

		t.recognizeValue()

		// 	// for t.currChar() == ' ' {
		// 	// 	t.nextChar()
		// 	// }

		t.nextChar()
	}

	return t.tokens
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
		for t.isAlpha(t.currChar()) {
			t.nextChar()
		}

		identifier := t.text[begin:t.cursor]
		t.createToken(identifier, TypeIdentifier)
	}
}

func (t *tokenizer) recognizeValue() {
	if t.currChar() == '=' || t.currChar() == ':' {
		separator := string(t.currChar())
		t.createToken(separator, TypeSeparator)
		t.nextChar()

		// 	for t.currChar() == ' ' {
		// 		t.nextChar()
		// 	}

		begin := t.cursor
		joinedText := ""

		for t.currChar() != '\n' && t.currChar() != '\\' {

			if t.currChar() == '\\' || t.currChar() == '\r' {
				joinedText += t.text[begin:t.cursor]
				begin = t.cursor
				t.nextChar()
			}

			t.nextChar()
		}

		value := joinedText + t.text[begin:t.cursor]
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

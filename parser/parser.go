package parser

const (
	TypeIdentifier = "variable"
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
}

func (t *tokenizer) Tokenize() (tokens []Token) {
	t.cursor = 0
	for t.cursor < len(t.text) {

		if t.currChar() == '#' || t.currChar() == '!' {
			t.nextChar()

			for t.currChar() != '\n' {
				t.nextChar()
			}
		}

		if t.isAlpha(t.currChar()) {
			begin := t.cursor
			t.nextChar()
			for t.isAlpha(t.currChar()) {
				t.nextChar()
			}
			tokens = append(tokens, Token{
				Text: t.text[begin:t.cursor],
				Type: TypeIdentifier,
			})
		}

		// 	// for t.currChar() == ' ' {
		// 	// 	t.nextChar()
		// 	// }

		if t.currChar() == '=' || t.currChar() == ':' {
			tokens = append(tokens, Token{
				Text: string(t.currChar()),
				Type: TypeSeparator,
			})
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

			tokens = append(tokens, Token{
				Text: joinedText + t.text[begin:t.cursor],
				Type: TypeValue,
			})
		}

		// 	// for t.currChar() == ' ' {
		// 	// 	t.nextChar()
		// 	// }

		t.nextChar()
	}
	return
}

func (t *tokenizer) currChar() byte {
	return t.text[t.cursor]
}

func (t *tokenizer) nextChar() {
	t.cursor++
}

func (t *tokenizer) isAlpha(currChar byte) bool {
	return (currChar >= 'a' && currChar <= 'z') || currChar >= 'A' && currChar <= 'Z'
}

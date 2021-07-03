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

func Tokenize(text string) (tokens []Token) {
	cursor := 0
	for cursor < len(text) {
		currChar := text[cursor]

		if currChar == '#' || currChar == '!' {
			cursor++
			currChar = text[cursor]
			for currChar != '\n' {
				cursor++
				currChar = text[cursor]
			}
		}

		if isAlpha(currChar) {
			begin := cursor
			cursor++
			currChar = text[cursor]
			for isAlpha(currChar) {
				cursor++
				currChar = text[cursor]
			}
			tokens = append(tokens, Token{
				Text: text[begin:cursor],
				Type: TypeIdentifier,
			})
		}

		// for currChar == ' ' {
		// 	cursor++
		// 	currChar = text[cursor]
		// }

		if currChar == '=' || currChar == ':' {
			tokens = append(tokens, Token{
				Text: string(currChar),
				Type: TypeSeparator,
			})
			cursor++

			// 	for currChar == ' ' {
			// 		cursor++
			// 		currChar = text[cursor]
			// 	}

			begin := cursor
			joinedText := ""

			for currChar != '\n' && currChar != '\\' {

				if currChar == '\\' || currChar == '\r' {
					joinedText += text[begin:cursor]
					begin = cursor
					cursor++
				}

				cursor++
				currChar = text[cursor]
			}

			tokens = append(tokens, Token{
				Text: joinedText + text[begin:cursor],
				Type: TypeValue,
			})
		}

		// for currChar == ' ' {
		// 	cursor++
		// 	currChar = text[cursor]
		// }

		cursor++
	}
	return
}

func isAlpha(currChar byte) bool {
	return (currChar >= 'a' && currChar <= 'z') || currChar >= 'A' && currChar <= 'Z'
}

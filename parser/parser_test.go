package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeParserSut(tokens []Token) *parser {
	return NewParser(tokens)
}

func TestParser(t *testing.T) {
	t.Run("Map tokens to a map of strings", func(t *testing.T) {
		tokens := []Token{{
			Text: "website",
			Type: TypeIdentifier,
		}, {
			Text: "=",
			Type: TypeSeparator,
		}, {
			Text: "https://en.wikipedia.org/",
			Type: TypeValue,
		}, {
			Text: "language",
			Type: TypeIdentifier,
		}, {
			Text: ":",
			Type: TypeSeparator,
		}, {
			Text: "English",
			Type: TypeValue,
		}}

		parser := makeParserSut(tokens)

		want := map[string]string{
			"website":  "https://en.wikipedia.org/",
			"language": "English",
		}
		have, err := parser.ParseToMap()

		fmt.Println(err)

		assert.Equal(t, want, have)
	})
}

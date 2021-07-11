package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	t.Run("Should generate a colon separated text from a map", func(t *testing.T) {
		mp := map[string]string{
			"website":  "https://en.wikipedia.org/",
			"language": "English",
		}
		have := GenerateFromMap(mp)
		want := `website: https://en.wikipedia.org/
language: English
`
		fmt.Println(have)
		assert.Equal(t, want, have)
	})

	t.Run("Should generate a colon separated text from a struct", func(t *testing.T) {
		strc := struct {
			Website   string `properties:"website"`
			Language  string `properties:"language"`
			language2 string
		}{
			Website:   "https://en.wikipedia.org/",
			Language:  "English",
			language2: "English",
		}
		have := GenerateFromStruct(strc)
		want := `website: https://en.wikipedia.org/
language: English
language2: English
`
		fmt.Println(have)
		assert.Equal(t, want, have)
	})
}

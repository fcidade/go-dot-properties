package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {

	t.Run("Should generate a colon separated text", func(t *testing.T) {
		mp := map[string]string{
			"website":  "https://en.wikipedia.org/",
			"language": "English",
		}
		have := GenerateFromMap(mp)
		want := `website: https://en.wikipedia.org/
language: English
`
		assert.Equal(t, want, have)
	})
}

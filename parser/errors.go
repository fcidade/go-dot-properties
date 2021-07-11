package parser

import "fmt"

type syntaxError struct {
	expecting string
	found     string
}

func (e *syntaxError) Error() (out string) {
	return fmt.Sprintf("expecting %s, found %s", e.expecting, e.found)
}

func InvalidTokenTypeError(expecting, found string) error {
	return &syntaxError{
		expecting: expecting,
		found:     found,
	}
}

func ExpectingTokenError(expecting string) error {
	return &syntaxError{
		expecting: expecting,
		found:     "nothing",
	}
}

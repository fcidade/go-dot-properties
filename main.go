package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-dot-properties/parser"
)

func Parse(text string) (interface{}, error) {
	tokenizer := parser.NewTokenizer(text)
	tokens := tokenizer.Tokenize()
	pars := parser.NewParser(tokens)
	return pars.ParseToMap()
}

func main() {

	fmt.Println(os.Args)
	filepath := ""

	if len(os.Args) >= 2 {
		filepath = os.Args[1]
	}

	bytes, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	text := string(bytes)

	mp, err := Parse(text)
	fmt.Printf("%#v, %e\n", mp, err)

	js, err := json.Marshal(mp)
	fmt.Println(string(js), err)

}

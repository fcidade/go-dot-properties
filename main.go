package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-dot-properties/parser"
)

func main() {
	ParseToJSON()
}

func ParseToJSON() {

	filepath := ""

	if len(os.Args) >= 2 {
		filepath = os.Args[1]
	}

	bytes, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	text := string(bytes)
	tokenizer := parser.NewTokenizer(text)
	tokens := tokenizer.Tokenize()
	pars := parser.NewParser(tokens)

	mp, err := pars.ParseToMap()
	if err != nil {
		panic(err)
	}

	js, err := json.Marshal(mp)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(js), err)
}

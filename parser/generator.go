package parser

import (
	"fmt"
)

func GenerateFromMap(in map[string]string) (out string) {
	for key, value := range in {
		out += fmt.Sprintf("%s: %s\n", key, value)
	}
	return
}

func GenerateFromStruct(in interface{}) (out string) {
	return GenerateFromMap(structToMap(in))
}

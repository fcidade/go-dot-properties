package parser

import (
	"fmt"
	"reflect"
)

func GenerateFromMap(in map[string]string) (out string) {
	for key, value := range in {
		out += fmt.Sprintf("%s: %s\n", key, value)
	}
	return
}

func GenerateFromStruct(in interface{}) (out string) {
	fields := reflect.TypeOf(in)
	values := reflect.ValueOf(in)

	num := fields.NumField()

	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)

		tagName := field.Tag.Get("properties")
		if len(tagName) == 0 {
			tagName = field.Name
		}
		out += fmt.Sprintf("%s: %s\n", tagName, value.String())
	}

	return
}

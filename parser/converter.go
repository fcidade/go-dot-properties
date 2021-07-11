package parser

import "reflect"

func structToMap(in interface{}) (out map[string]string) {
	out = make(map[string]string)
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
		out[tagName] = value.String()
	}
	return
}

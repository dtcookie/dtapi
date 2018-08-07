package dtapi

import (
	"errors"
	"reflect"
	"strings"
)

func omitEmpty(tag reflect.StructTag) bool {
	jsonTag, ok := tag.Lookup("json")
	if !ok {
		return false
	}
	for _, value := range strings.Split(jsonTag, ",") {
		if value == "omitempty" {
			return true
		}
	}
	return false
}

func assert(v interface{}) error {
	structValue := reflect.ValueOf(v)
	structType := reflect.TypeOf(v)

	structName := structType.Name()
	for idx := 0; idx < structType.NumField(); idx++ {
		field := structType.Field(idx)
		fieldName := field.Name
		fieldType := field.Type

		if fieldType.Kind() == reflect.String {
			fieldValue := structValue.Field(idx)
			value := fieldValue.Interface()
			tag := field.Tag
			if value == "" && !omitEmpty(tag) {
				return errors.New(structName + "." + fieldName + " must not be empty")
			}
		}
	}

	return nil
}

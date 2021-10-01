package share

import (
	"reflect"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// GetTagName get the tag name of the reflect.StructField
func GetTagName(field reflect.StructField, name string) string {
	tag := field.Tag.Get(name)
	if tag == "" {
		tag = Snake(field.Name)
		return tag
	}

	tagr := strings.Split(tag, ",")
	return tagr[0]
}

// Snake The Snake method converts the given string to snake_case
func Snake(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

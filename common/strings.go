package common

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchAllCap = regexp.MustCompile("([a-z])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

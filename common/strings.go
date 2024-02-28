package common

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var MatchAllCap = regexp.MustCompile("(a-z0-9)([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "%{1}_${2}")
	snake = MatchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

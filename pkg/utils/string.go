package utils

import (
	"regexp"
	"strings"
)

// ToSnakeCase is a function to convert a string to snake case
func ToSnakeCase(str string) string {
	snake := regexp.MustCompile("(.)([A-Z][a-z]+)").ReplaceAllString(str, "${1}_${2}")
	snake = regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

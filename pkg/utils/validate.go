package utils

import (
	"regexp"
)

// ValidateDash is a function to validate a string only contains letters, numbers, underscores, and dashes
func ValidateDash(name string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9_-]*$`)

	return re.MatchString(name)
}

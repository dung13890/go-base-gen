package template

import (
	"embed"
	"strings"
	"text/template"

	"github.com/dung13890/go-base-gen/pkg/utils"
)

var (
	//go:embed tmpl/*
	templateDir embed.FS
	// Funcs is a map of functions for template
	Funcs = template.FuncMap{
		"upper":      strings.ToUpper,
		"capitalize": capitalize,
		"snakecase":  snakecase,
	}
)

// NewTemplate is a function to create new template
func NewTemplate(name string, patterns []string) (*template.Template, error) {
	return template.
		New(name).
		Funcs(Funcs).
		ParseFS(templateDir, patterns...)
}

// capitalize is a function to uppercase first character of string
func capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func snakecase(s string) string {
	return utils.ToSnakeCase(s)
}

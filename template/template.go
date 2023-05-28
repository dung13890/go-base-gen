package template

import (
	"embed"
	"strings"
	"text/template"
)

var (
	//go:embed tmpl/*
	templateDir embed.FS
	// Funcs is a map of functions for template
	Funcs = template.FuncMap{
		"upper":      strings.ToUpper,
		"capitalize": capitalize,
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

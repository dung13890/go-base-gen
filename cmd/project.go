package cmd

import (
	"context"
	"embed"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
)

var templateFs embed.FS

var (
	errDirExists = errors.New("Directory already exist!")
	structs      = []string{
		"cmd/app",
		"cmd/migrate",
		"cmd/seed",
		"config",
		"db/migrations",
		"db/seeds",
		"internal/app",
		"internal/constants",
		"internal/domain/auth/delivery/http",
		"internal/domain/auth/delivery/grpc",
		"internal/domain/auth/repository",
		"internal/domain/auth/usecase",
		"internal/modules",
		"internal/impl/pubsub",
		"internal/impl/service",
		"internal/registry",
		"pkg",
	}

	files = []string{
		".golint.yaml.tmpl",
		".editorconfig.tmpl",
		"go.mod.tmpl",
		"Makefile.tmpl",
	}
)

type project struct {
	Module string
	Path   string
}

// NewProject is a function to create new project command
func NewProject(tFs embed.FS) *cli.Command {
	templateFs = tFs
	return &cli.Command{
		Name:  "project",
		Usage: "Generate base code for go project use clean architecture",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "module",
				Aliases:  []string{"m"},
				Usage:    "Module name for the project",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "path",
				Aliases:     []string{"p"},
				Usage:       "Path is a path to generate the project",
				DefaultText: "./",
			},
		},
		Action: func(ctx *cli.Context) error {
			dir, err := os.Getwd()
			if err != nil {
				return err
			}
			if ctx.String("path") != "" {
				dir = ctx.String("path")
			}

			rltDir, err := filepath.Abs(dir)
			if err != nil {
				return err
			}

			p := &project{
				Module: ctx.String("module"),
				Path:   rltDir,
			}
			// Create dir
			if err := p.createDir(ctx.Context); err != nil {
				return err
			}

			// Generate struct
			if err := p.generateStruct(ctx.Context); err != nil {
				if dErr := p.destroy(ctx.Context); dErr != nil {
					return errors.Join(err, dErr)
				}
				return err
			}

			// Generate file
			if err := p.generateFile(ctx.Context, templateFs); err != nil {
				if dErr := p.destroy(ctx.Context); dErr != nil {
					return errors.Join(err, dErr)
				}
				return err
			}

			return cli.Exit("Successfully created!", 0)
		},
	}
}

// createDir is a function to create directory for project
func (p *project) createDir(context.Context) error {
	dir := filepath.Join(p.Path, p.Module)
	// Check directory exist or not
	if _, err := os.Stat(dir); !errors.Is(err, os.ErrNotExist) {
		return errDirExists
	}
	// Create dir for project
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// destroy is a function to destroy project
func (p *project) destroy(context.Context) error {
	dir := filepath.Join(p.Path, p.Module)

	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	return nil
}

// generateStruct is a function to generate struct for project
func (p *project) generateStruct(context.Context) error {
	dir := filepath.Join(p.Path, p.Module)
	// Generate struct
	for _, s := range structs {
		if err := os.MkdirAll(filepath.Join(dir, s), os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

// generateFile is a function to generate file for project
func (p *project) generateFile(_ context.Context, content embed.FS) error {
	dir := filepath.Join(p.Path, p.Module)
	tmpl := template.Must(template.New("tmpl").ParseFS(content, "template/*.tmpl"))

	for _, f := range files {
		target := filepath.Join(dir, strings.TrimSuffix(f, ".tmpl"))
		f, err := os.Create(filepath.Clean(target))
		if err != nil {
			return err
		}

		fileName := filepath.Base(target)
		if err := tmpl.ExecuteTemplate(f, fileName, p); err != nil {
			return err
		}

		if err := f.Close(); err != nil {
			return err
		}
	}

	return nil
}

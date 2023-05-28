package cmd

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/dung13890/go-base-gen/pkg/utils"
	"github.com/dung13890/go-base-gen/template"
	"github.com/urfave/cli/v2"
)

var (
	fCreated []string
)

var (
	errProjectNotExists  = errors.New("Project is not exists!")
	errNameDomainInvalid = errors.New("Name of domain is invalid!")
	errProjectInvalid    = errors.New("Name of project is invalid!")
	errModuleInvalid     = errors.New("Name of module is invalid!")
	dStructs             = []string{
		"internal/domain",
		"internal/modules/:name/delivery/http",
		"internal/modules/:name/usecase",
		"internal/modules/:name/repository",
	}
	dFiles = []string{
		"internal/domain/domain.go.tmpl",
		"internal/modules/name/delivery/http/domain_http.go.tmpl",
		"internal/modules/name/delivery/http/domain_dto.go.tmpl",
		"internal/modules/name/delivery/http/handler.go.tmpl",
		"internal/modules/name/delivery/http/middleware.go.tmpl",
		"internal/modules/name/usecase/domain_uc.go.tmpl",
		"internal/modules/name/usecase/usecase.go.tmpl",
		"internal/modules/name/repository/domain_repo.go.tmpl",
		"internal/modules/name/repository/domain_dao.go.tmpl",
		"internal/modules/name/repository/repository.go.tmpl",
	}
)

type domain struct {
	Domain  string
	Project string
	Module  string
	Path    string
}

// NewDomain is a function to create new domain in project
func NewDomain() *cli.Command {
	return &cli.Command{
		Name:  "domain",
		Usage: "Create new domain in project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "Name of domain (ex: product)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "project",
				Aliases:  []string{"pj"},
				Usage:    "Name of project (ex: project-demo)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "module",
				Aliases:  []string{"m"},
				Usage:    "Name of module inside project (ex: ecommerce)",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "path",
				Aliases:     []string{"p"},
				Usage:       "Path is a path to generate for domain will stay in project path",
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
			// Validate
			if err := domainValidate(ctx); err != nil {
				return err
			}

			d := &domain{
				Domain:  strings.ToLower(ctx.String("name")),
				Project: strings.ToLower(ctx.String("project")),
				Module:  strings.ToLower(ctx.String("module")),
				Path:    rltDir,
			}
			if err := d.checkProject(ctx.Context); err != nil {
				return err
			}
			if err := d.generateStruct(ctx.Context); err != nil {
				return err
			}

			if err := d.generateFile(ctx.Context); err != nil {
				if dErr := d.destroy(ctx.Context); dErr != nil {
					return errors.Join(err, dErr)
				}
				return err
			}

			fCreated = []string{}
			return cli.Exit("Successfully created!", 0)
		},
	}
}

func domainValidate(ctx *cli.Context) error {
	// Validate
	if ok := utils.ValidateDash(ctx.String("name")); !ok {
		return errNameDomainInvalid
	}
	if ok := utils.ValidateDash(ctx.String("project")); !ok {
		return errProjectInvalid
	}
	if ok := utils.ValidateDash(ctx.String("module")); !ok {
		return errModuleInvalid
	}

	return nil
}

// checkProject is a function to check project exist or not
func (d *domain) checkProject(context.Context) error {
	dir := filepath.Join(d.Path, d.Project)
	// Check project exist or not
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		return errProjectNotExists
	}

	return nil
}

// destroy is a function to destroy created file if generate file failed
func (*domain) destroy(context.Context) error {
	if len(fCreated) > 0 {
		for _, f := range fCreated {
			if err := os.Remove(f); err != nil {
				return err
			}
		}
	}

	return nil
}

// generateStruct is a function to generate struct for domain
func (d *domain) generateStruct(context.Context) error {
	dir := filepath.Join(d.Path, d.Project)
	// Generate struct
	for _, s := range dStructs {
		target := strings.Replace(s, ":name", d.Module, 1)
		// Check directory exist or not
		if _, err := os.Stat(filepath.Join(dir, target)); !errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err := os.MkdirAll(filepath.Join(dir, target), os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

// generateFile is a function to generate file for domain
func (d *domain) generateFile(_ context.Context) error {
	dir := filepath.Join(d.Path, d.Project)
	tmpl, err := template.NewTemplate("tmpl", []string{
		"tmpl/*/*.tmpl",
		"tmpl/*/*/*.tmpl",
		"tmpl/*/*/*/*/*.tmpl",
		"tmpl/*/*/*/*/*/*.tmpl",
	})
	if err != nil {
		return err
	}
	rl := strings.NewReplacer(
		"/modules/name/", "/modules/"+d.Module+"/",
		"domain.go", d.Domain+".go",
		"domain_http.go", d.Domain+"_http.go",
		"domain_grpc.go", d.Domain+"_grpc.go",
		"domain_dto.go", d.Domain+"_dto.go",
		"domain_repo.go", d.Domain+"_repo.go",
		"domain_dao.go", d.Domain+"_dao.go",
		"domain_uc.go", d.Domain+"_uc.go",
	)
	for _, f := range dFiles {
		index := strings.TrimSuffix(f, ".tmpl")
		target := filepath.Join(dir, rl.Replace(index))

		// Check file exist or not
		if _, err := os.Stat(filepath.Clean(target)); !errors.Is(err, os.ErrNotExist) {
			continue
		}

		f, err := os.Create(filepath.Clean(target))
		if err != nil {
			return err
		}
		if err := tmpl.ExecuteTemplate(f, index, d); err != nil {
			fCreated = append(fCreated, filepath.Clean(target))
			return err
		}
		fCreated = append(fCreated, filepath.Clean(target))

		if err := f.Close(); err != nil {
			return err
		}
	}

	return nil
}

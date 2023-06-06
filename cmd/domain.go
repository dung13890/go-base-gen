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
	errProjectNotExists  = errors.New("Project is not exists!. Please create project first!")
	errNameDomainInvalid = errors.New("Name of domain is invalid!")
	errProjectInvalid    = errors.New("Name of project is invalid!")
	errModuleInvalid     = errors.New("Name of module is invalid!")
	dStructs             = []string{
		"internal/domain",
		"internal/modules/:name/delivery/http",
		"internal/modules/:name/usecase",
		"internal/modules/:name/repository",
	}
	dStructsWithoutModule = []string{
		"internal/domain",
		"internal/delivery/http",
		"internal/usecase",
		"internal/repository",
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
	dFilesWithoutModule = []string{
		"internal/domain/domain.go.tmpl",
		"internal/delivery/http/domain_http.go.tmpl",
		"internal/delivery/http/domain_dto.go.tmpl",
		"internal/delivery/http/handler.go.tmpl",
		"internal/delivery/http/middleware.go.tmpl",
		"internal/usecase/domain_uc.go.tmpl",
		"internal/usecase/usecase.go.tmpl",
		"internal/repository/domain_repo.go.tmpl",
		"internal/repository/domain_dao.go.tmpl",
		"internal/repository/repository.go.tmpl",
	}
)

type domain struct {
	Domain    string
	Project   string
	Module    string
	Path      string
	ForcePath bool
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
				Usage:    "Name of domain is required! (ex: product)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "project",
				Aliases:  []string{"pj"},
				Usage:    "Name of project is required! (ex: project-demo)",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "module",
				Aliases: []string{"m"},
				Usage:   "Name of module inside project (ex: ecommerce)",
			},
			&cli.StringFlag{
				Name:        "path",
				Aliases:     []string{"p"},
				Usage:       "Path is a path to generate for domain will stay in project path",
				DefaultText: "./",
			},
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Force path is a flag to force path to generate for domain will stay in path",
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
				Domain:    ctx.String("name"),
				Project:   strings.ToLower(ctx.String("project")),
				Module:    strings.ToLower(ctx.String("module")),
				Path:      rltDir,
				ForcePath: ctx.Bool("force"),
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

// domainValidate is a function to validate input
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

// getDir will check path is contain project or not and return path include project
func getDir(path, project string, forcePath bool) string {
	if forcePath {
		return path
	}
	if ok := strings.HasSuffix(path, project); ok {
		return path
	}

	return filepath.Join(path, project)
}

// checkProject is a function to check project exist or not
func (d *domain) checkProject(context.Context) error {
	dir := getDir(d.Path, d.Project, d.ForcePath)
	modDir := filepath.Join(dir, "go.mod")
	// Check project exist or not
	if _, err := os.Stat(modDir); errors.Is(err, os.ErrNotExist) {
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
	dir := getDir(d.Path, d.Project, d.ForcePath)
	structs := dStructsWithoutModule
	if d.Module != "" {
		structs = dStructs
	}
	// Generate struct
	for _, s := range structs {
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
	dir := getDir(d.Path, d.Project, d.ForcePath)
	tmpl, err := template.NewTemplate("tmpl", []string{
		"tmpl/*/*.tmpl",
		"tmpl/*/*/*.tmpl",
		"tmpl/*/*/*/*.tmpl",
		"tmpl/*/*/*/*/*.tmpl",
		"tmpl/*/*/*/*/*/*.tmpl",
	})
	if err != nil {
		return err
	}
	snakeDomain := utils.ToSnakeCase(d.Domain)
	rl := strings.NewReplacer(
		"/modules/name/", "/modules/"+d.Module+"/",
		"domain.go", snakeDomain+".go",
		"domain_http.go", snakeDomain+"_http.go",
		"domain_grpc.go", snakeDomain+"_grpc.go",
		"domain_dto.go", snakeDomain+"_dto.go",
		"domain_repo.go", snakeDomain+"_repo.go",
		"domain_dao.go", snakeDomain+"_dao.go",
		"domain_uc.go", snakeDomain+"_uc.go",
	)
	files := dFilesWithoutModule
	if d.Module != "" {
		files = dFiles
	}
	for _, f := range files {
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

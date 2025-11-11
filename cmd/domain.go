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
	dStructs             = []string{
		"internal/adapter/repository",
		"internal/delivery/http",
		"internal/delivery/http/dto",
		"internal/service",
		"internal/domain/repository",
		"internal/domain/entity",
		"internal/domain/service",
		"internal/usecase/:name",
	}
	dFiles = []string{
		"internal/adapter/repository/domain_dao.go.tmpl",
		"internal/adapter/repository/domain_repo.go.tmpl",
		"internal/delivery/http/dto/domain_dto.go.tmpl",
		"internal/delivery/http/domain_handler.go.tmpl",
		"internal/service/domain_svc.go.tmpl",
		"internal/domain/repository/domain_repo.go.tmpl",
		"internal/domain/entity/domain.go.tmpl",
		"internal/domain/service/domain_svc.go.tmpl",
		"internal/usecase/name/usecase.go.tmpl",
	}
)

type domain struct {
	Domain    string
	Project   string
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
				Name:     "domain",
				Aliases:  []string{"dn"},
				Usage:    "Name of domain is required! (ex: product)",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "package",
				Aliases:     []string{"pkg"},
				Usage:       "Package is the name of the project",
				DefaultText: "github.com/username/projectname",
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "path",
				Aliases:     []string{"p"},
				Usage:       "Path is a path to generate for domain will stay in project path",
				DefaultText: "project_new",
				Value:       "project_new",
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
				Domain:    ctx.String("domain"),
				Project:   strings.ToLower(ctx.String("package")),
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
	if ok := utils.ValidateDomain(ctx.String("domain")); !ok {
		return errNameDomainInvalid
	}
	if ok := utils.ValidateDash(ctx.String("package")); !ok {
		return errProjectInvalid
	}

	return nil
}

// getDir will check path is contain project or not and return path include project
func getDir(path string, forcePath bool) string {
	if forcePath {
		return path
	}

	return filepath.Join(path)
}

// checkProject is a function to check project exist or not
func (d *domain) checkProject(context.Context) error {
	dir := getDir(d.Path, d.ForcePath)
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
	dir := getDir(d.Path, d.ForcePath)
	structs := dStructs
	snakeDomain := utils.ToSnakeCase(d.Domain)
	// Generate struct
	for _, s := range structs {
		target := strings.Replace(s, ":name", snakeDomain, 1)
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
	dir := getDir(d.Path, d.ForcePath)
	tmpl, err := template.NewTemplate("tmpl", []string{
		"tmpl/*/*/*.tmpl",
		"tmpl/*/*/*/*.tmpl",
		"tmpl/*/*/*/*/*.tmpl",
	})
	if err != nil {
		return err
	}
	snakeDomain := utils.ToSnakeCase(d.Domain)
	rl := strings.NewReplacer(
		"usecase/name/", "/usecase/"+snakeDomain+"/",
		"domain.go", snakeDomain+".go",
		"domain_handler.go", snakeDomain+"_handler.go",
		"domain_dto.go", snakeDomain+"_dto.go",
		"domain_repo.go", snakeDomain+"_repo.go",
		"domain_dao.go", snakeDomain+"_dao.go",
		"domain_svc.go", snakeDomain+"_svc.go",
	)
	files := dFiles
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

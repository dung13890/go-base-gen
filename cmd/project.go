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
	errDirExists          = errors.New("Directory already exist!")
	errNameProjectInvalid = errors.New("Name of project is invalid!")
	pStructs              = []string{
		"cmd/app",
		"cmd/migrate",
		"db/migrations",
		"internal/adapter/repository",
		"internal/adapter/cache",
		"internal/registry",
		"internal/delivery/http",
		"internal/delivery/http/dto",
		"internal/service",
		"internal/infrastructure/config",
		"internal/infrastructure/constants",
		"internal/infrastructure/database",
		"internal/infrastructure/redis",
		"internal/usecase",
		"internal/domain/entity",
		"internal/domain/repository",
		"internal/domain/service",
		"internal/usecase/auth",
		"pkg/logger",
		"pkg/utils",
		"pkg/errors",
	}

	pFiles = []string{
		".golint.yml.tmpl",
		".editorconfig.tmpl",
		".env.example.tmpl",
		"go.mod.tmpl",
		"Makefile.tmpl",
		"cmd/app/main.go.tmpl",
		"cmd/migrate/main.go.tmpl",
		"db/migrations/20220705080200_create_auth_table.down.sql.tmpl",
		"db/migrations/20220705080200_create_auth_table.up.sql.tmpl",
		"internal/adapter/repository/user_dao.go.tmpl",
		"internal/adapter/repository/user_repo.go.tmpl",
		"internal/adapter/cache/redis_store.go.tmpl",
		"internal/adapter/cache/client.go.tmpl",
		"internal/registry/registry.go.tmpl",
		"internal/delivery/http/dto/auth_dto.go.tmpl",
		"internal/delivery/http/middleware.go.tmpl",
		"internal/delivery/http/auth_handler.go.tmpl",
		"internal/delivery/http/handler.go.tmpl",
		"internal/service/jwt_svc.go.tmpl",
		"internal/infrastructure/database/migrate.go.tmpl",
		"internal/infrastructure/database/postgres.go.tmpl",
		"internal/infrastructure/config/redis.go.tmpl",
		"internal/infrastructure/config/app.go.tmpl",
		"internal/infrastructure/config/database.go.tmpl",
		"internal/infrastructure/redis/redis.go.tmpl",
		"internal/infrastructure/constants/http.go.tmpl",
		"internal/usecase/auth/usecase.go.tmpl",
		"internal/usecase/auth/signin_uc.go.tmpl",
		"internal/domain/repository/user_repo.go.tmpl",
		"internal/domain/entity/user.go.tmpl",
		"internal/domain/service/jwt_svc.go.tmpl",
		"pkg/logger/logger.go.tmpl",
		"pkg/utils/bcrypt.go.tmpl",
		"pkg/errors/errors.go.tmpl",
		"pkg/errors/code.go.tmpl",
	}
)

type project struct {
	Project string
	Path    string
}

// NewProject is a function to create new project command
func NewProject() *cli.Command {
	return &cli.Command{
		Name:  "project",
		Usage: "Generate base code for go project use clean architecture",
		Flags: []cli.Flag{
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
				Usage:       "Path is a path to generate the project",
				DefaultText: "project_new",
				Value:       "project_new",
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
			if err := projectValidate(ctx); err != nil {
				return err
			}

			p := &project{
				Project: strings.ToLower(ctx.String("package")),
				Path:    rltDir,
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
			if err := p.generateFile(ctx.Context); err != nil {
				if dErr := p.destroy(ctx.Context); dErr != nil {
					return errors.Join(err, dErr)
				}
				return err
			}

			return cli.Exit("Successfully created!", 0)
		},
	}
}

// projectValidate is a function to validate project name
func projectValidate(ctx *cli.Context) error {
	if ok := utils.ValidateDash(ctx.String("package")); !ok {
		return errNameProjectInvalid
	}
	return nil
}

// createDir is a function to create directory for project
func (p *project) createDir(context.Context) error {
	dir := filepath.Join(p.Path)
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
	dir := filepath.Join(p.Path, p.Project)

	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	return nil
}

// generateStruct is a function to generate struct for project
func (p *project) generateStruct(context.Context) error {
	dir := filepath.Join(p.Path)
	// Generate struct
	for _, s := range pStructs {
		if err := os.MkdirAll(filepath.Join(dir, s), os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

// generateFile is a function to generate file for project
func (p *project) generateFile(_ context.Context) error {
	dir := filepath.Join(p.Path)
	tmpl, err := template.NewTemplate("tmpl", []string{
		"tmpl/*.tmpl",
		"tmpl/*/*/*.tmpl",
		"tmpl/*/*/*/*.tmpl",
		"tmpl/*/*/*/*/*.tmpl",
	})
	if err != nil {
		return err
	}

	for _, f := range pFiles {
		index := strings.TrimSuffix(f, ".tmpl")
		target := filepath.Join(dir, index)
		f, err := os.Create(filepath.Clean(target))
		if err != nil {
			return err
		}

		if err := tmpl.ExecuteTemplate(f, index, p); err != nil {
			return err
		}

		if err := f.Close(); err != nil {
			return err
		}
	}

	return nil
}

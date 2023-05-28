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
		"cmd/seed",
		"config",
		"db/migrations",
		"db/seeds",
		"internal/app",
		"internal/constants",
		"internal/domain",
		"internal/modules/auth/delivery/http",
		"internal/modules/auth/delivery/grpc",
		"internal/modules/auth/repository",
		"internal/modules/auth/usecase",
		"internal/impl/pubsub",
		"internal/impl/service",
		"internal/registry",
		"pkg/cache",
		"pkg/errors",
		"pkg/logger",
		"pkg/postgres",
		"pkg/redis",
		"pkg/utils",
		"pkg/validate",
	}

	pFiles = []string{
		".golint.yml.tmpl",
		".editorconfig.tmpl",
		".env.example.tmpl",
		"go.mod.tmpl",
		"Makefile.tmpl",
		"cmd/app/main.go.tmpl",
		"cmd/app/air.toml.tmpl",
		"cmd/migrate/main.go.tmpl",
		"cmd/seed/main.go.tmpl",
		"config/app.go.tmpl",
		"config/database.go.tmpl",
		"config/redis.go.tmpl",
		"db/migrations/20220705080200_create_auth_table.down.sql.tmpl",
		"db/migrations/20220705080200_create_auth_table.up.sql.tmpl",
		"db/seeds/data.json.tmpl",
		"internal/app/app.go.tmpl",
		"internal/app/seed.go.tmpl",
		"internal/constants/http.go.tmpl",
		"internal/domain/auth.go.tmpl",
		"internal/domain/role.go.tmpl",
		"internal/domain/user.go.tmpl",
		"internal/domain/password_reset.go.tmpl",
		"internal/domain/jwt_svc.go.tmpl",
		"internal/domain/throttle_svc.go.tmpl",
		"internal/impl/service/jwt_svc.go.tmpl",
		"internal/impl/service/throttle_svc.go.tmpl",
		"internal/modules/auth/delivery/http/middleware.go.tmpl",
		"internal/modules/auth/delivery/http/handler.go.tmpl",
		"internal/modules/auth/delivery/http/auth_dto.go.tmpl",
		"internal/modules/auth/delivery/http/auth_http.go.tmpl",
		"internal/modules/auth/delivery/http/role_dto.go.tmpl",
		"internal/modules/auth/delivery/http/role_http.go.tmpl",
		"internal/modules/auth/delivery/http/user_dto.go.tmpl",
		"internal/modules/auth/delivery/http/user_http.go.tmpl",
		"internal/modules/auth/repository/repository.go.tmpl",
		"internal/modules/auth/repository/password_reset_dao.go.tmpl",
		"internal/modules/auth/repository/password_reset_repo.go.tmpl",
		"internal/modules/auth/repository/role_dao.go.tmpl",
		"internal/modules/auth/repository/role_repo.go.tmpl",
		"internal/modules/auth/repository/user_dao.go.tmpl",
		"internal/modules/auth/repository/user_repo.go.tmpl",
		"internal/modules/auth/usecase/usecase.go.tmpl",
		"internal/modules/auth/usecase/auth_uc.go.tmpl",
		"internal/modules/auth/usecase/user_uc.go.tmpl",
		"internal/modules/auth/usecase/role_uc.go.tmpl",
		"internal/registry/http.go.tmpl",
		"internal/registry/repository.go.tmpl",
		"internal/registry/service.go.tmpl",
		"internal/registry/usecase.go.tmpl",
		"pkg/cache/client.go.tmpl",
		"pkg/cache/redis_store.go.tmpl",
		"pkg/errors/code.go.tmpl",
		"pkg/errors/errors.go.tmpl",
		"pkg/logger/logger.go.tmpl",
		"pkg/postgres/migrate.go.tmpl",
		"pkg/postgres/postgres.go.tmpl",
		"pkg/redis/redis.go.tmpl",
		"pkg/utils/bcrypt.go.tmpl",
		"pkg/utils/md5.go.tmpl",
		"pkg/utils/rand_string.go.tmpl",
		"pkg/utils/slug.go.tmpl",
		"pkg/utils/uuid.go.tmpl",
		"pkg/validate/base.go.tmpl",
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
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "Name of project with module format (ex: project-demo)",
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

			// Validate
			if err := projectValidate(ctx); err != nil {
				return err
			}

			p := &project{
				Project: strings.ToLower(ctx.String("name")),
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

func projectValidate(ctx *cli.Context) error {
	if ok := utils.ValidateDash(ctx.String("name")); !ok {
		return errNameProjectInvalid
	}
	return nil
}

// createDir is a function to create directory for project
func (p *project) createDir(context.Context) error {
	dir := filepath.Join(p.Path, p.Project)
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
	dir := filepath.Join(p.Path, p.Project)
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
	dir := filepath.Join(p.Path, p.Project)
	tmpl, err := template.NewTemplate("tmpl", []string{
		"tmpl/*.tmpl",
		"tmpl/*/*.tmpl",
		"tmpl/*/*/*.tmpl",
		"tmpl/*/*/*/*.tmpl",
		"tmpl/*/*/*/*/*.tmpl",
		"tmpl/*/*/*/*/*/*.tmpl",
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
